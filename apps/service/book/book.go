package book

import (
	"digitalbooklending/apps/repositories"

	"github.com/go-playground/validator/v10"
)

type book struct {
	validator *validator.Validate
	db        repositories.IDatabase
}

func NewBook(
	validator *validator.Validate,
	db repositories.IDatabase) *book {
	return &book{validator: validator, db: db}
}
