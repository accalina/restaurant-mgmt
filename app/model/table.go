package model

import (
	"time"
)

type TableResponse struct {
	ID             string     `json:"id"`
	No             int        `json:"no"`
	NumberOfGuests int        `json:"numberOfGuests"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
	DeletedAt      *time.Time `json:"deletedAt"`
}

type TableFilter struct {
	Filter
	ID *string `json:"id"`
}

func NewTableFilter() *TableFilter {
	return &TableFilter{
		Filter: *DefaultFilter(),
		ID:     new(string),
	}
}

type TableCreateOrUpdateModel struct {
	ID             string `json:"id" validate:"max=36"`
	No             int    `json:"no" validate:"required,min=1"`
	NumberOfGuests int    `json:"numberOfGuests" validate:"required,min=1"`
}
