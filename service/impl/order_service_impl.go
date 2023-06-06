package impl

import (
	"context"
	"time"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/accalina/restaurant-mgmt/service"
)

type orderServiceImpl struct {
	repository.OrderRepository
}

func NewOrderServiceImpl(r *repository.OrderRepository) service.OrderService {
	return &orderServiceImpl{OrderRepository: *r}
}

func (s *orderServiceImpl) GetAllOrder(ctx context.Context, filter *model.OrderFilter) (result []model.OrderResponse, meta model.Meta, err error) {
	orders, err := s.OrderRepository.FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.OrderRepository.Count(ctx, filter)
	meta = model.NewMeta(filter.Page, filter.Limit, count)
	for _, order := range orders {
		result = append(result, model.OrderResponse{
			ID:        order.ID,
			OrderDate: order.OrderDate,
			TableID:   order.TableID,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			DeletedAt: order.DeletedAt,
		})
	}
	return
}

func (s *orderServiceImpl) GetDetailOrder(ctx context.Context, id string) (result model.OrderResponse, err error) {
	var data entity.Order
	filter := model.OrderFilter{ID: &id}
	data, err = s.OrderRepository.Find(ctx, &filter)
	if err != nil {
		return
	}

	result.ID = data.ID
	result.OrderDate = data.OrderDate
	result.TableID = data.TableID
	result.CreatedAt = data.CreatedAt
	result.UpdatedAt = data.UpdatedAt
	result.DeletedAt = data.DeletedAt

	return
}

func (s *orderServiceImpl) CreateOrder(ctx context.Context, orderModel model.OrderCreateOrUpdateModel) (*entity.Order, error) {
	order := entity.Order{TableID: orderModel.TableID}
	return s.OrderRepository.Save(ctx, &order)
}

func (s *orderServiceImpl) UpdateOrder(ctx context.Context, orderModel model.OrderCreateOrUpdateModel) (*entity.Order, error) {
	var order entity.Order
	filter := model.OrderFilter{ID: &orderModel.ID}
	order, err := s.OrderRepository.Find(ctx, &filter)
	if err != nil {
		return &entity.Order{}, err
	}

	order.TableID = orderModel.TableID

	return s.OrderRepository.Save(ctx, &order)
}

func (s *orderServiceImpl) DeleteOrder(ctx context.Context, id string) (err error) {
	filter := model.OrderFilter{ID: &id}
	order, err := s.OrderRepository.Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	order.DeletedAt = &deleted_at
	_, err = s.OrderRepository.Save(ctx, &order)
	return
}
