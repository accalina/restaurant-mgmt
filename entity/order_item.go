package entity

import (
	"time"
)

type OrderItem struct {
	ID        string  `gorm:"primaryKey;column:id;type:varchar(36)" json:"id"`
	Qty       int     `gorm:"column:qty" json:"qty"`
	FoodId    *string `gorm:"column:food_id;type:varchar(36)" json:"foodId"`
	Food      *Food
	OrderId   *string `gorm:"column:order_id;type:varchar(36)" json:"orderId"`
	Order     *Order
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:null" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null" json:"deletedAt"`
}
