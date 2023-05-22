package entity

import (
	"time"

	"github.com/google/uuid"
)

type Food struct {
	Id        uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36)" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Price     float64   `gorm:"column:price;type:float64" json:"price"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
