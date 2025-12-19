package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type TreatmentPlanRepository interface {
	// CRUD operations
	Create(ctx context.Context, plan *entities.TreatmentPlan) error
	GetByID(ctx context.Context, id int) (*entities.TreatmentPlan, error)
	Update(ctx context.Context, plan *entities.TreatmentPlan) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.TreatmentPlan, error)
	GetByPatientID(ctx context.Context, patientID int) ([]entities.TreatmentPlan, error)
	GetByDoctorID(ctx context.Context, doctorID int) ([]entities.TreatmentPlan, error)
	GetActiveByPatient(ctx context.Context, patientID int) ([]entities.TreatmentPlan, error)
	GetByStatus(ctx context.Context, status string) ([]entities.TreatmentPlan, error)
	GetActivePlans(ctx context.Context) ([]entities.TreatmentPlan, error)
}
