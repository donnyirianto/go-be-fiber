package repository

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/entity"
)

type CoverageRepository interface {
	// FindAll retrieves all coverages.
	FindAll(ctx context.Context) ([]entity.Coverage, error)

	// Insert adds a new coverage.
	Insert(ctx context.Context, coverage entity.Coverage) (entity.Coverage, error)

	// Update modifies an existing coverage.
	Update(ctx context.Context, coverage entity.Coverage) (entity.Coverage, error)

	// Delete removes a coverage.
	Delete(ctx context.Context, coverage entity.Coverage) error

	// FindByID retrieves a coverage by ID.
	FindByID(ctx context.Context, id int64) (entity.Coverage, error)
}
