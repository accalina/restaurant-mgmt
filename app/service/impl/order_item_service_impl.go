package impl

import (
	"context"
	"errors"
	"time"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/app/service"
	"github.com/accalina/restaurant-mgmt/pkg/configuration"
	"github.com/accalina/restaurant-mgmt/pkg/shared/repository"
	"github.com/go-redis/redis/v8"
)

type orderItemServiceImpl struct {
	repoSQL repository.RepoSQL
	Redis   *redis.Client
}

func NewOrderItemServiceImpl() service.OrderItemService {
	return &orderItemServiceImpl{repoSQL: repository.GetSharedRepoSQL(), Redis: configuration.GetRedisCache()}
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

func (s *orderItemServiceImpl) GetDetailOrderItem(ctx context.Context, id string) (result model.OrderItemResponse, err error) {
	var data entity.OrderItem
	filter := model.OrderItemFilter{ID: &id}
	data, err = s.repoSQL.OrderItemRepo().Find(ctx, &filter)
	if err != nil {
		return
	}

	result.ID = data.ID
	result.Qty = data.Qty
	result.FoodId = data.FoodId
	result.OrderId = data.OrderId
	result.CreatedAt = data.CreatedAt
	result.UpdatedAt = data.UpdatedAt
	result.DeletedAt = data.DeletedAt

	return
}

func (s *orderItemServiceImpl) CreateOrderItem(ctx context.Context, orderItemModel model.OrderItemCreateOrUpdateModel) (*entity.OrderItem, error) {
	// sid = transaction
	
	// check food quoantity
	var food entity.Food
	foodFilter := model.FoodFilter{ID: &orderItemModel.FoodId}
	food, err := s.repoSQL.FoodRepo().Find(ctx, &foodFilter)
	if err != nil {
		return &entity.OrderItem{}, err
	}
	if food.Qty < int32(orderItemModel.Qty) {
		return &entity.OrderItem{}, errors.New("Not enough Food Quoantity")
	}

	var orderItem entity.OrderItem
	filter := model.NewOrderItemFilter()
	filter.FoodID = &orderItemModel.FoodId
	orderItem, err = s.repoSQL.OrderItemRepo().Find(ctx, filter)
	if err != nil {
		orderItem.Qty = orderItemModel.Qty
		orderItem.FoodId = orderItemModel.FoodId
		orderItem.OrderId = orderItemModel.OrderId
	} else {
		orderItem.Qty += orderItemModel.Qty
	}

	tx := s.repoSQL.GetDB().Begin()
	orderItemInstance, err := s.repoSQL.OrderItemRepo().Save(tx, &orderItem)
	if err != nil{
		tx.Rollback()
		return &entity.OrderItem{}, err
	}
	
	// Decrease food stockctx
	food.Qty -= int32(orderItemModel.Qty)
	_, err = s.repoSQL.FoodRepo().Save(tx, &food)
	if err != nil{
		tx.Rollback()
		return &entity.OrderItem{}, err
	}

	tx.Commit()

	return orderItemInstance, err
}

func (s *orderItemServiceImpl) UpdateOrderItem(ctx context.Context, orderItemModel model.OrderItemCreateOrUpdateModel) (*entity.OrderItem, error) {
	var orderItem entity.OrderItem
	filter := model.OrderItemFilter{ID: &orderItemModel.ID}
	orderItem, err := s.repoSQL.OrderItemRepo().Find(ctx, &filter)
	if err != nil {
		return &entity.OrderItem{}, err
	}

	orderItem.Qty = orderItemModel.Qty
	orderItem.FoodId = orderItemModel.FoodId
	orderItem.OrderId = orderItemModel.OrderId

	return s.repoSQL.OrderItemRepo().Save(s.repoSQL.GetDB(), &orderItem)
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
