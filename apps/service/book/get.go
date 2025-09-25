package book

import (
	"context"
	"digitalbooklending/apps/domain"
	"digitalbooklending/apps/models"

	"github.com/vizucode/gokit/logger"
)

func (uc *book) GetListBooks(ctx context.Context, limit int, page int) (resp domain.ResponseBook, err error) {

	books, err := uc.db.GetListBooks(ctx, models.Filter{
		Limit: limit,
		Page:  page,
	})
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	for _, book := range books {
		resp.Books = append(resp.Books, domain.Book{
			Id:       book.Id,
			Title:    book.Title,
			Author:   book.Author,
			Isbn:     book.Isbn,
			Quantity: book.Quantity,
			Category: book.Category,
		})
	}

	booksCount, err := uc.db.GetListBooks(ctx, models.Filter{
		Limit: 0,
		Page:  0,
	})
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	resp.TotalPage = int(int64(len(booksCount)) / int64(limit))
	if int64(len(booksCount))%int64(limit) != 0 {
		resp.TotalPage++
	}
	resp.Limit = limit
	resp.Page = page
	resp.CurrentPage = page

	return resp, nil
}

func (uc *book) FirstBookById(ctx context.Context, id uint) (book domain.Book, err error) {
	bookModel, err := uc.db.FirstBookById(ctx, id)
	if err != nil {
		logger.Log.Error(ctx, err)
		return book, err
	}

	book = domain.Book{
		Id:       bookModel.Id,
		Title:    bookModel.Title,
		Author:   bookModel.Author,
		Isbn:     bookModel.Isbn,
		Quantity: bookModel.Quantity,
		Category: bookModel.Category,
	}
	return book, nil
}
