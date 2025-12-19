package repositories

import (
	"context"
	"time"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type MedicalTestRepository interface {
	// CRUD operations
	Create(ctx context.Context, test *entities.MedicalTest) error
	GetByID(ctx context.Context, id int) (*entities.MedicalTest, error)
	Update(ctx context.Context, test *entities.MedicalTest) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.MedicalTest, error)
	GetByPatientID(ctx context.Context, patientID int) ([]entities.MedicalTest, error)
	GetByTestCatalogID(ctx context.Context, testCatalogID int) ([]entities.MedicalTest, error)
	GetAbnormalResults(ctx context.Context) ([]entities.MedicalTest, error)
	GetCriticalResults(ctx context.Context) ([]entities.MedicalTest, error)
	GetByPatientAndTestType(ctx context.Context, patientID, testCatalogID int) ([]entities.MedicalTest, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]entities.MedicalTest, error)
	GetByResultStatus(ctx context.Context, status string) ([]entities.MedicalTest, error)
}
