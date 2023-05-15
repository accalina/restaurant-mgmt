package impl

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/accalina/restaurant-mgmt/service"
)

func NewFoodServiceImpl(foodRepository *repository.FoodRepository) service.FoodServie {
	return &foodServiceImpl{FoodRepository: *foodRepository}
}

type foodServiceImpl struct {
	repository.FoodRepository
}

func (service *foodServiceImpl) Create(ctx context.Context, foodModel model.FoodCreteOrUpdateModel) entity.Food {
	food := entity.Food{
		Field: foodModel.Field,
	}
	return service.FoodRepository.Insert(ctx, food)
}
