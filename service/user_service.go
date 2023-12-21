package service

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/entity"
	"github.com/donnyirianto/go-be-fiber/model"
)

type UserService interface {
	Authentication(ctx context.Context, model model.UserModel) entity.User
}
