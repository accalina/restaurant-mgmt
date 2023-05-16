package entity

import (
	"time"

	"github.com/google/uuid"
)

type Food struct {
	Id        uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36) json:"id"`
	Field     string    `gorm:"column:field;type:varchar(255)" json:"field"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
