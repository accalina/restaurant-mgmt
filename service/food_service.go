package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
)

type FoodServie interface {
	Create(ctx context.Context, model model.FoodCreteOrUpdateModel) entity.Food
}
