package repositories

import (
	"context"
	"digitalbooklending/apps/models"
)

type IDatabase interface {
	CreateUser(ctx context.Context, user models.Users) (err error)
	FirstUserById(ctx context.Context, id string) (user models.Users, err error)
	FirstUserByEmail(ctx context.Context, email string) (user models.Users, err error)

	CreateBook(ctx context.Context, book models.Book) (err error)
	FirstBookById(ctx context.Context, id uint) (book models.Book, err error)
	GetListBooks(ctx context.Context, filter models.Filter) (books []models.Book, err error)
	DeleteBook(ctx context.Context, id uint) (err error)
	UpdateBook(ctx context.Context, id uint, book models.Book) (err error)

	CreateLendingRecord(ctx context.Context, lendingRecord models.LendingRecord) (err error)
	FirstLendingRecordByUserId(ctx context.Context, userId uint) (lendingRecord models.LendingRecord, err error)
	CountUserLendingRecords(ctx context.Context, userId uint) (count int64, err error)
	UpdateLendingRecord(ctx context.Context, lendingRecord models.LendingRecord) (err error)

	CheckLimitBorrow(ctx context.Context, userId uint) (err error)
}
