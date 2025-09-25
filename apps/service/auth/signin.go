package auth

import (
	"context"
	"digitalbooklending/apps/domain"
	"digitalbooklending/apps/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/env"
	"github.com/vizucode/gokit/utils/errorkit"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) SignIn(ctx context.Context, req domain.SignInRequest) (resp domain.SignInResponse, err error) {

	// validate request
	if err = s.validator.StructCtx(ctx, req); err != nil {
		logger.Log.Error(ctx, err)
		return domain.SignInResponse{}, err
	}

	// TODO: implement
	user, err := s.db.FirstUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Log.Error(ctx, err)
		return domain.SignInResponse{}, err
	}

	if user.Id < 1 {
		logger.Log.Error(ctx, err)
		return domain.SignInResponse{}, errorkit.NewErrorStd(http.StatusNotFound, "", "user not found")
	}

	// compare password
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		logger.Log.Error(ctx, err)
		return domain.SignInResponse{}, errorkit.NewErrorStd(http.StatusUnauthorized, "", "invalid password")
	}

	// generate access token
	accessToken, err := s.generateAccessToken(ctx, user)
	if err != nil {
		logger.Log.Error(ctx, err)
		return domain.SignInResponse{}, err
	}

	return domain.SignInResponse{
		Token: accessToken,
	}, nil
}

func (s *authService) generateAccessToken(ctx context.Context, user models.Users) (string, error) {

	// Get secret key from environment
	secretKey := env.GetString("ACCESS_SECRET_KEY")
	if secretKey == "" {
		err := errorkit.NewErrorStd(http.StatusInternalServerError, "", "JWT secret key not configured")
		logger.Log.Error(ctx, err)
		return "", err
	}

	// Set token expiration time (24 hours from now)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create JWT claims
	claims := jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   expirationTime.Unix(),
		"iat":   time.Now().Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(env.GetString("ACCESS_SECRET_KEY")))
	if err != nil {
		logger.Log.Error(ctx, err)
		return "", err
	}

	return tokenString, nil
}
