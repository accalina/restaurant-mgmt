package entity

import (
	"time"
)

type Menu struct {
	ID        string     `gorm:"primaryKey;column:id;type:varchar(36)"`
	Name      string     `gorm:"column:name;type:varchar(255)"`
	Category  string     `gorm:"column:category;type:varchar(255)"`
	StartDate *time.Time `gorm:"column:start_date;default:null"`
	EndDate   *time.Time `gorm:"column:end_date;default:null"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:null"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null"`
	Foods     []Food
}
