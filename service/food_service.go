package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/model"
)

type FoodService interface {
	FindAll(ctx context.Context) []model.FoodModel
	FindById(ctx context.Context, id string) (model.FoodModel, error)
	Create(ctx context.Context, model model.FoodCreteOrUpdateModel) model.FoodCreteOrUpdateModel
}
