package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/model"
)

type UserService interface {
	Register(ctx context.Context, model model.UserCreateModel) model.UserCreateModel
	Login(ctx context.Context, username string, password string) (model.UserModel, error)
	FindAll(ctx context.Context) []model.UserModel
	FindById(ctx context.Context, id string) (model.UserModel, error)
	Update(ctx context.Context, model model.UserUpdateModel, id string) model.UserUpdateModel
	Delete(ctx context.Context, id string) bool
	Promote(ctx context.Context, id string) bool
	Demote(ctx context.Context, id string) bool
}
