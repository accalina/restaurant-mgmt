package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/model"
)

type OrderItemService interface {
	GetAllOrderItem(ctx context.Context, filter *model.OrderItemFilter) (result []model.OrderItemResponse, meta model.Meta, err error)
	GetDetailOrderItem(ctx context.Context, id string) (result *model.OrderItemResponse, err error)
	CreateOrderItem(ctx context.Context, data model.OrderItemCreateOrUpdateModel) (*model.OrderItemResponse, error)
	UpdateOrderItem(ctx context.Context, data model.OrderItemChangeQtyModel) (*model.OrderItemResponse, error)
	DeleteOrderItem(ctx context.Context, id string) error
}
