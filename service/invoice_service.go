package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
)

type InvoiceService interface {
	GetAllInvoice(ctx context.Context, filter *model.InvoiceFilter) (result []model.InvoiceResponse, meta model.Meta, err error)
	GetDetailInvoice(ctx context.Context, id string) (result model.InvoiceResponse, err error)
	CreateInvoice(ctx context.Context, data model.InvoiceCreateOrUpdateModel) (*entity.Invoice, error)
	UpdateInvoice(ctx context.Context, data model.InvoiceCreateOrUpdateModel) (*entity.Invoice, error)
	DeleteInvoice(ctx context.Context, id string) error
}
