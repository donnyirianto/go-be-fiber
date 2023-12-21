package impl

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/exception"
	"github.com/donnyirianto/go-be-fiber/model"
	"github.com/donnyirianto/go-be-fiber/repository"
	"github.com/donnyirianto/go-be-fiber/service"
)

func NewTransactionDetailServiceImpl(transactionDetailRepository *repository.TransactionDetailRepository) service.TransactionDetailService {
	return &transactionDetailServiceImpl{TransactionDetailRepository: *transactionDetailRepository}
}

type transactionDetailServiceImpl struct {
	repository.TransactionDetailRepository
}

func (transactionDetailService *transactionDetailServiceImpl) FindById(ctx context.Context, id string) model.TransactionDetailModel {
	transactionDetail, err := transactionDetailService.TransactionDetailRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return model.TransactionDetailModel{
		Id:            transactionDetail.Id.String(),
		SubTotalPrice: transactionDetail.SubTotalPrice,
		Price:         transactionDetail.Price,
		Quantity:      transactionDetail.Quantity,
		Product: model.ProductModel{
			Id:       transactionDetail.Product.Id.String(),
			Name:     transactionDetail.Product.Name,
			Price:    transactionDetail.Product.Price,
			Quantity: transactionDetail.Product.Quantity,
		},
	}
}
