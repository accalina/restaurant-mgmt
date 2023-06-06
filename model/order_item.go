package model

import (
	"time"

	"github.com/go-playground/validator/v10"
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
}

func NewOrderItemFilter() *OrderItemFilter {
	return &OrderItemFilter{
		Filter: *DefaultFilter(),
		ID:     new(string),
	}
}

type OrderItemCreateOrUpdateModel struct {
	ID      string `json:"id" validate:"max=36"`
	Qty     int    `json:"qty" validate:"required,min=1"`
	FoodId  string `json:"foodId" validate:"required,len=36"`
	OrderId string `json:"orderId" validate:"required,len=36"`
}

func (f *OrderItemCreateOrUpdateModel) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}
