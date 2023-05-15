package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
)

type FoodRepository interface {
	Insert(ctx context.Context, food entity.Food) entity.Food
}
