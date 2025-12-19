package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type MedicationRepository interface {
	// CRUD operations
	Create(ctx context.Context, medication *entities.Medication) error
	GetByID(ctx context.Context, id int) (*entities.Medication, error)
	Update(ctx context.Context, medication *entities.Medication) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Medication, error)
	GetByName(ctx context.Context, name string) (*entities.Medication, error)
	GetByGenericName(ctx context.Context, genericName string) (*entities.Medication, error)
	SearchByName(ctx context.Context, name string) ([]entities.Medication, error)
	GetByForm(ctx context.Context, form string) ([]entities.Medication, error)
}
