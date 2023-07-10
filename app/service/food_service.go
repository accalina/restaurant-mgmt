package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/model"
)

type FoodService interface {
	GetAllFood(ctx context.Context, filter *model.FoodFilter) (result []model.FoodResponse, meta model.Meta, err error)
	GetDetailFood(ctx context.Context, id string) (result *model.FoodResponse, err error)
	CreateFood(ctx context.Context, data model.FoodCreateOrUpdateModel) (*model.FoodResponse, error)
	UpdateFood(ctx context.Context, data model.FoodCreateOrUpdateModel) (*model.FoodResponse, error)
	DeleteFood(ctx context.Context, id string) error
}
