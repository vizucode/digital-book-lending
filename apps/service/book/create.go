package book

import (
	"context"
	"digitalbooklending/apps/domain"
	"digitalbooklending/apps/models"

	"github.com/vizucode/gokit/logger"
)

func (uc *book) CreateBook(ctx context.Context, req domain.RequestBook) (resp domain.Book, err error) {

	if err = uc.validator.Struct(req); err != nil {
		return resp, err
	}

	payload := models.Book{
		Title:    req.Title,
		Author:   req.Author,
		Isbn:     req.Isbn,
		Quantity: req.Quantity,
		Category: req.Category,
	}

	if err = uc.db.CreateBook(ctx, payload); err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	return resp, nil
}
