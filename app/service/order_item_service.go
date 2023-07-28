package service

import (
	"context"
	"errors"
	"time"

	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/pkg/shared/repository"
	"github.com/accalina/restaurant-mgmt/platform/cache"
	"github.com/go-redis/redis/v8"
)

type OrderItemService interface {
	GetAllOrderItem(ctx context.Context, filter *model.OrderItemFilter) (result []model.OrderItemResponse, meta model.Meta, err error)
	GetDetailOrderItem(ctx context.Context, id string) (result *model.OrderItemResponse, err error)
	CreateOrderItem(ctx context.Context, data model.OrderItemCreateOrUpdateModel) (*model.OrderItemResponse, error)
	UpdateOrderItem(ctx context.Context, data model.OrderItemChangeQtyModel) (*model.OrderItemResponse, error)
	DeleteOrderItem(ctx context.Context, id string) error
}

type orderItemServiceImpl struct {
	repoSQL repository.RepoSQL
	Redis   *redis.Client
}

func NewOrderItemServiceImpl() OrderItemService {
	return &orderItemServiceImpl{repoSQL: repository.GetSharedRepoSQL(), Redis: cache.GetRedisCache()}
}

func (s *orderItemServiceImpl) GetAllOrderItem(ctx context.Context, filter *model.OrderItemFilter) (result []model.OrderItemResponse, meta model.Meta, err error) {
	orderItems, err := s.repoSQL.OrderItemRepo().FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.repoSQL.OrderItemRepo().Count(ctx, filter)
	meta = model.NewMeta(filter.Page, filter.Limit, count)
	for _, orderItem := range orderItems {
		result = append(result, model.OrderItemResponse{
			ID:        orderItem.ID,
			Qty:       orderItem.Qty,
			FoodId:    orderItem.FoodId,
			OrderId:   orderItem.OrderId,
			CreatedAt: orderItem.CreatedAt,
			UpdatedAt: orderItem.UpdatedAt,
			DeletedAt: orderItem.DeletedAt,
		})
	}
	return
}

func (s *orderItemServiceImpl) GetDetailOrderItem(ctx context.Context, id string) (result *model.OrderItemResponse, err error) {
	filter := model.NewOrderItemFilter()
	filter.ID = &id
	orderItem, err := s.repoSQL.OrderItemRepo().Find(ctx, filter)
	if err != nil {
		return
	}

	result = &model.OrderItemResponse{
		ID:        orderItem.ID,
		Qty:       orderItem.Qty,
		FoodId:    orderItem.FoodId,
		OrderId:   orderItem.OrderId,
		CreatedAt: orderItem.CreatedAt,
		UpdatedAt: orderItem.UpdatedAt,
		DeletedAt: orderItem.DeletedAt,
	}

	return
}

func (s *orderItemServiceImpl) CreateOrderItem(ctx context.Context, orderItemModel model.OrderItemCreateOrUpdateModel) (result *model.OrderItemResponse, err error) {
	foodFilter := model.FoodFilter{ID: &orderItemModel.FoodId}
	food, err := s.repoSQL.FoodRepo().Find(ctx, &foodFilter)
	if err != nil {
		return
	}
	if food.Qty < int32(orderItemModel.Qty) {
		return result, errors.New("Insufficient food quoantity")
	}

	// create or update order item
	filter := model.NewOrderItemFilter()
	filter.FoodID = &orderItemModel.FoodId
	filter.OrderID = &orderItemModel.OrderId
	orderItem, err := s.repoSQL.OrderItemRepo().Find(ctx, filter)
	if err != nil {
		orderItem.Qty = orderItemModel.Qty
		orderItem.FoodId = orderItemModel.FoodId
		orderItem.OrderId = orderItemModel.OrderId
	} else {
		orderItem.Qty += orderItemModel.Qty
	}

	tx := s.repoSQL.GetDB().Begin()
	if _, err = s.repoSQL.OrderItemRepo().Save(tx, &orderItem); err != nil {
		tx.Rollback()
		return
	}

	// Decrease food stockctx
	food.Qty -= int32(orderItemModel.Qty)
	if _, err = s.repoSQL.FoodRepo().Save(tx, food); err != nil {
		tx.Rollback()
		return
	}

	result = &model.OrderItemResponse{
		ID:        orderItem.ID,
		Qty:       orderItem.Qty,
		FoodId:    orderItem.FoodId,
		OrderId:   orderItem.OrderId,
		CreatedAt: orderItem.CreatedAt,
		UpdatedAt: orderItem.UpdatedAt,
		DeletedAt: orderItem.DeletedAt,
	}

	tx.Commit()
	return
}

func (s *orderItemServiceImpl) UpdateOrderItem(ctx context.Context, orderItemModel model.OrderItemChangeQtyModel) (result *model.OrderItemResponse, err error) {
	filter := model.NewOrderItemFilter("Food")
	filter.ID = &orderItemModel.ID
	orderItem, err := s.repoSQL.OrderItemRepo().Find(ctx, filter)
	if err != nil {
		return
	}

	return_qty := orderItem.Qty - orderItemModel.Qty
	food := orderItem.Food
	food.Qty += int32(return_qty)
	if food.Qty < 0 {
		return result, errors.New("Insufficient food stock")
	}

	tx := s.repoSQL.GetDB().Begin()
	// return / decrease food stock
	if _, err = s.repoSQL.FoodRepo().Save(s.repoSQL.GetDB(), &food); err != nil {
		tx.Rollback()
		return
	}

	// update order item quantity
	orderItem.Qty = orderItemModel.Qty
	if _, err = s.repoSQL.OrderItemRepo().Save(tx, &orderItem); err != nil {
		tx.Rollback()
		return
	}

	// delete order item
	if orderItemModel.Qty == 0 {
		if err = s.repoSQL.OrderItemRepo().Delete(tx, &orderItem); err != nil {
			tx.Rollback()
			return
		}
	}

	result = &model.OrderItemResponse{
		ID:        orderItem.ID,
		Qty:       orderItem.Qty,
		FoodId:    orderItem.FoodId,
		OrderId:   orderItem.OrderId,
		CreatedAt: orderItem.CreatedAt,
		UpdatedAt: orderItem.UpdatedAt,
		DeletedAt: orderItem.DeletedAt,
	}

	tx.Commit()
	return
}

func (s *orderItemServiceImpl) DeleteOrderItem(ctx context.Context, id string) (err error) {
	filter := model.OrderItemFilter{ID: &id}
	orderItem, err := s.repoSQL.OrderItemRepo().Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	orderItem.DeletedAt = &deleted_at
	_, err = s.repoSQL.OrderItemRepo().Save(s.repoSQL.GetDB(), &orderItem)
	return
}
