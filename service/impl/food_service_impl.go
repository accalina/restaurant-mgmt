package impl

import (
	"context"
	"errors"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/accalina/restaurant-mgmt/service"
)

func NewFoodServiceImpl(foodRepository *repository.FoodRepository) service.FoodService {
	return &foodServiceImpl{FoodRepository: *foodRepository}
}

type foodServiceImpl struct {
	repository.FoodRepository
}

func (service *foodServiceImpl) Create(ctx context.Context, foodModel model.FoodCreteOrUpdateModel) model.FoodCreteOrUpdateModel {
	food := entity.Food{
		Name:  foodModel.Name,
		Price: foodModel.Price,
		Qty:   foodModel.Qty,
	}
	service.FoodRepository.Insert(ctx, food)
	return foodModel
}

func (service *foodServiceImpl) FindAll(ctx context.Context) (response []model.FoodModel) {
	foods := service.FoodRepository.FindAll(ctx)
	for _, food := range foods {
		response = append(response, model.FoodModel{
			Id:        food.Id,
			Name:      food.Name,
			Price:     food.Price,
			Qty:       food.Qty,
			CreatedAt: food.CreatedAt,
			UpdatedAt: food.UpdatedAt,
		})
	}
	if len(foods) == 0 {
		return []model.FoodModel{}
	}
	return response
}

func (service *foodServiceImpl) FindById(ctx context.Context, id string) (model.FoodModel, error) {
	food, err := service.FoodRepository.FindById(ctx, id)
	if err != nil {
		return model.FoodModel{}, errors.New("food not found")
	}

	return model.FoodModel{
		Id:        food.Id,
		Name:      food.Name,
		Price:     food.Price,
		Qty:       food.Qty,
		CreatedAt: food.CreatedAt,
		UpdatedAt: food.UpdatedAt,
	}, nil
}
