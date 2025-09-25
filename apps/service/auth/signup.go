package auth

import (
	"context"
	"digitalbooklending/apps/domain"
	"digitalbooklending/apps/models"
	"net/http"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) SignUp(ctx context.Context, req domain.SignUpRequest) (err error) {

	// validate request
	if err = s.validator.StructCtx(ctx, req); err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusInternalServerError, "", "failed to hash password")
	}

	// TODO: implement
	err = s.db.CreateUser(ctx, models.Users{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}
