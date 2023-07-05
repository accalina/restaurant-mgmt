package entity

import (
	"time"

	"github.com/google/uuid"
)

type RoleChoice struct {
	User  string
	Admin string
}

var Role = RoleChoice{
	User:  "user",
	Admin: "admin",
}

type User struct {
	ID        uuid.UUID  `gorm:"primaryKey;column:id;type:varchar(36)" json:"id"`
	Username  string     `gorm:"unique;error:username must be unique;column:username;type:varchar(255)" json:"username"`
	Password  string     `gorm:"column:password" json:"password"`
	Role      string     `gorm:"column:role" json:"role"`
	LastLogin *time.Time `gorm:"column:last_login;default:null" json:"lastLogin"`
	CreatedAt *time.Time `gorm:"column:created_at;default:null" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:null" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null" json:"deletedAt"`
}
