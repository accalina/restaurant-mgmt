package entity

import (
	"time"
)

type Food struct {
	ID        string `gorm:"primaryKey;column:id;type:varchar(36)"`
	Name      string `gorm:"column:name;type:varchar(255)"`
	Price     int64  `gorm:"column:price"`
	Qty       int32  `gorm:"column:qty"`
	MenuID    string `gorm:"column:menu_id;type:varchar(36)"`
	Menu      Menu
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:null"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null"`
}
