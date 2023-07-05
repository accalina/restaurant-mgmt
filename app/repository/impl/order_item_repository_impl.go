package impl

import (
	"context"
	"strings"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/app/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderItemRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderItemRepositoryImpl(DB *gorm.DB) repository.OrderItemRepository {
	return &orderItemRepositoryImpl{DB}
}

func (r *orderItemRepositoryImpl) FetchAll(ctx context.Context, filter *model.OrderItemFilter) (orderItems []entity.OrderItem, err error) {
	err = r.setFilter(r.DB, filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&orderItems).Error
	return
}

func (r *orderItemRepositoryImpl) Count(ctx context.Context, filter *model.OrderItemFilter) (count int) {
	var total int64
	r.setFilter(r.DB, filter).Model(&entity.OrderItem{}).Count(&total)
	count = int(total)
	return
}

func (r *orderItemRepositoryImpl) Find(ctx context.Context, filter *model.OrderItemFilter) (result entity.OrderItem, err error) {
	err = r.setFilter(r.DB, filter).First(&result).Error
	return
}

func (r *orderItemRepositoryImpl) Save(tx *gorm.DB, orderItem *entity.OrderItem) (*entity.OrderItem, error) {
	if orderItem.ID == "" {
		orderItem.ID = uuid.NewString()
		err := tx.Create(orderItem).Error
		return orderItem, err
	}
	err := tx.Save(orderItem).Error
	return orderItem, err
}

func (r *orderItemRepositoryImpl) setFilter(db *gorm.DB, filter *model.OrderItemFilter) *gorm.DB {
	if *filter.ID != "" {
		db = db.Where("id = ?", *filter.ID)
	}

	if *filter.FoodID != "" {
		db = db.Where("food_id = ?", *filter.FoodID)
	}

	if filter.Search != "" {
		db = db.Where("id ILIKE '%%' || ? || '%%' OR name ILIKE '%%' || ? || '%%' OR category ILIKE '%%' || ? || '%%'", filter.Search, filter.Search, filter.Search)
	}

	return db
}
