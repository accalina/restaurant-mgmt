package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
)

type OrderRepository interface {
	FetchAll(ctx context.Context, filter *model.OrderFilter) ([]entity.Order, error)
	Count(ctx context.Context, filter *model.OrderFilter) int
	Find(ctx context.Context, filter *model.OrderFilter) (entity.Order, error)
	Save(ctx context.Context, data *entity.Order) (*entity.Order, error)
}
