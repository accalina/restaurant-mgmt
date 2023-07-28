package service

import (
	"context"
	"errors"
	"time"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/pkg/shared/repository"
	"github.com/accalina/restaurant-mgmt/platform/cache"
	"github.com/go-redis/redis/v8"
)

type OrderService interface {
	GetAllOrder(ctx context.Context, filter *model.OrderFilter) (result []model.OrderResponse, meta model.Meta, err error)
	GetDetailOrder(ctx context.Context, id string) (result *model.OrderResponse, err error)
	CreateOrder(ctx context.Context, data model.OrderCreateOrUpdateModel) (*model.OrderResponse, error)
	UpdateOrder(ctx context.Context, data model.OrderCreateOrUpdateModel) (*model.OrderResponse, error)
	DoneOrder(ctx context.Context, data model.OrderDoneModel) (*model.OrderResponse, error)
	DeleteOrder(ctx context.Context, id string) error
}

type orderServiceImpl struct {
	repoSQL repository.RepoSQL
	Redis   *redis.Client
}

func NewOrderServiceImpl() OrderService {
	return &orderServiceImpl{repoSQL: repository.GetSharedRepoSQL(), Redis: cache.GetRedisCache()}
}

func (s *orderServiceImpl) GetAllOrder(ctx context.Context, filter *model.OrderFilter) (result []model.OrderResponse, meta model.Meta, err error) {
	orders, err := s.repoSQL.OrderRepo().FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.repoSQL.OrderRepo().Count(ctx, filter)
	meta = model.NewMeta(filter.Page, filter.Limit, count)
	for _, order := range orders {
		table := model.TableResponse{
			ID:             order.TableID,
			No:             order.Table.No,
			NumberOfGuests: order.Table.NumberOfGuests,
		}

		result = append(result, model.OrderResponse{
			ID:        order.ID,
			Name:      order.Name,
			OrderDate: order.OrderDate,
			Status:    string(order.Status),
			Table:     table,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			DeletedAt: order.DeletedAt,
		})
	}
	return
}

func (s *orderServiceImpl) GetDetailOrder(ctx context.Context, id string) (result *model.OrderResponse, err error) {
	filter := model.NewOrderFilter("Table", "OrderItems")
	filter.ID = &id
	order, err := s.repoSQL.OrderRepo().Find(ctx, filter)
	if err != nil {
		return
	}

	var orderItems []model.OrderItemResponse
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, model.OrderItemResponse{
			ID:        orderItem.ID,
			Qty:       orderItem.Qty,
			FoodId:    orderItem.FoodId,
			OrderId:   orderItem.OrderId,
			CreatedAt: orderItem.CreatedAt,
			UpdatedAt: orderItem.UpdatedAt,
			DeletedAt: orderItem.DeletedAt,
		})
	}

	result = &model.OrderResponse{
		ID:        order.ID,
		Name:      order.Name,
		OrderDate: order.OrderDate,
		Status:    string(order.Status),
		Table: model.TableResponse{
			ID:             order.TableID,
			No:             order.Table.No,
			NumberOfGuests: order.Table.NumberOfGuests,
		},
		OrderItems: orderItems,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		DeletedAt:  order.DeletedAt,
	}

	return
}

func (s *orderServiceImpl) CreateOrder(ctx context.Context, orderModel model.OrderCreateOrUpdateModel) (result *model.OrderResponse, err error) {
	isAvailable := "true"
	filter := model.TableFilter{ID: &orderModel.TableID, IsAvailable: &isAvailable}
	table, err := s.repoSQL.TableRepo().Find(ctx, &filter)
	if err != nil {
		err = errors.New("Table unavailable")
		return
	}

	order := &entity.Order{
		Name:    orderModel.Name,
		TableID: orderModel.TableID,
	}

	tx := s.repoSQL.GetDB().Begin()
	order, err = s.repoSQL.OrderRepo().Save(tx, order)
	if err != nil {
		tx.Rollback()
		return &model.OrderResponse{}, err
	}

	table.IsAvailable = false
	if _, err = s.repoSQL.TableRepo().Save(tx, table); err != nil {
		tx.Rollback()
		return
	}

	result = &model.OrderResponse{
		ID:        order.ID,
		Name:      order.Name,
		OrderDate: order.OrderDate,
		Status:    string(order.Status),
		Table: model.TableResponse{
			ID:             order.TableID,
			No:             order.Table.No,
			NumberOfGuests: order.Table.NumberOfGuests,
		},
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		DeletedAt: order.DeletedAt,
	}

	tx.Commit()

	return
}

func (s *orderServiceImpl) UpdateOrder(ctx context.Context, orderModel model.OrderCreateOrUpdateModel) (result *model.OrderResponse, err error) {
	filter := model.NewOrderFilter()
	filter.ID = &orderModel.ID
	order, err := s.repoSQL.OrderRepo().Find(ctx, filter)
	if err != nil {
		return
	}

	order.Name = orderModel.Name
	order.TableID = orderModel.TableID
	order, err = s.repoSQL.OrderRepo().Save(s.repoSQL.GetDB(), order)
	if err != nil {
		return
	}
	result = &model.OrderResponse{
		ID:        order.ID,
		Name:      order.Name,
		OrderDate: order.OrderDate,
		Status:    string(order.Status),
		Table: model.TableResponse{
			ID:             order.TableID,
			No:             order.Table.No,
			NumberOfGuests: order.Table.NumberOfGuests,
		},
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		DeletedAt: order.DeletedAt,
	}
	return
}

func (s *orderServiceImpl) DoneOrder(ctx context.Context, orderModel model.OrderDoneModel) (result *model.OrderResponse, err error) {
	filter := model.NewOrderFilter("Table")
	filter.ID = &orderModel.ID
	order, err := s.repoSQL.OrderRepo().Find(ctx, filter)
	if err != nil {
		return
	}

	tx := s.repoSQL.GetDB().Begin()
	order.Status = entity.OrderStatus(orderModel.Status)
	if _, err = s.repoSQL.OrderRepo().Save(tx, order); err != nil {
		tx.Rollback()
		return
	}

	table := &order.Table
	table.IsAvailable = true
	table, err = s.repoSQL.TableRepo().Save(tx, table)
	if err != nil {
		tx.Rollback()
		return
	}

	result = &model.OrderResponse{
		ID:        order.ID,
		Name:      order.Name,
		OrderDate: order.OrderDate,
		Status:    string(order.Status),
		Table: model.TableResponse{
			ID:             order.TableID,
			No:             order.Table.No,
			NumberOfGuests: order.Table.NumberOfGuests,
			IsAvailable:    order.Table.IsAvailable,
		},
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		DeletedAt: order.DeletedAt,
	}

	tx.Commit()

	return
}

func (s *orderServiceImpl) DeleteOrder(ctx context.Context, id string) (err error) {
	filter := model.NewOrderFilter()
	filter.ID = &id
	order, err := s.repoSQL.OrderRepo().Find(ctx, filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	order.DeletedAt = &deleted_at
	_, err = s.repoSQL.OrderRepo().Save(s.repoSQL.GetDB(), order)
	return
}
