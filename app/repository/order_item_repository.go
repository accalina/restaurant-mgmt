package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	FetchAll(ctx context.Context, filter *model.OrderItemFilter) ([]entity.OrderItem, error)
	Count(ctx context.Context, filter *model.OrderItemFilter) int
	Find(ctx context.Context, filter *model.OrderItemFilter) (entity.OrderItem, error)
	Save(tx *gorm.DB, data *entity.OrderItem) (*entity.OrderItem, error)
	Delete(tx *gorm.DB, data *entity.OrderItem) (error)
}
