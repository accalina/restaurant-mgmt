package impl

import (
	"context"
	"strings"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepositoryImpl(DB *gorm.DB) repository.OrderRepository {
	return &orderRepositoryImpl{DB}
}

func (r *orderRepositoryImpl) FetchAll(ctx context.Context, filter *model.OrderFilter) (orders []entity.Order, err error) {
	err = r.setFilter(r.DB, filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&orders).Error
	return
}

func (r *orderRepositoryImpl) Count(ctx context.Context, filter *model.OrderFilter) (count int) {
	var total int64
	r.setFilter(r.DB, filter).Model(&entity.Order{}).Count(&total)
	count = int(total)
	return
}

func (r *orderRepositoryImpl) Find(ctx context.Context, filter *model.OrderFilter) (result entity.Order, err error) {
	err = r.setFilter(r.DB, filter).First(&result).Error
	return
}

func (r *orderRepositoryImpl) Save(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	if order.ID == "" {
		order.ID = uuid.NewString()
		err := r.DB.Create(order).Error
		return order, err
	}
	err := r.DB.Save(order).Error
	return order, err
}

func (r *orderRepositoryImpl) setFilter(db *gorm.DB, filter *model.OrderFilter) *gorm.DB {
	if *filter.ID != "" {
		db = db.Where("id = ?", *filter.ID)
	}

	if filter.Search != "" {
		db = db.Where("id ILIKE '%%' || ? || '%%' OR name ILIKE '%%' || ? || '%%' OR category ILIKE '%%' || ? || '%%'", filter.Search, filter.Search, filter.Search)
	}

	return db
}
