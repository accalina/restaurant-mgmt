package entity

import (
	"time"
)

type Table struct {
	ID             string     `gorm:"primaryKey;column:id;type:varchar(36)" json:"id"`
	No             int        `gorm:"column:no;unique" json:"no"`
	NumberOfGuests int        `gorm:"column:number_of_guests" json:"numberOfGuests"`
	CreatedAt      time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;default:null" json:"updatedAt"`
	DeletedAt      *time.Time `gorm:"column:deleted_at;default:null" json:"deletedAt"`
}
