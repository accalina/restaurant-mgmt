package impl

import (
	"context"
	"errors"
	"time"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/exception"
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

func (service *foodServiceImpl) FindAll(ctx context.Context, filter *model.FoodFilter) (response []model.FoodModel, err error) {
	foods, err := service.FoodRepository.FindAll(ctx, filter)
	if err != nil {
		return response, err
	}
	for _, food := range foods {

		updatedAt := time.Time{}
		updatedAt = *food.UpdatedAt
		response = append(response, model.FoodModel{
			Id:        food.Id,
			Name:      food.Name,
			Price:     food.Price,
			Qty:       food.Qty,
			CreatedAt: food.CreatedAt,
			UpdatedAt: updatedAt,
		})
	}
	return
}

func (service *foodServiceImpl) FindById(ctx context.Context, id string) (model.FoodModel, error) {
	food, err := service.FoodRepository.FindById(ctx, id)
	if err != nil {
		return model.FoodModel{}, errors.New("food not found")
	}
	updatedAt := time.Time{}
	updatedAt = *food.UpdatedAt
	return model.FoodModel{
		Id:        food.Id,
		Name:      food.Name,
		Price:     food.Price,
		Qty:       food.Qty,
		CreatedAt: food.CreatedAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (service *foodServiceImpl) Update(ctx context.Context, foodModel model.FoodCreteOrUpdateModel, id string) model.FoodCreteOrUpdateModel {
	currentTime := time.Now()
	food := entity.Food{
		Name:      foodModel.Name,
		Price:     foodModel.Price,
		Qty:       foodModel.Qty,
		UpdatedAt: &currentTime,
	}
	_, err := service.FoodRepository.Update(ctx, food, id)
	exception.PanicLogging(err)
	return foodModel
}

func (service *foodServiceImpl) Delete(ctx context.Context, id string) bool {
	currentTime := time.Now()
	food := entity.Food{
		DeletedAt: &currentTime,
	}
	_, err := service.FoodRepository.Update(ctx, food, id)
	exception.PanicLogging(err)
	return err != nil
}
