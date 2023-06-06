package impl

import (
	"context"
	"strings"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type invoiceRepositoryImpl struct {
	DB *gorm.DB
}

func NewInvoiceRepositoryImpl(DB *gorm.DB) repository.InvoiceRepository {
	return &invoiceRepositoryImpl{DB}
}

func (r *invoiceRepositoryImpl) FetchAll(ctx context.Context, filter *model.InvoiceFilter) (invoices []entity.Invoice, err error) {
	err = r.setFilter(r.DB, filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&invoices).Error
	return
}

func (r *invoiceRepositoryImpl) Count(ctx context.Context, filter *model.InvoiceFilter) (count int) {
	var total int64
	r.setFilter(r.DB, filter).Model(&entity.Invoice{}).Count(&total)
	count = int(total)
	return
}

func (r *invoiceRepositoryImpl) Find(ctx context.Context, filter *model.InvoiceFilter) (result entity.Invoice, err error) {
	err = r.setFilter(r.DB, filter).First(&result).Error
	return
}

func (r *invoiceRepositoryImpl) Save(ctx context.Context, invoice *entity.Invoice) (*entity.Invoice, error) {
	if invoice.ID == "" {
		invoice.ID = uuid.NewString()
		err := r.DB.Create(invoice).Error
		return invoice, err
	}
	err := r.DB.Save(invoice).Error
	return invoice, err
}

func (r *invoiceRepositoryImpl) setFilter(db *gorm.DB, filter *model.InvoiceFilter) *gorm.DB {
	if *filter.ID != "" {
		db = db.Where("id = ?", *filter.ID)
	}

	if filter.Search != "" {
		db = db.Where("id ILIKE '%%' || ? || '%%' OR name ILIKE '%%' || ? || '%%' OR category ILIKE '%%' || ? || '%%'", filter.Search, filter.Search, filter.Search)
	}

	return db
}
