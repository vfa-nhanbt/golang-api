package repositories

import (
	"github.com/vfa-nhanbt/todo-api/app/models"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	DB *gorm.DB
}

func (r *ReviewRepository) AddReview(review *models.ReviewModel) error {
	err := r.DB.Create(&review).Error
	return err
}

func (r *ReviewRepository) UpdateReview(review *models.ReviewModel, updateFields map[string]interface{}) error {
	err := r.DB.Model(&review).Updates(updateFields).Error
	return err
}

func (r *ReviewRepository) DeleteReviewById(id string) error {
	err := r.DB.Delete(&models.ReviewModel{}, "id = ?", id).Error
	return err
}

func (r *ReviewRepository) GetReviewById(id string) (*models.ReviewModel, error) {
	reviewModel := models.ReviewModel{}
	err := r.DB.First(&reviewModel, "id = ?", id).Preload("Viewer").Preload("Book").Find(&reviewModel).Error
	if err != nil {
		return nil, err
	}
	return &reviewModel, nil
}
