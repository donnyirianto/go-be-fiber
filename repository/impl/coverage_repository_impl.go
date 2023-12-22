package impl

import (
	"context"
	"errors"

	"github.com/donnyirianto/go-be-fiber/entity"
	"github.com/donnyirianto/go-be-fiber/exception"
	"github.com/donnyirianto/go-be-fiber/repository"
	"gorm.io/gorm"
)

type coverageRepositoryImpl struct {
	db *gorm.DB
}

func NewCoverageRepositoryImpl(db *gorm.DB) repository.CoverageRepository {
	return &coverageRepositoryImpl{db: db}
}

func (db *coverageRepositoryImpl) Insert(ctx context.Context, coverage entity.Coverage) (entity.Coverage, error) {
	err := db.db.WithContext(ctx).Create(&coverage).Error
	exception.PanicLogging(err)
	return coverage, nil
}

func (db *coverageRepositoryImpl) Update(ctx context.Context, coverage entity.Coverage) (entity.Coverage, error) {
	err := db.db.WithContext(ctx).Where("id = ?", coverage.Id).Updates(&coverage).Error
	exception.PanicLogging(err)
	return coverage, nil
}

func (db *coverageRepositoryImpl) Delete(ctx context.Context, coverage entity.Coverage) error {
	err := db.db.WithContext(ctx).Delete(&coverage).Error
	exception.PanicLogging(err)
	return nil
}

func (db *coverageRepositoryImpl) FindByID(ctx context.Context, id int64) (entity.Coverage, error) {
	var coverage entity.Coverage
	result := db.db.WithContext(ctx).Unscoped().Where("id = ?", id).First(&coverage)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return entity.Coverage{}, errors.New("COVERAGE NOT FOUND")
	}
	exception.PanicLogging(result.Error)
	return coverage, nil
}

func (db *coverageRepositoryImpl) FindAll(ctx context.Context) ([]entity.Coverage, error) {
	var coverages []entity.Coverage
	if err := db.db.WithContext(ctx).Find(&coverages).Error; err != nil {
		return nil, err
	}
	return coverages, nil
}
