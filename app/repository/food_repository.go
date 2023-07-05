package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"gorm.io/gorm"
)

type FoodRepository interface {
	FetchAll(ctx context.Context, filter *model.FoodFilter) ([]entity.Food, error)
	Count(ctx context.Context, filter *model.FoodFilter) int
	Find(ctx context.Context, filter *model.FoodFilter) (entity.Food, error)
	Save(tx *gorm.DB, data *entity.Food) (*entity.Food, error)
}
