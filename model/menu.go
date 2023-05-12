package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name     string `gorm:"unique;error:name must be unique"`
	Category string
}
