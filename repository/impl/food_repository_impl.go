package impl

import (
	"context"

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

func NewFoodRepositoryImpl2(DB *gorm.DB) repository.FoodRepository {
	return &foodRepositoryImpl{DB: DB}
}
