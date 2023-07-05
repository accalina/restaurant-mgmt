package impl

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/app/service"
	"github.com/accalina/restaurant-mgmt/pkg/common"
	"github.com/accalina/restaurant-mgmt/pkg/configuration"
	"github.com/accalina/restaurant-mgmt/pkg/exception"
	"github.com/accalina/restaurant-mgmt/pkg/shared/repository"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type userServiceImpl struct {
	repoSQL repository.RepoSQL
	Redis   *redis.Client
}

func NewUserServiceImpl() service.UserService {
	return &userServiceImpl{repoSQL: repository.GetSharedRepoSQL(), Redis: configuration.GetRedisCache()}
}

func (u *userServiceImpl) FindAll(ctx context.Context) (response []model.UserModel) {
	users := u.repoSQL.UserRepo().FindAll(ctx)
	for _, user := range users {

		createdAt := time.Time{}
		createdAt = *user.CreatedAt
		response = append(response, model.UserModel{
			Id:        user.ID,
			Username:  user.Username,
			CreatedAt: createdAt,
		})
	}
	if len(users) == 0 {
		return []model.UserModel{}
	}
	return response
}

func (u *userServiceImpl) FindById(ctx context.Context, id string) (model.UserModel, error) {
	user, err := u.repoSQL.UserRepo().FindById(ctx, id)
	if err != nil {
		return model.UserModel{}, errors.New("user not found")
	}
	createdAt := time.Time{}
	updatedAt := time.Time{}
	createdAt = *user.CreatedAt
	updatedAt = *user.UpdatedAt
	return model.UserModel{
		Id:        user.ID,
		Username:  user.Username,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (u *userServiceImpl) Register(ctx context.Context, userModel model.UserCreateModel) model.UserCreateModel {
	id := uuid.New()
	password, err := common.HashPassword(userModel.Password)
	currentTime := time.Now()
	exception.PanicLogging(err)

	user := entity.User{
		ID:        id,
		Username:  userModel.Username,
		Password:  password,
		Role:      entity.Role.User,
		CreatedAt: &currentTime,
	}
	u.repoSQL.UserRepo().Register(ctx, user)
	return userModel
}

func (u *userServiceImpl) Login(ctx context.Context, username string, password string) (model.ResponseLogin, error) {
	currentTime := time.Now()
	user, err := u.repoSQL.UserRepo().Login(ctx, username, password)
	if err != nil {
		return model.ResponseLogin{}, errors.New("user not found")
	}

	// Update lastLogin
	userEntity := entity.User{LastLogin: &currentTime}
	errUpdateLastlogin := u.repoSQL.UserRepo().Update(ctx, userEntity, user.ID.String())
	if errUpdateLastlogin != nil {
		return model.ResponseLogin{}, err
	}
	token := common.GenerateToken(user.Username, user.Role)
	expiredTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	err = u.Redis.Set(ctx, user.Username, user.Role, time.Duration(expiredTime)*time.Minute).Err()
	exception.PanicLogging(err)

	return model.ResponseLogin{
		Token: token,
	}, nil
}

func (u userServiceImpl) Logout(ctx context.Context, username string) {
	u.Redis.Del(ctx, username)
}

func (u *userServiceImpl) Update(ctx context.Context, userModel model.UserUpdateModel, id string) model.UserUpdateModel {
	currentTime := time.Now()
	password, err := common.HashPassword(userModel.Password)
	exception.PanicLogging(err)
	user := entity.User{
		Password:  password,
		UpdatedAt: &currentTime,
	}
	u.repoSQL.UserRepo().Update(ctx, user, id)
	return userModel
}

func (u *userServiceImpl) Delete(ctx context.Context, id string) bool {
	currentTime := time.Now()
	user := entity.User{
		DeletedAt: &currentTime,
	}
	err := u.repoSQL.UserRepo().Update(ctx, user, id)
	return err != nil
}

func (u *userServiceImpl) Promote(ctx context.Context, id string) bool {
	currentTime := time.Now()
	user := entity.User{
		Role:      entity.Role.Admin,
		UpdatedAt: &currentTime,
	}
	err := u.repoSQL.UserRepo().Update(ctx, user, id)
	return err != nil
}

func (u *userServiceImpl) Demote(ctx context.Context, id string) bool {
	currentTime := time.Now()
	user := entity.User{
		Role:      entity.Role.User,
		UpdatedAt: &currentTime,
	}
	err := u.repoSQL.UserRepo().Update(ctx, user, id)
	return err != nil
}
