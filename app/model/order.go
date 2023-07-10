package model

import (
	"time"
)

type OrderResponse struct {
	ID         string              `json:"id"`
	Name       string              `json:"name"`
	OrderDate  time.Time           `json:"orderDate"`
	Status     string              `json:"status"`
	Table      TableResponse       `json:"table"`
	OrderItems []OrderItemResponse `json:"orderItems"`
	CreatedAt  time.Time           `json:"createdAt"`
	UpdatedAt  *time.Time          `json:"updatedAt"`
	DeletedAt  *time.Time          `json:"deletedAt"`
}

type OrderFilter struct {
	Filter
	ID      *string `json:"id"`
	TableID *string `json:"tableId"`
}

func NewOrderFilter(preloads ...string) *OrderFilter {
	return &OrderFilter{
		Filter:  *DefaultFilter(preloads...),
		ID:      new(string),
		TableID: new(string),
	}
}

type OrderCreateOrUpdateModel struct {
	ID      string `json:"id" validate:"max=36"`
	Name    string `json:"name" validate:"max=50"`
	TableID string `json:"tableId" validate:"required,len=36"`
}

type OrderDoneModel struct {
	ID      string `json:"id" validate:"max=36"`
	Status string `json:"status" validate:"oneof=Completed Cancelled"`
}
