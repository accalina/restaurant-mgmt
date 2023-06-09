package model

import (
	"time"
)

type OrderResponse struct {
	ID        string     `json:"id"`
	OrderDate time.Time  `json:"orderDate"`
	TableID   string     `json:"tableId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type OrderFilter struct {
	Filter
	ID      *string `json:"id"`
	TableID *string `json:"tableId"`
}

func NewOrderFilter() *OrderFilter {
	return &OrderFilter{
		Filter:  *DefaultFilter(),
		ID:      new(string),
		TableID: new(string),
	}
}

type OrderCreateOrUpdateModel struct {
	ID      string `json:"id" validate:"max=36"`
	TableID string `json:"tableId" validate:"required,len=36"`
}
