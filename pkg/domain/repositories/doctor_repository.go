package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DoctorRepository interface {
	// CRUD operations
	Create(ctx context.Context, doctor *entities.Doctor) error
	GetByID(ctx context.Context, id int) (*entities.Doctor, error)
	Update(ctx context.Context, doctor *entities.Doctor) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Doctor, error)
	GetByEmail(ctx context.Context, email string) (*entities.Doctor, error)
	GetByLicenseNumber(ctx context.Context, licenseNumber string) (*entities.Doctor, error)
	GetBySpecialization(ctx context.Context, specialization string) ([]entities.Doctor, error)
	GetByDepartment(ctx context.Context, department string) ([]entities.Doctor, error)
}
