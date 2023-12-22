package service

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/model"
)

type CoverageService interface {
	Create(ctx context.Context, model model.CoverageCreatUpdateModel) model.CoverageCreatUpdateModel
	Update(ctx context.Context, CoverageModel model.CoverageCreatUpdateModel, id int64) model.CoverageCreatUpdateModel
	Delete(ctx context.Context, id int64)
	FindById(ctx context.Context, id int64) model.CoverageModel
	FindAll(ctx context.Context) []model.CoverageModel
}
