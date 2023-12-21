package impl

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/common"
	"github.com/donnyirianto/go-be-fiber/configuration"
	"github.com/donnyirianto/go-be-fiber/entity"
	"github.com/donnyirianto/go-be-fiber/exception"
	"github.com/donnyirianto/go-be-fiber/model"
	"github.com/donnyirianto/go-be-fiber/repository"
	"github.com/donnyirianto/go-be-fiber/service"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func NewProductServiceImpl(productRepository *repository.ProductRepository, cache *redis.Client) service.ProductService {
	return &productServiceImpl{ProductRepository: *productRepository, Cache: cache}
}

type productServiceImpl struct {
	repository.ProductRepository
	Cache *redis.Client
}

func (service *productServiceImpl) Create(ctx context.Context, productModel model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel {
	common.Validate(productModel)
	product := entity.Product{
		Name:     productModel.Name,
		Price:    productModel.Price,
		Quantity: productModel.Quantity,
	}
	service.ProductRepository.Insert(ctx, product)
	return productModel
}

func (service *productServiceImpl) Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id string) model.ProductCreateOrUpdateModel {
	common.Validate(productModel)
	product := entity.Product{
		Id:       uuid.MustParse(id),
		Name:     productModel.Name,
		Price:    productModel.Price,
		Quantity: productModel.Quantity,
	}
	service.ProductRepository.Update(ctx, product)
	return productModel
}

func (service *productServiceImpl) Delete(ctx context.Context, id string) {
	product, err := service.ProductRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.ProductRepository.Delete(ctx, product)
}

func (service *productServiceImpl) FindById(ctx context.Context, id string) model.ProductModel {
	productCache := configuration.SetCache[entity.Product](service.Cache, ctx, "product", id, service.ProductRepository.FindById)
	return model.ProductModel{
		Id:       productCache.Id.String(),
		Name:     productCache.Name,
		Price:    productCache.Price,
		Quantity: productCache.Quantity,
	}
}

func (service *productServiceImpl) FindAll(ctx context.Context) (responses []model.ProductModel) {
	products := service.ProductRepository.FindAl(ctx)
	for _, product := range products {
		responses = append(responses, model.ProductModel{
			Id:       product.Id.String(),
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	if len(products) == 0 {
		return []model.ProductModel{}
	}
	return responses
}
