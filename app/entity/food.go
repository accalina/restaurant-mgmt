package entity

import (
	"time"
)

type Food struct {
	ID        string `gorm:"primaryKey;column:id;type:varchar(36)" json:"id"`
	Name      string `gorm:"column:name;type:varchar(255)" json:"name"`
	Price     int64  `gorm:"column:price" json:"price"`
	Qty       int32  `gorm:"column:qty" json:"qty"`
	MenuID    string `gorm:"column:menu_id;type:varchar(36)" json:"menuId"`
	Menu      Menu
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:null" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null" json:"deletedAt"`
}
