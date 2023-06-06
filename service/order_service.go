package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
)

type OrderService interface {
	GetAllOrder(ctx context.Context, filter *model.OrderFilter) (result []model.OrderResponse, meta model.Meta, err error)
	GetDetailOrder(ctx context.Context, id string) (result model.OrderResponse, err error)
	CreateOrder(ctx context.Context, data model.OrderCreateOrUpdateModel) (*entity.Order, error)
	UpdateOrder(ctx context.Context, data model.OrderCreateOrUpdateModel) (*entity.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}
