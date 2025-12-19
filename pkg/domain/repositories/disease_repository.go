package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DiseaseRepository interface {
	// CRUD operations
	Create(ctx context.Context, disease *entities.Disease) error
	GetByID(ctx context.Context, id int) (*entities.Disease, error)
	Update(ctx context.Context, disease *entities.Disease) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Disease, error)
	GetByName(ctx context.Context, name string) (*entities.Disease, error)
	GetByICDCode(ctx context.Context, icdCode string) (*entities.Disease, error)
	GetByCategory(ctx context.Context, category string) ([]entities.Disease, error)
	SearchByName(ctx context.Context, name string) ([]entities.Disease, error)
}
