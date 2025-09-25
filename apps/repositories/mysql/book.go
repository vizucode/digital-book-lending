package mysql

import (
	"context"
	"digitalbooklending/apps/models"
	"digitalbooklending/helpers/constants/httpstd"
	"errors"
	"net/http"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"gorm.io/gorm"
)

func (r *mysql) CreateBook(ctx context.Context, book models.Book) error {
	result := r.db.WithContext(ctx).Create(&book)
	if result.Error != nil {
		logger.Log.Error(ctx, result.Error)
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return errorkit.NewErrorStd(http.StatusBadRequest, "", "Error Duplicate")
		}

		return errorkit.NewErrorStd(http.StatusBadRequest, "", httpstd.InternalServerError)
	}
	return nil
}

func (r *mysql) FirstBookById(ctx context.Context, id uint) (models.Book, error) {
	var book models.Book
	result := r.db.WithContext(ctx).First(&book, id)
	if result.Error != nil {
		logger.Log.Error(ctx, result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Book{}, nil
		}
		return models.Book{}, errorkit.NewErrorStd(http.StatusBadRequest, "", httpstd.InternalServerError)
	}
	return book, nil
}

func (r *mysql) GetListBooks(ctx context.Context, filter models.Filter) ([]models.Book, error) {
	var books []models.Book
	query := r.db.Model(&models.Book{})

	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}

	if filter.Page > 0 {
		query = query.Offset((filter.Page - 1) * filter.Limit)
	}

	result := query.Find(&books)
	if result.Error != nil {
		logger.Log.Error(ctx, result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errorkit.NewErrorStd(http.StatusBadRequest, "", httpstd.InternalServerError)
	}
	return books, nil
}

func (r *mysql) DeleteBook(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Book{}, id)
	if result.Error != nil {
		logger.Log.Error(ctx, result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errorkit.NewErrorStd(http.StatusNotFound, "", httpstd.NotFound)
		}
		return errorkit.NewErrorStd(http.StatusBadRequest, "", httpstd.InternalServerError)
	}
	return nil
}

func (r *mysql) UpdateBook(ctx context.Context, id uint, book models.Book) error {

	selectedField := []string{}

	if book.Title != "" {
		selectedField = append(selectedField, "title")
	}

	if book.Author != "" {
		selectedField = append(selectedField, "author")
	}

	if book.Isbn != "" {
		selectedField = append(selectedField, "isbn")
	}

	if book.Category != "" {
		selectedField = append(selectedField, "category")
	}

	if book.Quantity > 0 {
		selectedField = append(selectedField, "quantity")
	}

	result := r.db.WithContext(ctx).Model(&models.Book{}).Select(selectedField).Where("id = ?", id).Updates(book)
	if result.Error != nil {
		logger.Log.Error(ctx, result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errorkit.NewErrorStd(http.StatusNotFound, "", httpstd.NotFound)
		}
		return errorkit.NewErrorStd(http.StatusBadRequest, "", httpstd.InternalServerError)
	}
	return nil
}
