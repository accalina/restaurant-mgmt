package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
)

type InvoiceRepository interface {
	FetchAll(ctx context.Context, filter *model.InvoiceFilter) ([]entity.Invoice, error)
	Count(ctx context.Context, filter *model.InvoiceFilter) int
	Find(ctx context.Context, filter *model.InvoiceFilter) (entity.Invoice, error)
	Save(ctx context.Context, data *entity.Invoice) (*entity.Invoice, error)
}