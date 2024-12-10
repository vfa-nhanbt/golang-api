package helpers

import "gorm.io/gorm"

type Paginate struct {
	Page int
	Limit int
}

func NewPagination(page int, limit int) *Paginate {
	return &Paginate{Page: page, Limit: limit}
}

func (p *Paginate) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit
	return db.Offset(offset).Limit(p.Limit)
}