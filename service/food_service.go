package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/model"
)

type FoodService interface {
	FindAll(ctx context.Context, queryParams *model.FoodFilter) ([]model.FoodModel, error)
	FindById(ctx context.Context, id string) (model.FoodModel, error)
	Create(ctx context.Context, model model.FoodCreteOrUpdateModel) model.FoodCreteOrUpdateModel
	Update(ctx context.Context, foodModel model.FoodCreteOrUpdateModel, id string) model.FoodCreteOrUpdateModel
	Delete(ctx context.Context, id string) bool
}
