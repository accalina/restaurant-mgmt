package model

import (
	"time"
)

type FoodResponseBase struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Price     int64      `json:"price"`
	Qty       int32      `json:"qty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type FoodResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Price     int64      `json:"price"`
	Qty       int32      `json:"qty"`
	MenuID    string     `json:"menuId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type FoodFilter struct {
	Filter
	ID   *string `json:"id"`
	Name string  `json:"name"`
}

func NewFoodFilter(preloads ...string) *FoodFilter {
	return &FoodFilter{
		Filter: *DefaultFilter(preloads...),
		ID:     new(string),
		Name:   "",
	}
}

type FoodCreateOrUpdateModel struct {
	ID     string `json:"id" validate:"max=36"`
	Name   string `json:"name" validate:"required,min=1"`
	Price  int64  `json:"price" validate:"required,min=1"`
	Qty    int32  `json:"qty" validate:"required,min=1"`
	MenuID string `json:"menuId" validate:"required,len=36"`
}

type FoodCreateOrUpdateSwaggerModel struct {
	Name   string `json:"name" validate:"required,min=1"`
	Price  int64  `json:"price" validate:"required,min=1"`
	Qty    int32  `json:"qty" validate:"required,min=1"`
	MenuID string `json:"menuId" validate:"required,len=36"`
}
