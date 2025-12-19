package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type PrescriptionRepository interface {
	// CRUD operations
	Create(ctx context.Context, prescription *entities.Prescription) error
	GetByID(ctx context.Context, id int) (*entities.Prescription, error)
	Update(ctx context.Context, prescription *entities.Prescription) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Prescription, error)
	GetByPatientID(ctx context.Context, patientID int) ([]entities.Prescription, error)
	GetByDoctorID(ctx context.Context, doctorID int) ([]entities.Prescription, error)
	GetActivePrescriptionsByPatient(ctx context.Context, patientID int) ([]entities.Prescription, error)
	GetExpiringPrescriptionsByPatient(ctx context.Context, patientID int, days int) ([]entities.Prescription, error)
}
