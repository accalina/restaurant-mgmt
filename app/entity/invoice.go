package entity

import (
	"time"
)

type Invoice struct {
	ID             string     `gorm:"primaryKey;column:id;type:varchar(36)"`
	PaymentMethod  string     `gorm:"column:payment_method;type:varchar(50)"`
	PaymentStatus  string     `gorm:"column:payment_status;type:varchar(30)"`
	PaymentDueDate time.Time  `gorm:"column:payment_due_date"`
	OrderId        string     `gorm:"column:order_id;type:varchar(36)"`
	Order          Order      
	CreatedAt      time.Time  `gorm:"column:created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;default:null"`
	DeletedAt      *time.Time `gorm:"column:deleted_at;default:null"`
}
