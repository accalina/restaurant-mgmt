package entity

import (
	"time"
)

type Order struct {
	ID        string     `gorm:"primaryKey;column:id;type:varchar(36)" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(50)" json:"name"`
	OrderDate time.Time  `gorm:"column:order_date" json:"orderDate"`
	TableID   string     `gorm:"column:table_id;type:varchar(36)" json:"tableId"`
	Table     Table      `json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:null" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null" json:"deletedAt"`
}
