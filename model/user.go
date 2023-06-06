package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserModel struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreateModel struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdateModel struct {
	Password string `json:"password" validate:"required,min=6"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}

func (f *UserCreateModel) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}

func (f *UserUpdateModel) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}
