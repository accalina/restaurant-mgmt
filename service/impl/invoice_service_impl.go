package impl

import (
	"context"
	"time"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/accalina/restaurant-mgmt/service"
)

type invoiceServiceImpl struct {
	repository.InvoiceRepository
}

func NewInvoiceServiceImpl(r *repository.InvoiceRepository) service.InvoiceService {
	return &invoiceServiceImpl{InvoiceRepository: *r}
}

func (s *invoiceServiceImpl) GetAllInvoice(ctx context.Context, filter *model.InvoiceFilter) (result []model.InvoiceResponse, meta model.Meta, err error) {
	invoices, err := s.InvoiceRepository.FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.InvoiceRepository.Count(ctx, filter)
	meta = model.NewMeta(filter.Page, filter.Limit, count)
	for _, invoice := range invoices {
		result = append(result, model.InvoiceResponse{
			ID:             invoice.ID,
			PaymentMethod:  invoice.PaymentMethod,
			PaymentStatus:  invoice.PaymentStatus,
			PaymentDueDate: invoice.PaymentDueDate,
			OrderId:        invoice.OrderId,
			CreatedAt:      invoice.CreatedAt,
			UpdatedAt:      invoice.UpdatedAt,
			DeletedAt:      invoice.DeletedAt,
		})
	}
	return
}

func (s *invoiceServiceImpl) GetDetailInvoice(ctx context.Context, id string) (result model.InvoiceResponse, err error) {
	var data entity.Invoice
	filter := model.InvoiceFilter{ID: &id}
	data, err = s.InvoiceRepository.Find(ctx, &filter)
	if err != nil {
		return
	}

	result.ID = data.ID
	result.PaymentMethod = data.PaymentMethod
	result.PaymentStatus = data.PaymentStatus
	result.PaymentDueDate = data.PaymentDueDate
	result.OrderId = data.OrderId
	result.CreatedAt = data.CreatedAt
	result.UpdatedAt = data.UpdatedAt
	result.DeletedAt = data.DeletedAt

	return
}

func (s *invoiceServiceImpl) CreateInvoice(ctx context.Context, invoiceModel model.InvoiceCreateOrUpdateModel) (*entity.Invoice, error) {
	invoice := entity.Invoice{
		PaymentMethod: invoiceModel.PaymentMethod,
		PaymentStatus: invoiceModel.PaymentStatus,
		OrderId:       invoiceModel.OrderId,
	}
	return s.InvoiceRepository.Save(ctx, &invoice)
}

func (s *invoiceServiceImpl) UpdateInvoice(ctx context.Context, invoiceModel model.InvoiceCreateOrUpdateModel) (*entity.Invoice, error) {
	var invoice entity.Invoice
	filter := model.InvoiceFilter{ID: &invoiceModel.ID}
	invoice, err := s.InvoiceRepository.Find(ctx, &filter)
	if err != nil {
		return &entity.Invoice{}, err
	}

	invoice.PaymentMethod = invoiceModel.PaymentMethod
	invoice.PaymentStatus = invoiceModel.PaymentStatus

	return s.InvoiceRepository.Save(ctx, &invoice)
}

func (s *invoiceServiceImpl) DeleteInvoice(ctx context.Context, id string) (err error) {
	filter := model.InvoiceFilter{ID: &id}
	invoice, err := s.InvoiceRepository.Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	invoice.DeletedAt = &deleted_at
	_, err = s.InvoiceRepository.Save(ctx, &invoice)
	return
}
