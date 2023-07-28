package repository

import (
	"context"
	"strings"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FoodRepository interface {
	FetchAll(ctx context.Context, filter *model.FoodFilter) ([]entity.Food, error)
	Count(ctx context.Context, filter *model.FoodFilter) int
	Find(ctx context.Context, filter *model.FoodFilter) (*entity.Food, error)
	Save(tx *gorm.DB, data *entity.Food) (*entity.Food, error)
}

type foodRepositoryImpl struct {
	DB *gorm.DB
}

func NewFoodRepositoryImpl(DB *gorm.DB) FoodRepository {
	return &foodRepositoryImpl{DB}
}

func (r *foodRepositoryImpl) FetchAll(ctx context.Context, filter *model.FoodFilter) (foods []entity.Food, err error) {
	err = r.setFilter(r.DB, filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&foods).Error
	return
}

func (r *foodRepositoryImpl) Count(ctx context.Context, filter *model.FoodFilter) (count int) {
	var total int64
	r.setFilter(r.DB, filter).Model(&entity.Food{}).Count(&total)
	count = int(total)
	return
}

func (r *foodRepositoryImpl) Find(ctx context.Context, filter *model.FoodFilter) (result *entity.Food, err error) {
	err = r.setFilter(r.DB, filter).First(&result).Error
	return
}

func (r *foodRepositoryImpl) Save(tx *gorm.DB, food *entity.Food) (*entity.Food, error) {
	if food.ID == "" {
		food.ID = uuid.NewString()
		err := tx.Create(food).Error
		return food, err
	}
	err := tx.Save(food).Error
	return food, err
}

func (r *foodRepositoryImpl) setFilter(db *gorm.DB, filter *model.FoodFilter) *gorm.DB {
	if *filter.ID != "" {
		db = db.Where("id = ?", *filter.ID)
	}

	if filter.Search != "" {
		db = db.Where("id ILIKE '%%' || ? || '%%' OR name ILIKE '%%' || ? || '%%' OR category ILIKE '%%' || ? || '%%'", filter.Search, filter.Search, filter.Search)
	}

	for _, preload := range filter.Preloads {
		db = db.Preload(preload)
	}

	return db
}
