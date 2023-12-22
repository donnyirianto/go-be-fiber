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
	"github.com/redis/go-redis/v9"
)

func NewCoverageServiceImpl(coverageRepository *repository.CoverageRepository, cache *redis.Client) service.CoverageService {
	return &coverageServiceImpl{CoverageRepository: *coverageRepository, Cache: cache}
}

type coverageServiceImpl struct {
	repository.CoverageRepository
	Cache *redis.Client
}

func (service *coverageServiceImpl) FindAll(ctx context.Context) (responses []model.CoverageModel) {
	coverage = service.CoverageRepository.FindAll(ctx)

	for _, coverage := range coverage {
		responses = append(responses, model.CoverageModel{
			Id:    coverage.Id,
			Kdcab: coverage.Kdcab,
			Nik:   coverage.Nik,
		})
	}
	if len(coverage) == 0 {
		return []model.CoverageModel{}
	}
	return responses
}

func (service *coverageServiceImpl) Create(ctx context.Context, coverageModel model.CoverageCreatUpdateModel) model.CoverageCreatUpdateModel {
	common.Validate(coverageModel)
	coverage := entity.Coverage{
		Id:    int64(coverageModel.Id),
		Kdcab: coverageModel.Kdcab,
		Nik:   coverageModel.Nik,
	}
	service.CoverageRepository.Insert(ctx, coverage)
	return coverageModel
}

func (service *coverageServiceImpl) Update(ctx context.Context, coverageModel model.CoverageCreatUpdateModel, id string) model.CoverageCreatUpdateModel {
	common.Validate(coverageModel)
	coverage := entity.Coverage{
		Id:    int64(coverageModel.Id),
		Kdcab: coverageModel.Kdcab,
		Nik:   coverageModel.Nik,
	}
	service.CoverageRepository.Update(ctx, coverage)
	return coverageModel
}

func (service *coverageServiceImpl) Delete(ctx context.Context, id int64) {
	product, err := service.CoverageRepository.FindByID(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.CoverageRepository.Delete(ctx, product)
}

func (service *coverageServiceImpl) FindByID(ctx context.Context, id int64) model.CoverageModel {
	coverageCache := configuration.SetCache[entity.Coverage](service.Cache, ctx, "coverage", id, service.CoverageRepository.FindByID)
	return model.CoverageModel{
		Id:    coverageCache.Id,
		Kdcab: coverageCache.Kdcab,
		Nik:   coverageCache.Nik,
	}
}
