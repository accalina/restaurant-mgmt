package impl

import (
	"context"
	"time"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/app/service"
	"github.com/accalina/restaurant-mgmt/pkg/configuration"
	"github.com/accalina/restaurant-mgmt/pkg/shared/repository"
	"github.com/go-redis/redis/v8"
)

type orderServiceImpl struct {
	repoSQL repository.RepoSQL
	Redis   *redis.Client
}

func NewOrderServiceImpl() service.OrderService {
	return &orderServiceImpl{repoSQL: repository.GetSharedRepoSQL(), Redis: configuration.GetRedisCache()}
}

func (s *orderServiceImpl) GetAllOrder(ctx context.Context, filter *model.OrderFilter) (result []model.OrderResponse, meta model.Meta, err error) {
	orders, err := s.repoSQL.OrderRepo().FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.repoSQL.OrderRepo().Count(ctx, filter)
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
	data, err = s.repoSQL.OrderRepo().Find(ctx, &filter)
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
	order := entity.Order{
		Name:    orderModel.Name,
		TableID: orderModel.TableID,
	}
	return s.repoSQL.OrderRepo().Save(ctx, &order)
}

func (s *orderServiceImpl) UpdateOrder(ctx context.Context, orderModel model.OrderCreateOrUpdateModel) (*entity.Order, error) {
	var order entity.Order
	filter := model.OrderFilter{ID: &orderModel.ID}
	order, err := s.repoSQL.OrderRepo().Find(ctx, &filter)
	if err != nil {
		return &entity.Order{}, err
	}

	order.TableID = orderModel.TableID

	return s.repoSQL.OrderRepo().Save(ctx, &order)
}

func (s *orderServiceImpl) DeleteOrder(ctx context.Context, id string) (err error) {
	filter := model.OrderFilter{ID: &id}
	order, err := s.repoSQL.OrderRepo().Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	order.DeletedAt = &deleted_at
	_, err = s.repoSQL.OrderRepo().Save(ctx, &order)
	return
}
