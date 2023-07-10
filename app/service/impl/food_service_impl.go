package impl

import (
	"context"
	"time"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/app/service"
	"github.com/accalina/restaurant-mgmt/pkg/configuration"
	"github.com/accalina/restaurant-mgmt/pkg/shared/repository"
	"github.com/go-redis/redis/v8"
)

type foodServiceImpl struct {
	repoSQL repository.RepoSQL
	Redis   *redis.Client
}

func NewFoodServiceImpl() service.FoodService {
	return &foodServiceImpl{repoSQL: repository.GetSharedRepoSQL(), Redis: configuration.GetRedisCache()}
}

func (s *foodServiceImpl) GetAllFood(ctx context.Context, filter *model.FoodFilter) (result []model.FoodResponse, meta model.Meta, err error) {
	foods, err := s.repoSQL.FoodRepo().FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.repoSQL.FoodRepo().Count(ctx, filter)
	meta = model.NewMeta(filter.Page, filter.Limit, count)
	for _, food := range foods {
		result = append(result, model.FoodResponse{
			ID:        food.ID,
			Name:      food.Name,
			Price:     food.Price,
			Qty:       food.Qty,
			MenuID:    food.MenuID,
			CreatedAt: food.CreatedAt,
			UpdatedAt: food.UpdatedAt,
			DeletedAt: food.DeletedAt,
		})
	}
	return
}

func (s *foodServiceImpl) GetDetailFood(ctx context.Context, id string) (result *model.FoodResponse, err error) {
	filter := model.NewFoodFilter("Menu")
	filter.ID = &id
	food, err := s.repoSQL.FoodRepo().Find(ctx, filter)
	if err != nil {
		return
	}
	result = &model.FoodResponse{
		ID:        food.ID,
		Name:      food.Name,
		Price:     food.Price,
		Qty:       food.Qty,
		MenuID:    food.MenuID,
		CreatedAt: food.CreatedAt,
		UpdatedAt: food.UpdatedAt,
		DeletedAt: food.DeletedAt,
	}

	return
}

func (s *foodServiceImpl) CreateFood(ctx context.Context, foodModel model.FoodCreateOrUpdateModel) (result *model.FoodResponse, err error) {
	food := &entity.Food{
		Name:   foodModel.Name,
		Price:  foodModel.Price,
		Qty:    foodModel.Qty,
		MenuID: foodModel.MenuID,
	}

	food, err = s.repoSQL.FoodRepo().Save(s.repoSQL.GetDB(), food)
	if err != nil {
		return
	}
	result = &model.FoodResponse{
		ID:        food.ID,
		Name:      food.Name,
		Price:     food.Price,
		Qty:       food.Qty,
		MenuID:    food.MenuID,
		CreatedAt: food.CreatedAt,
		UpdatedAt: food.UpdatedAt,
		DeletedAt: food.DeletedAt,
	}
	return
}

func (s *foodServiceImpl) UpdateFood(ctx context.Context, foodModel model.FoodCreateOrUpdateModel) (result *model.FoodResponse, err error) {
	filter := model.FoodFilter{ID: &foodModel.ID}
	food, err := s.repoSQL.FoodRepo().Find(ctx, &filter)
	if err != nil {
		return
	}

	food.Name = foodModel.Name
	food.Price = foodModel.Price
	food.Qty = foodModel.Qty
	food.MenuID = foodModel.MenuID
	food, err = s.repoSQL.FoodRepo().Save(s.repoSQL.GetDB(), food)
	result = &model.FoodResponse{
		ID:        food.ID,
		Name:      food.Name,
		Price:     food.Price,
		Qty:       food.Qty,
		MenuID:    food.MenuID,
		CreatedAt: food.CreatedAt,
		UpdatedAt: food.UpdatedAt,
		DeletedAt: food.DeletedAt,
	}

	return
}

func (s *foodServiceImpl) DeleteFood(ctx context.Context, id string) (err error) {
	filter := model.FoodFilter{ID: &id}
	food, err := s.repoSQL.FoodRepo().Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	food.DeletedAt = &deleted_at
	_, err = s.repoSQL.FoodRepo().Save(s.repoSQL.GetDB(), food)
	return
}
