package entity

import (
	"time"
)

type Invoice struct {
	ID             string     `gorm:"primaryKey;column:id;type:varchar(36)" json:"id"`
	PaymentMethod  string     `gorm:"column:payment_method;type:varchar(50)" json:"paymentMethod"`
	PaymentStatus  string     `gorm:"column:payment_status;type:varchar(30)" json:"paymentStatus"`
	PaymentDueDate time.Time  `gorm:"column:payment_due_date" json:"paymentDueDate"`
	OrderId        string     `gorm:"column:order_id;type:varchar(36)" json:"orderId"`
	Order          Order      `json:"-"`
	CreatedAt      time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;default:null" json:"updatedAt"`
	DeletedAt      *time.Time `gorm:"column:deleted_at;default:null" json:"deletedAt"`
}
