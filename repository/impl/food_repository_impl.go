package impl

import (
	"context"
	"errors"
	"strings"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/exception"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewFoodRepositoryImpl(DB *gorm.DB) repository.FoodRepository {
	return &foodRepositoryImpl{DB: DB}
}

type foodRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *foodRepositoryImpl) Insert(ctx context.Context, food entity.Food) entity.Food {
	food.ID = uuid.New()
	err := repository.DB.Create(&food).Error
	exception.PanicLogging(err)
	return food
}

func (repository *foodRepositoryImpl) FindAll(ctx context.Context, filter *model.FoodFilter) (foods []entity.Food, err error) {
	err = repository.setFilter(repository.DB, filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&foods).Error
	return
}

func (repository *foodRepositoryImpl) Count(ctx context.Context, filter *model.FoodFilter) (count int) {
	var total int64
	repository.setFilter(repository.DB, filter).Model(&entity.Food{}).Count(&total)
	count = int(total)
	return
}

func (repository *foodRepositoryImpl) FindById(ctx context.Context, id string) (entity.Food, error) {
	var food entity.Food
	result := repository.DB.WithContext(ctx).Unscoped().Where("deleted_at is null").Where("id = ?", id).First(&food)
	if result.RowsAffected == 0 {
		return entity.Food{}, errors.New("food not found")
	}
	return food, nil
}

func (repository *foodRepositoryImpl) Update(ctx context.Context, food entity.Food, id string) (entity.Food, error) {
	result := repository.DB.WithContext(ctx).Unscoped().Where("deleted_at is null").Where("id = ?", id).Updates(&food)
	if result.RowsAffected == 0 {
		return entity.Food{}, errors.New("food not found")
	}
	return food, nil
}

func (repository *foodRepositoryImpl) setFilter(db *gorm.DB, filter *model.FoodFilter) *gorm.DB {
	if *filter.ID != "" {
		db = db.Where("id = ?", *filter.ID)
	}
	if filter.Search != "" {
		db = db.Where("id ILIKE '%%' || ? || '%%' OR name ILIKE '%%' || ? || '%%'", filter.Search, filter.Search)
	}

	return db
}
