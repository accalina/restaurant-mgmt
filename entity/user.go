package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36)" json:"id"`
	Username  string    `gorm:"unique;error:username must be unique;column:username;type:varchar(255)" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
