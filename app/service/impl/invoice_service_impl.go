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

type invoiceServiceImpl struct {
	repoSQL repository.RepoSQL
	Redis   *redis.Client
}

func NewInvoiceServiceImpl() service.InvoiceService {
	return &invoiceServiceImpl{repoSQL: repository.GetSharedRepoSQL(), Redis: configuration.GetRedisCache()}
}

func (s *invoiceServiceImpl) GetAllInvoice(ctx context.Context, filter *model.InvoiceFilter) (result []model.InvoiceResponse, meta model.Meta, err error) {
	invoices, err := s.repoSQL.InvoiceRepo().FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.repoSQL.InvoiceRepo().Count(ctx, filter)
	meta = model.NewMeta(filter.Page, filter.Limit, count)
	for _, invoice := range invoices {
		result = append(result, model.InvoiceResponse{
			ID:             invoice.ID,
			PaymentMethod:  invoice.PaymentMethod,
			PaymentStatus:  invoice.PaymentStatus,
			PaymentDueDate: invoice.PaymentDueDate,
			Order: model.OrderResponse{
				ID:        invoice.Order.ID,
				Name:      invoice.Order.Name,
				OrderDate: invoice.Order.OrderDate,
				Status:    string(invoice.Order.Status),
				Table: model.TableResponse{
					ID:             invoice.Order.Table.ID,
					No:             invoice.Order.Table.No,
					NumberOfGuests: invoice.Order.Table.NumberOfGuests,
				},
				CreatedAt: invoice.Order.CreatedAt,
				UpdatedAt: invoice.Order.UpdatedAt,
				DeletedAt: invoice.Order.UpdatedAt,
			},
			CreatedAt: invoice.CreatedAt,
			UpdatedAt: invoice.UpdatedAt,
			DeletedAt: invoice.DeletedAt,
		})
	}
	return
}

func (s *invoiceServiceImpl) GetDetailInvoice(ctx context.Context, id string) (result model.InvoiceResponse, err error) {
	var data entity.Invoice
	filter := model.NewInvoiceFilter("Order.Table")
	filter.ID = &id
	data, err = s.repoSQL.InvoiceRepo().Find(ctx, filter)
	if err != nil {
		return
	}

	result.ID = data.ID
	result.PaymentMethod = data.PaymentMethod
	result.PaymentStatus = data.PaymentStatus
	result.PaymentDueDate = data.PaymentDueDate
	result.Order = model.OrderResponse{
		ID:        data.Order.ID,
		Name:      data.Order.Name,
		OrderDate: data.Order.OrderDate,
		Status:    string(data.Order.Status),
		Table: model.TableResponse{
			ID:             data.Order.Table.ID,
			No:             data.Order.Table.No,
			NumberOfGuests: data.Order.Table.NumberOfGuests,
		},
		CreatedAt: data.Order.CreatedAt,
		UpdatedAt: data.Order.UpdatedAt,
		DeletedAt: data.Order.UpdatedAt,
	}
	result.CreatedAt = data.CreatedAt
	result.UpdatedAt = data.UpdatedAt
	result.DeletedAt = data.DeletedAt

	return
}

func (s *invoiceServiceImpl) CreateInvoice(ctx context.Context, invoiceModel model.InvoiceCreateOrUpdateModel) (result *model.InvoiceResponse, err error) {

	invoice := &entity.Invoice{
		PaymentMethod: invoiceModel.PaymentMethod,
		PaymentStatus: invoiceModel.PaymentStatus,
		OrderId:       invoiceModel.OrderId,
	}
	tx := s.repoSQL.GetDB().Begin()
	invoice, err = s.repoSQL.InvoiceRepo().Save(tx, invoice)
	if err != nil {
		tx.Rollback()
		return
	}

	order := &invoice.Order
	order.Status = entity.StatusInProgress
	order, err = s.repoSQL.OrderRepo().Save(tx, order)
	if err != nil {
		tx.Rollback()
		return
	}

	result = &model.InvoiceResponse{
		ID:             invoice.ID,
		PaymentMethod:  invoice.PaymentMethod,
		PaymentStatus:  invoice.PaymentStatus,
		PaymentDueDate: invoice.PaymentDueDate,
		Order: model.OrderResponse{
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
			DeletedAt: order.UpdatedAt,
		},
		CreatedAt: invoice.CreatedAt,
		UpdatedAt: invoice.UpdatedAt,
		DeletedAt: invoice.DeletedAt,
	}

	tx.Commit()

	return
}

func (s *invoiceServiceImpl) UpdateInvoice(ctx context.Context, invoiceModel model.InvoiceCreateOrUpdateModel) (*entity.Invoice, error) {
	var invoice entity.Invoice
	filter := model.InvoiceFilter{ID: &invoiceModel.ID}
	invoice, err := s.repoSQL.InvoiceRepo().Find(ctx, &filter)
	if err != nil {
		return &entity.Invoice{}, err
	}

	invoice.PaymentMethod = invoiceModel.PaymentMethod
	invoice.PaymentStatus = invoiceModel.PaymentStatus

	return s.repoSQL.InvoiceRepo().Save(s.repoSQL.GetDB(), &invoice)
}

func (s *invoiceServiceImpl) DeleteInvoice(ctx context.Context, id string) (err error) {
	filter := model.InvoiceFilter{ID: &id}
	invoice, err := s.repoSQL.InvoiceRepo().Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	invoice.DeletedAt = &deleted_at
	_, err = s.repoSQL.InvoiceRepo().Save(s.repoSQL.GetDB(), &invoice)
	return
}
