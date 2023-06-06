package impl

import (
	"context"
	"time"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/accalina/restaurant-mgmt/service"
)

type orderItemServiceImpl struct {
	repository.OrderItemRepository
}

func NewOrderItemServiceImpl(r *repository.OrderItemRepository) service.OrderItemService {
	return &orderItemServiceImpl{OrderItemRepository: *r}
}

func (s *orderItemServiceImpl) GetAllOrderItem(ctx context.Context, filter *model.OrderItemFilter) (result []model.OrderItemResponse, meta model.Meta, err error) {
	orderItems, err := s.OrderItemRepository.FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.OrderItemRepository.Count(ctx, filter)
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
	data, err = s.OrderItemRepository.Find(ctx, &filter)
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
	orderItem := entity.OrderItem{
		Qty:     orderItemModel.Qty,
		FoodId:  orderItemModel.FoodId,
		OrderId: orderItemModel.OrderId,
	}
	return s.OrderItemRepository.Save(ctx, &orderItem)
}

func (s *orderItemServiceImpl) UpdateOrderItem(ctx context.Context, orderItemModel model.OrderItemCreateOrUpdateModel) (*entity.OrderItem, error) {
	var orderItem entity.OrderItem
	filter := model.OrderItemFilter{ID: &orderItemModel.ID}
	orderItem, err := s.OrderItemRepository.Find(ctx, &filter)
	if err != nil {
		return &entity.OrderItem{}, err
	}

	orderItem.Qty = orderItemModel.Qty
	orderItem.FoodId = orderItemModel.FoodId
	orderItem.OrderId = orderItemModel.OrderId

	return s.OrderItemRepository.Save(ctx, &orderItem)
}

func (s *orderItemServiceImpl) DeleteOrderItem(ctx context.Context, id string) (err error) {
	filter := model.OrderItemFilter{ID: &id}
	orderItem, err := s.OrderItemRepository.Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	orderItem.DeletedAt = &deleted_at
	_, err = s.OrderItemRepository.Save(ctx, &orderItem)
	return
}
