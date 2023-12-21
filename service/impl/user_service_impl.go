package impl

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/entity"
	"github.com/donnyirianto/go-be-fiber/exception"
	"github.com/donnyirianto/go-be-fiber/model"
	"github.com/donnyirianto/go-be-fiber/repository"
	"github.com/donnyirianto/go-be-fiber/service"
	"golang.org/x/crypto/bcrypt"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	repository.UserRepository
}

func (userService *userServiceImpl) Authentication(ctx context.Context, model model.UserModel) entity.User {
	userResult, err := userService.UserRepository.Authentication(ctx, model.Username)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(model.Password))
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "incorrect username and password",
		})
	}
	return userResult
}
