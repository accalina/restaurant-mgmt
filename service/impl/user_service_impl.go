package impl

import (
	"context"
	"errors"
	"time"

	"github.com/accalina/restaurant-mgmt/common"
	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/exception"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/accalina/restaurant-mgmt/service"
	"github.com/google/uuid"
)

type userServiceImpl struct {
	repository.UserRepository
}

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

func (userService *userServiceImpl) FindAll(ctx context.Context) (response []model.UserModel) {
	users := userService.UserRepository.FindAll(ctx)
	for _, user := range users {

		createdAt := time.Time{}
		createdAt = *user.CreatedAt
		response = append(response, model.UserModel{
			Id:        user.Id,
			Username:  user.Username,
			CreatedAt: createdAt,
		})
	}
	if len(users) == 0 {
		return []model.UserModel{}
	}
	return response
}

func (userService *userServiceImpl) FindById(ctx context.Context, id string) (model.UserModel, error) {
	user, err := userService.UserRepository.FindById(ctx, id)
	if err != nil {
		return model.UserModel{}, errors.New("user not found")
	}
	createdAt := time.Time{}
	updatedAt := time.Time{}
	createdAt = *user.CreatedAt
	updatedAt = *user.UpdatedAt
	return model.UserModel{
		Id:        user.Id,
		Username:  user.Username,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (userService *userServiceImpl) Register(ctx context.Context, userModel model.UserCreateModel) model.UserCreateModel {
	id := uuid.New()
	password, err := common.HashPassword(userModel.Password)
	currentTime := time.Now()
	exception.PanicLogging(err)

	user := entity.User{
		Id:        id,
		Username:  userModel.Username,
		Password:  password,
		Role:      entity.Role.User,
		CreatedAt: &currentTime,
	}
	userService.UserRepository.Register(ctx, user)
	return userModel
}

func (userService *userServiceImpl) Login(ctx context.Context, username string, password string) (model.UserModel, error) {
	currentTime := time.Now()
	user, err := userService.UserRepository.Login(ctx, username, password)
	if err != nil {
		return model.UserModel{}, errors.New("user not found")
	}

	// Update lastLogin
	userEntity := entity.User{LastLogin: &currentTime}
	errUpdateLastlogin := userService.UserRepository.Update(ctx, userEntity, user.Id.String())
	if errUpdateLastlogin != nil {
		return model.UserModel{}, err
	}
	createdAt := time.Time{}
	updatedAt := time.Time{}
	createdAt = *user.CreatedAt
	updatedAt = *user.UpdatedAt

	return model.UserModel{
		Id:        user.Id,
		Username:  user.Username,
		LastLogin: currentTime,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (userService *userServiceImpl) Update(ctx context.Context, userModel model.UserUpdateModel, id string) model.UserUpdateModel {
	currentTime := time.Now()
	password, err := common.HashPassword(userModel.Password)
	exception.PanicLogging(err)
	user := entity.User{
		Password:  password,
		UpdatedAt: &currentTime,
	}
	userService.UserRepository.Update(ctx, user, id)
	return userModel
}

func (userService *userServiceImpl) Delete(ctx context.Context, id string) bool {
	currentTime := time.Now()
	user := entity.User{
		DeletedAt: &currentTime,
	}
	err := userService.UserRepository.Update(ctx, user, id)
	return err != nil
}

func (userService *userServiceImpl) Promote(ctx context.Context, id string) bool {
	currentTime := time.Now()
	user := entity.User{
		Role:      entity.Role.Admin,
		UpdatedAt: &currentTime,
	}
	err := userService.UserRepository.Update(ctx, user, id)
	return err != nil
}

func (userService *userServiceImpl) Demote(ctx context.Context, id string) bool {
	currentTime := time.Now()
	user := entity.User{
		Role:      entity.Role.User,
		UpdatedAt: &currentTime,
	}
	err := userService.UserRepository.Update(ctx, user, id)
	return err != nil
}
