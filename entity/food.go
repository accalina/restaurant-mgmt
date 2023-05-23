package entity

import (
	"time"

	"github.com/google/uuid"
)

type Food struct {
	Id        uuid.UUID  `gorm:"primaryKey;column:id;type:varchar(36)" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(255)" json:"name"`
	Price     int64      `gorm:"column:price" json:"price"`
	Qty       int32      `gorm:"column:qty" json:"qty"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null" json:"deleted_at"`
}
