package repository

import (
	"context"
	"errors"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/pkg/common"
	"github.com/accalina/restaurant-mgmt/pkg/exception"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(ctx context.Context, user entity.User) entity.User
	Login(ctx context.Context, username string, password string) (entity.User, error)
	FindAll(ctx context.Context) []entity.User
	FindById(ctx context.Context, id string) (entity.User, error)
	Update(ctx context.Context, user entity.User, id string) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (userRepository UserRepositoryImpl) FindAll(ctx context.Context) []entity.User {
	var users []entity.User
	userRepository.DB.WithContext(ctx).Unscoped().Where("deleted_at is null").Find(&users)
	return users
}

func (userRepository UserRepositoryImpl) FindById(ctx context.Context, id string) (entity.User, error) {
	var user entity.User
	result := userRepository.DB.WithContext(ctx).Unscoped().Where("deleted_at is null").Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	return user, nil
}

func (userRepository UserRepositoryImpl) Register(ctx context.Context, user entity.User) entity.User {
	err := userRepository.DB.Create(&user).Error
	exception.PanicLogging(err)
	return user
}

func (userRepository UserRepositoryImpl) Login(ctx context.Context, username string, password string) (entity.User, error) {
	var user entity.User
	err := userRepository.DB.WithContext(ctx).Unscoped().
		Where("deleted_at is null").
		Where("username = ?", username).
		First(&user).Error
	exception.PanicLogging(err)

	isSuccess := common.CompareHashPassword(password, user.Password)
	if !isSuccess {
		return entity.User{}, errors.New("invalid username / password")
	}
	return user, nil
}

func (userRepository UserRepositoryImpl) Update(ctx context.Context, user entity.User, id string) error {
	result := userRepository.DB.WithContext(ctx).Unscoped().Where("deleted_at is null").Where("id = ?", id).Updates(&user)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
