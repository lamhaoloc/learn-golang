package userService

import (
	"awesomeProject/common"
	"context"

	userModel "awesomeProject/module/user/model"
)

type RegisterStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*userModel.User, error)
	CreateUser(ctx context.Context, data *userModel.CreateUserDTO) error
}

type Hasher interface {
	Hash(data string) string
}

type registerService struct {
	registerStore RegisterStore
	hasher        Hasher
}

func NewRegisterService(registerStore RegisterStore, hasher Hasher) *registerService {
	return &registerService{
		registerStore: registerStore,
		hasher:        hasher,
	}
}

func (service *registerService) Register(ctx context.Context, data *userModel.CreateUserDTO) error {
	user, _ := service.registerStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return userModel.ErrEmailExisted
	}

	salt := common.GenSalt(50)
	password := service.hasher.Hash(data.Password + salt)

	data.Salt = salt
	data.Role = common.User
	data.Status = 1
	data.Password = password

	if err := service.registerStore.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(userModel.EntityName, err)
	}

	return nil
}
