package service

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/model"
)

type TransactionDetailService interface {
	FindById(ctx context.Context, id string) model.TransactionDetailModel
}
