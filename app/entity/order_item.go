package entity

import (
	"time"
)

type OrderItem struct {
	ID        string `gorm:"primaryKey;column:id;type:varchar(36)"`
	Qty       int    `gorm:"column:qty"`
	FoodId    string `gorm:"column:food_id;type:varchar(36)"`
	Food      Food
	OrderId   string `gorm:"column:order_id;type:varchar(36)"`
	Order     Order
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:null"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null"`
}
