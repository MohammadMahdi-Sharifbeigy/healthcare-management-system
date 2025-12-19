package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type PatientRepository interface {
	// CRUD operations
	Create(ctx context.Context, patient *entities.Patient) error
	GetByID(ctx context.Context, id int) (*entities.Patient, error)
	Update(ctx context.Context, patient *entities.Patient) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Patient, error)
	GetByEmail(ctx context.Context, email string) (*entities.Patient, error)
	GetByPhone(ctx context.Context, phone string) (*entities.Patient, error)
	SearchByName(ctx context.Context, name string) ([]entities.Patient, error)
}
