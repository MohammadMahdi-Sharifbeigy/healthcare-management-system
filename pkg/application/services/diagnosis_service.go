package services

import (
	"context"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type DiagnosisService struct {
	diagnosisRepo repositories.DiagnosisRepository
	patientRepo   repositories.PatientRepository
	doctorRepo    repositories.DoctorRepository
	diseaseRepo   repositories.DiseaseRepository
}

func NewDiagnosisService(
	diagnosisRepo repositories.DiagnosisRepository,
	patientRepo repositories.PatientRepository,
	doctorRepo repositories.DoctorRepository,
	diseaseRepo repositories.DiseaseRepository,
) *DiagnosisService {
	return &DiagnosisService{
		diagnosisRepo: diagnosisRepo,
		patientRepo:   patientRepo,
		doctorRepo:    doctorRepo,
		diseaseRepo:   diseaseRepo,
	}
}

func (s *DiagnosisService) CreateDiagnosis(ctx context.Context, diagnosis *entities.Diagnosis) error {
	if err := diagnosis.Validate(); err != nil {
		return err
	}

	// Verify patient exists
	if _, err := s.patientRepo.GetByID(ctx, diagnosis.PatientID); err != nil {
		return entities.ErrNotFound("patient not found")
	}

	// Verify doctor exists
	if _, err := s.doctorRepo.GetByID(ctx, diagnosis.DoctorID); err != nil {
		return entities.ErrNotFound("doctor not found")
	}

	// Verify disease exists
	if _, err := s.diseaseRepo.GetByID(ctx, diagnosis.DiseaseID); err != nil {
		return entities.ErrNotFound("disease not found")
	}

	return s.diagnosisRepo.Create(ctx, diagnosis)
}

func (s *DiagnosisService) GetDiagnosis(ctx context.Context, diagnosisID int) (*entities.Diagnosis, error) {
	if diagnosisID == 0 {
		return nil, entities.ErrInvalidInput("diagnosis ID required")
	}

	return s.diagnosisRepo.GetByID(ctx, diagnosisID)
}

func (s *DiagnosisService) UpdateDiagnosis(ctx context.Context, diagnosis *entities.Diagnosis) error {
	if err := diagnosis.Validate(); err != nil {
		return err
	}

	existing, err := s.diagnosisRepo.GetByID(ctx, diagnosis.DiagnosisID)
	if err != nil {
		return err
	}
	if existing == nil {
		return entities.ErrNotFound("diagnosis not found")
	}

	return s.diagnosisRepo.Update(ctx, diagnosis)
}

func (s *DiagnosisService) DeleteDiagnosis(ctx context.Context, diagnosisID int) error {
	if diagnosisID == 0 {
		return entities.ErrInvalidInput("diagnosis ID required")
	}

	return s.diagnosisRepo.Delete(ctx, diagnosisID)
}

func (s *DiagnosisService) GetPatientDiagnoses(ctx context.Context, patientID int) ([]entities.Diagnosis, error) {
	if patientID == 0 {
		return nil, entities.ErrInvalidInput("patient ID required")
	}

	return s.diagnosisRepo.GetByPatientID(ctx, patientID)
}

func (s *DiagnosisService) GetSevereDiagnoses(ctx context.Context, patientID int) ([]entities.Diagnosis, error) {
	if patientID == 0 {
		return nil, entities.ErrInvalidInput("patient ID required")
	}

	return s.diagnosisRepo.GetSevereForPatient(ctx, patientID)
}

func (s *DiagnosisService) GetDoctorDiagnoses(ctx context.Context, doctorID int) ([]entities.Diagnosis, error) {
	if doctorID == 0 {
		return nil, entities.ErrInvalidInput("doctor ID required")
	}

	return s.diagnosisRepo.GetByDoctorID(ctx, doctorID)
}
