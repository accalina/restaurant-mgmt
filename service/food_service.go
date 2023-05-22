package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
)

type FoodService interface {
	Create(ctx context.Context, model model.FoodCreteOrUpdateModel) entity.Food
}
