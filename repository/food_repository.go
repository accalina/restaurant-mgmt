package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
)

type FoodRepository interface {
	FindAll(ctx context.Context) []entity.Food
	FindById(ctx context.Context, id string) (entity.Food, error)
	Insert(ctx context.Context, food entity.Food) entity.Food
	Delete(ctx context.Context, food entity.Food, id string) error
}
