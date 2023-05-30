package entity

import (
	"time"
)

type Menu struct {
	ID        string     `gorm:"primaryKey;column:id;type:varchar(36)" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(255)" json:"name"`
	Category  string     `gorm:"column:category;type:varchar(255)" json:"category"`
	StartDate *time.Time `gorm:"column:start_date;default:null" json:"startDate"`
	EndDate   *time.Time `gorm:"column:end_date;default:null" json:"endDate"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:null" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null" json:"deletedAt"`
}
