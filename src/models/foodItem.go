package models

import (
	"gorm.io/gorm"
)

type FoodItem struct {
	gorm.Model
	Name  string `gorm:"unique;error:name must be unique"`
	Price int64
}
