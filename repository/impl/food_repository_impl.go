package impl

import (
	"context"
	"errors"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/exception"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewFoodRepositoryImpl(DB *gorm.DB) repository.FoodRepository {
	return &foodRepositoryImpl{DB: DB}
}

type foodRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *foodRepositoryImpl) Insert(ctx context.Context, food entity.Food) entity.Food {
	food.Id = uuid.New()
	err := repository.DB.Create(&food).Error
	exception.PanicLogging(err)
	return food
}

func (repository *foodRepositoryImpl) FindAll(ctx context.Context) []entity.Food {
	var foods []entity.Food
	repository.DB.WithContext(ctx).Unscoped().Where("deleted_at is null").Find(&foods)
	return foods
}

func (repository *foodRepositoryImpl) FindById(ctx context.Context, id string) (entity.Food, error) {
	var food entity.Food
	result := repository.DB.WithContext(ctx).Unscoped().Where("deleted_at is null").Where("id = ?", id).First(&food)
	if result.RowsAffected == 0 {
		return entity.Food{}, errors.New("food not found")
	}
	return food, nil
}

func (repository *foodRepositoryImpl) Delete(ctx context.Context, food entity.Food, id string) error {
	result := repository.DB.WithContext(ctx).Unscoped().Where("deleted_at is null").Where("id = ?", id).Updates(&food)
	if result.RowsAffected == 0 {
		return errors.New("food not found")
	}
	return nil
}
