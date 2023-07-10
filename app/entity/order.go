package entity

import (
	"time"
)

type OrderStatus string

const (
	StatusPending     OrderStatus = "Pending"
	StatusInProgress  OrderStatus = "In Progress"
	StatusInCompleted OrderStatus = "Completed"
	StatusInCancelled OrderStatus = "Cancelled"
)

type Order struct {
	ID         string    `gorm:"primaryKey;column:id;type:varchar(36)"`
	Name       string    `gorm:"column:name;type:varchar(50)"`
	OrderDate  time.Time `gorm:"column:order_date"`
	TableID    string    `gorm:"column:table_id;type:varchar(36)"`
	Table      Table
	OrderItems []OrderItem
	Status     OrderStatus `gorm:"column:status;default:Pending"`
	CreatedAt  time.Time   `gorm:"column:created_at"`
	UpdatedAt  *time.Time  `gorm:"column:updated_at;default:null"`
	DeletedAt  *time.Time  `gorm:"column:deleted_at;default:null"`
}
