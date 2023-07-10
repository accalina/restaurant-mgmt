package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FetchAll(ctx context.Context, filter *model.OrderFilter) ([]entity.Order, error)
	Count(ctx context.Context, filter *model.OrderFilter) int
	Find(ctx context.Context, filter *model.OrderFilter) (*entity.Order, error)
	Save(tx *gorm.DB, data *entity.Order) (*entity.Order, error)
}
