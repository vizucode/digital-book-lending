package auth

import (
	"digitalbooklending/apps/repositories"

	"github.com/go-playground/validator/v10"
)

type authService struct {
	db        repositories.IDatabase
	validator *validator.Validate
}

func NewAuthService(
	db repositories.IDatabase,
	validator *validator.Validate,
) *authService {
	return &authService{
		db:        db,
		validator: validator,
	}
}
