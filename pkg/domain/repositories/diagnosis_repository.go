package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DiagnosisRepository interface {
	// CRUD operations
	Create(ctx context.Context, diagnosis *entities.Diagnosis) error
	GetByID(ctx context.Context, id int) (*entities.Diagnosis, error)
	Update(ctx context.Context, diagnosis *entities.Diagnosis) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Diagnosis, error)
	GetByPatientID(ctx context.Context, patientID int) ([]entities.Diagnosis, error)
	GetByDoctorID(ctx context.Context, doctorID int) ([]entities.Diagnosis, error)
	GetByDiseaseID(ctx context.Context, diseaseID int) ([]entities.Diagnosis, error)
	GetBySeverity(ctx context.Context, severity string) ([]entities.Diagnosis, error)
	GetSevereForPatient(ctx context.Context, patientID int) ([]entities.Diagnosis, error)
}
