package service

import (
	"context"
	"digitalbooklending/apps/domain"
)

type IAuth interface {
	SignUp(ctx context.Context, req domain.SignUpRequest) (err error)
	SignIn(ctx context.Context, req domain.SignInRequest) (resp domain.SignInResponse, err error)
}
