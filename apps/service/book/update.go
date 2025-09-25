package book

import (
	"context"
	"digitalbooklending/apps/domain"
	"digitalbooklending/apps/models"

	"github.com/vizucode/gokit/logger"
)

func (uc *book) UpdateBook(ctx context.Context, id uint, req domain.RequestBook) (resp domain.Book, err error) {
	err = uc.db.UpdateBook(ctx, id, models.Book{
		Title:    req.Title,
		Author:   req.Author,
		Isbn:     req.Isbn,
		Quantity: req.Quantity,
		Category: req.Category,
	})

	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	return resp, nil
}
