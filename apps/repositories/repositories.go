package repositories

import (
	"context"
	"digitalbooklending/apps/models"
)

type IDatabase interface {
	CreateUser(ctx context.Context, user models.Users) (err error)
	FirstUserById(ctx context.Context, id string) (user models.Users, err error)
	FirstUserByEmail(ctx context.Context, email string) (user models.Users, err error)
}
