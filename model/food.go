package model

import (
	"encoding/json"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type FoodFilter struct {
	Filter
	ID   *string `json:"id"`
	Name string  `json:"name"`
}

func NewFoodFilter() *FoodFilter {
	return &FoodFilter{
		Filter: *DefaultFilter(),
		ID:     new(string),
		Name:   "",
	}
}

type FoodModel struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Price     int64      `json:"price"`
	Qty       int32      `json:"qty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type FoodCreteOrUpdateModel struct {
	Name  string `json:"name" validate:"required,min=1"`
	Price int64  `json:"price" validate:"required,min=1"`
	Qty   int32  `json:"qty" validate:"required"`
}

func (f *FoodCreteOrUpdateModel) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}

func (f *FoodCreteOrUpdateModel) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}
