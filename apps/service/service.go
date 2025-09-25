package service

import (
	"context"
	"digitalbooklending/apps/domain"
)

type IAuth interface {
	SignUp(ctx context.Context, req domain.SignUpRequest) (err error)
	SignIn(ctx context.Context, req domain.SignInRequest) (resp domain.SignInResponse, err error)
}

type IBook interface {
	CreateBook(ctx context.Context, req domain.RequestBook) (resp domain.Book, err error)
	GetListBooks(ctx context.Context, limit int, page int) (resp domain.ResponseBook, err error)
	FirstBookById(ctx context.Context, id uint) (book domain.Book, err error)
	UpdateBook(ctx context.Context, id uint, req domain.RequestBook) (resp domain.Book, err error)
	DeleteBook(ctx context.Context, id uint) (err error)
}
