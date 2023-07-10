package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/model"
)

type OrderService interface {
	GetAllOrder(ctx context.Context, filter *model.OrderFilter) (result []model.OrderResponse, meta model.Meta, err error)
	GetDetailOrder(ctx context.Context, id string) (result *model.OrderResponse, err error)
	CreateOrder(ctx context.Context, data model.OrderCreateOrUpdateModel) (*model.OrderResponse, error)
	UpdateOrder(ctx context.Context, data model.OrderCreateOrUpdateModel) (*model.OrderResponse, error)
	DoneOrder(ctx context.Context, data model.OrderDoneModel) (*model.OrderResponse, error)
	DeleteOrder(ctx context.Context, id string) error
}
