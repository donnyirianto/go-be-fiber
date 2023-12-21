package service

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/model"
)

type TransactionService interface {
	Create(ctx context.Context, model model.TransactionCreateUpdateModel) model.TransactionCreateUpdateModel
	Delete(ctx context.Context, id string)
	FindById(ctx context.Context, id string) model.TransactionModel
	FindAll(ctx context.Context) []model.TransactionModel
}
