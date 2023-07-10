package model

import (
	"time"
)

type OrderItemResponse struct {
	ID        string     `json:"id"`
	Qty       int        `json:"qty"`
	FoodId    string     `json:"foodId"`
	OrderId   string     `json:"orderId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type OrderItemFilter struct {
	Filter
	ID *string `json:"id"`
	FoodID *string `json:"foodId"`
	OrderID *string `json:"orderId"`
}

func NewOrderItemFilter(preloads ...string) *OrderItemFilter {
	return &OrderItemFilter{
		Filter: *DefaultFilter(preloads...),
		ID:     new(string),
		FoodID:     new(string),
		OrderID:     new(string),
	}
}

type OrderItemCreateOrUpdateModel struct {
	ID      string `json:"id" validate:"max=36"`
	Qty     int    `json:"qty" validate:"required,min=1"`
	FoodId  string `json:"foodId" validate:"required,len=36"`
	OrderId string `json:"orderId" validate:"required,len=36"`
}

type OrderItemChangeQtyModel struct {
	ID      string `json:"id" validate:"max=36"`
	Qty     int    `json:"qty" validate:"min=0"`
}
