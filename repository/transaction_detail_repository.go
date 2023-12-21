package repository

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/entity"
)

type TransactionDetailRepository interface {
	FindById(ctx context.Context, id string) (entity.TransactionDetail, error)
}
