package model

import (
	"time"
)

type InvoiceResponse struct {
	ID             string        `json:"id"`
	PaymentMethod  string        `json:"paymentMethod"`
	PaymentStatus  string        `json:"paymentStatus"`
	PaymentDueDate time.Time     `json:"paymentDueDate"`
	Order          OrderResponse `json:"order"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      *time.Time    `json:"updatedAt"`
	DeletedAt      *time.Time    `json:"deletedAt"`
}

type InvoiceFilter struct {
	Filter
	ID *string `json:"id"`
}

func NewInvoiceFilter(preloads ...string) *InvoiceFilter {
	return &InvoiceFilter{
		Filter: *DefaultFilter(preloads...),
		ID:     new(string),
	}
}

type InvoiceCreateOrUpdateModel struct {
	ID            string `json:"id" validate:"max=36"`
	PaymentMethod string `json:"paymentMethod" validate:"required,min=1"`
	PaymentStatus string `json:"paymentStatus" validate:"required,min=1"`
	OrderId       string `json:"orderId" validate:"required,len=36"`
}
