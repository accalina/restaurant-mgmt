package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
)

type OrderItemRepository interface {
	FetchAll(ctx context.Context, filter *model.OrderItemFilter) ([]entity.OrderItem, error)
	Count(ctx context.Context, filter *model.OrderItemFilter) int
	Find(ctx context.Context, filter *model.OrderItemFilter) (entity.OrderItem, error)
	Save(ctx context.Context, data *entity.OrderItem) (*entity.OrderItem, error)
}
