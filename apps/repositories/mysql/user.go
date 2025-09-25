package mysql

import (
	"context"
	"digitalbooklending/apps/models"
	"errors"
	"net/http"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"gorm.io/gorm"
)

func (r *mysql) CreateUser(ctx context.Context, user models.Users) (err error) {
	err = r.db.WithContext(ctx).Model(&models.Users{}).Create(&user).Error
	if err != nil {
		logger.Log.Error(ctx, err)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errorkit.NewErrorStd(http.StatusBadRequest, "", "email already exists")
		}

		if errors.Is(err, gorm.ErrInvalidData) {
			return errorkit.NewErrorStd(http.StatusBadRequest, "", "invalid data")
		}

		return errorkit.NewErrorStd(http.StatusInternalServerError, "", "internal server error")
	}

	return nil
}

func (r *mysql) FirstUserById(ctx context.Context, id string) (user models.Users, err error) {
	err = r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Log.Error(ctx, err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errorkit.NewErrorStd(http.StatusNotFound, "", "user not found")
		}

		return user, errorkit.NewErrorStd(http.StatusInternalServerError, "", "internal server error")
	}

	return
}

func (r *mysql) FirstUserByEmail(ctx context.Context, email string) (user models.Users, err error) {
	err = r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		logger.Log.Error(ctx, err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errorkit.NewErrorStd(http.StatusNotFound, "", "user not found")
		}

		return user, errorkit.NewErrorStd(http.StatusInternalServerError, "", "internal server error")
	}

	return
}
