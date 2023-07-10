package entity

import (
	"time"
)

type Table struct {
	ID             string     `gorm:"primaryKey;column:id;type:varchar(36)"`
	No             int        `gorm:"column:no;unique"`
	NumberOfGuests int        `gorm:"column:number_of_guests"`
	IsAvailable    bool       `gorm:"column:is_available;default:true"`
	CreatedAt      time.Time  `gorm:"column:created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;default:null"`
	DeletedAt      *time.Time `gorm:"column:deleted_at;default:null"`
}
