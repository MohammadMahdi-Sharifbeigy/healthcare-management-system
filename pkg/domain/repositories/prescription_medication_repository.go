package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type PrescriptionMedicationRepository interface {
	// CRUD operations
	Create(ctx context.Context, pm *entities.PrescriptionMedication) error
	GetByID(ctx context.Context, id int) (*entities.PrescriptionMedication, error)
	Update(ctx context.Context, pm *entities.PrescriptionMedication) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.PrescriptionMedication, error)
	GetByPrescriptionID(ctx context.Context, prescriptionID int) ([]entities.PrescriptionMedication, error)
	GetByMedicationID(ctx context.Context, medicationID int) ([]entities.PrescriptionMedication, error)
	DeleteByPrescriptionAndMedication(ctx context.Context, prescriptionID, medicationID int) error
}
