package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
)

type FoodRepository interface {
	FindAll(ctx context.Context, filter *model.FoodFilter) ([]entity.Food, error)
	Count(ctx context.Context, filter *model.FoodFilter) int
	FindById(ctx context.Context, id string) (entity.Food, error)
	Insert(ctx context.Context, food entity.Food) entity.Food
	Update(ctx context.Context, food entity.Food, id string) (entity.Food, error)
}
