package model

// type Food struct {
// 	Id        uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36) json:"id"`
// 	Field     string    `gorm:"column:field;type:varchar(255)" json:"field"`
// 	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
// 	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
// }

type FoodCreteOrUpdateModel struct {
	Field string `json:"field" validate:"required,min=1`
}
