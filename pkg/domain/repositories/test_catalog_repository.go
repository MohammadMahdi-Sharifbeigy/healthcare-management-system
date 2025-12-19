package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type TestCatalogRepository interface {
	// CRUD operations
	Create(ctx context.Context, catalog *entities.TestCatalog) error
	GetByID(ctx context.Context, id int) (*entities.TestCatalog, error)
	Update(ctx context.Context, catalog *entities.TestCatalog) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.TestCatalog, error)
	GetByName(ctx context.Context, name string) (*entities.TestCatalog, error)
	GetByCategory(ctx context.Context, category string) ([]entities.TestCatalog, error)
	SearchByName(ctx context.Context, name string) ([]entities.TestCatalog, error)
}
