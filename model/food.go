package model

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

// type Food struct {
// 	Id        uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36) json:"id"`
// 	Field     string    `gorm:"column:field;type:varchar(255)" json:"field"`
// 	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
// 	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
// }

type FoodCreteOrUpdateModel struct {
	Name  string  `json:"name" validate:"required,min=2"`
	Price float64 `json:"price" validate:"required,min=0"`
}

func (f *FoodCreteOrUpdateModel) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}

func (f *FoodCreteOrUpdateModel) Validate() error {
	validate := validator.New()
	return validate.Struct(f)

}
