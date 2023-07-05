package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/entity"
)

type UserRepository interface {
	Register(ctx context.Context, user entity.User) entity.User
	Login(ctx context.Context, username string, password string) (entity.User, error)
	FindAll(ctx context.Context) []entity.User
	FindById(ctx context.Context, id string) (entity.User, error)
	Update(ctx context.Context, user entity.User, id string) error
}
