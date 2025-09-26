package booklending

import (
	"digitalbooklending/apps/repositories"

	"github.com/go-playground/validator/v10"
)

type bookLending struct {
	db        repositories.IDatabase
	validator *validator.Validate
}

func NewBookLendingService(
	db repositories.IDatabase,
	validator *validator.Validate,
) *bookLending {
	return &bookLending{
		db:        db,
		validator: validator,
	}
}
