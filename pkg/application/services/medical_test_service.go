package services

import (
	"context"
	"time"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type MedicalTestService struct {
	medicalTestRepo repositories.MedicalTestRepository
	testCatalogRepo repositories.TestCatalogRepository
	patientRepo     repositories.PatientRepository
	doctorRepo      repositories.DoctorRepository
}

func NewMedicalTestService(
	medicalTestRepo repositories.MedicalTestRepository,
	testCatalogRepo repositories.TestCatalogRepository,
	patientRepo repositories.PatientRepository,
	doctorRepo repositories.DoctorRepository,
) *MedicalTestService {
	return &MedicalTestService{
		medicalTestRepo: medicalTestRepo,
		testCatalogRepo: testCatalogRepo,
		patientRepo:     patientRepo,
		doctorRepo:      doctorRepo,
	}
}

func (s *MedicalTestService) CreateTest(ctx context.Context, test *entities.MedicalTest) error {
	if err := test.Validate(); err != nil {
		return err
	}

	// Verify patient exists
	if _, err := s.patientRepo.GetByID(ctx, test.PatientID); err != nil {
		return entities.ErrNotFound("patient not found")
	}

	// Verify test catalog exists
	if _, err := s.testCatalogRepo.GetByID(ctx, test.TestCatalogID); err != nil {
		return entities.ErrNotFound("test catalog not found")
	}

	// Verify interpreter exists if provided
	if test.InterpretedBy != nil {
		if _, err := s.doctorRepo.GetByID(ctx, *test.InterpretedBy); err != nil {
			return entities.ErrNotFound("doctor not found")
		}
	}

	return s.medicalTestRepo.Create(ctx, test)
}

func (s *MedicalTestService) GetTest(ctx context.Context, testID int) (*entities.MedicalTest, error) {
	if testID == 0 {
		return nil, entities.ErrInvalidInput("test ID required")
	}

	return s.medicalTestRepo.GetByID(ctx, testID)
}

func (s *MedicalTestService) UpdateTestResults(ctx context.Context, test *entities.MedicalTest) error {
	if err := test.Validate(); err != nil {
		return err
	}

	existing, err := s.medicalTestRepo.GetByID(ctx, test.TestID)
	if err != nil {
		return err
	}
	if existing == nil {
		return entities.ErrNotFound("test not found")
	}

	return s.medicalTestRepo.Update(ctx, test)
}

func (s *MedicalTestService) InterpretTest(ctx context.Context, testID int, doctorID int) error {
	if testID == 0 || doctorID == 0 {
		return entities.ErrInvalidInput("test and doctor IDs required")
	}

	// Verify doctor exists
	if _, err := s.doctorRepo.GetByID(ctx, doctorID); err != nil {
		return entities.ErrNotFound("doctor not found")
	}

	test, err := s.medicalTestRepo.GetByID(ctx, testID)
	if err != nil {
		return err
	}

	test.InterpretedBy = &doctorID
	return s.medicalTestRepo.Update(ctx, test)
}

func (s *MedicalTestService) GetPatientTests(ctx context.Context, patientID int) ([]entities.MedicalTest, error) {
	if patientID == 0 {
		return nil, entities.ErrInvalidInput("patient ID required")
	}

	return s.medicalTestRepo.GetByPatientID(ctx, patientID)
}

func (s *MedicalTestService) GetAbnormalResults(ctx context.Context) ([]entities.MedicalTest, error) {
	return s.medicalTestRepo.GetAbnormalResults(ctx)
}

func (s *MedicalTestService) GetCriticalResults(ctx context.Context) ([]entities.MedicalTest, error) {
	return s.medicalTestRepo.GetCriticalResults(ctx)
}

func (s *MedicalTestService) GetTestsByType(ctx context.Context, patientID int, testCatalogID int) ([]entities.MedicalTest, error) {
	if patientID == 0 || testCatalogID == 0 {
		return nil, entities.ErrInvalidInput("patient and test catalog IDs required")
	}

	return s.medicalTestRepo.GetByPatientAndTestType(ctx, patientID, testCatalogID)
}

func (s *MedicalTestService) GetTestsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]entities.MedicalTest, error) {
	if startDate.IsZero() || endDate.IsZero() {
		return nil, entities.ErrInvalidInput("start and end dates required")
	}

	return s.medicalTestRepo.GetByDateRange(ctx, startDate, endDate)
}
