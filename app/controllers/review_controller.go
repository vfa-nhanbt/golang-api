package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/app/db/repositories"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	pkgRepo "github.com/vfa-nhanbt/todo-api/pkg/repositories"
)

type ReviewController struct {
	Repository *repositories.ReviewRepository
}

func (controller *ReviewController) AddReviewHandler(c *fiber.Ctx) error {
	// reviewModel := &models.ReviewModel{}
	var addReviewRequest struct {
		Description string    `json:"description" validate:"required"`
		BookID      uuid.UUID `json:"book_id" validate:"required, uuid"`
	}

	/// Validate body request
	err := helpers.ValidateRequestBody(&addReviewRequest, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	/// Set viewer for this review
	userId, err := helpers.GetUserIdFromToken(c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	reviewModel := &models.ReviewModel{
		ID:          uuid.New(),
		Description: addReviewRequest.Description,
		BookID:      addReviewRequest.BookID,
		Book: models.BookModel{
			ID: addReviewRequest.BookID,
		},
		ViewerID: userUUID,
		Viewer: models.ViewerModel{
			UserID: userUUID,
		},
	}
	/// Create review to DB
	err = controller.Repository.AddReview(reviewModel)
	if err != nil {
		res := pkgRepo.BaseResponse{
			Code:      "e-review-001",
			IsSuccess: false,
			Data:      "Cannot insert review record to table with error: " + err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	/// Insert success, return status OK
	res := pkgRepo.BaseResponse{
		Code:      "s-review-001",
		IsSuccess: true,
		Data:      reviewModel,
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func (controller *ReviewController) DeleteReviewByID(c *fiber.Ctx) error {
	reviewId := c.Params("id")

	/// Get current user id from token
	userId, err := helpers.GetUserIdFromToken(c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Find review from DB
	reviewFromDB, err := controller.Repository.GetReviewById(c.Params("id"))
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	/// Check if user is author of this review
	if reviewFromDB.ViewerID.String() != userId {
		res := pkgRepo.BaseResponse{
			Code:      "e-review-001",
			IsSuccess: false,
			Data:      "user is not owner of this review",
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	/// Delete review from DB
	err = controller.Repository.DeleteReviewById(reviewId)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Delete success, return status OK
	res := pkgRepo.BaseResponse{
		Code:      "s-review-001",
		IsSuccess: true,
		Data:      "Delete review successfully",
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func (controller *ReviewController) UpdateReview(c *fiber.Ctx) error {
	var updateReviewRequest struct {
		Description string `json:"description" validate:"required"`
		helpers.StructHelper
	}
	/// Validate request body
	err := helpers.ValidateRequestBody(&updateReviewRequest, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Get current user id from token
	userId, err := helpers.GetUserIdFromToken(c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Find review from DB
	reviewFromDB, err := controller.Repository.GetReviewById(c.Params("id"))
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	/// Check if user is author of this review
	if reviewFromDB.ViewerID.String() != userId {
		res := pkgRepo.BaseResponse{
			Code:      "e-review-001",
			IsSuccess: false,
			Data:      "user is not owner of this review",
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	/// Update book
	err = controller.Repository.UpdateReview(reviewFromDB, updateReviewRequest.StructToUnNilMap(updateReviewRequest))
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Update success, return status OK
	res := pkgRepo.BaseResponse{
		Code:      "s-review-001",
		IsSuccess: true,
		Data:      "update review successfully",
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}
