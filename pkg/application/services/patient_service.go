package services

import (
	"context"
	"fmt"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type PatientService struct {
	repo repositories.PatientRepository
}

func NewPatientService(repo repositories.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func (s *PatientService) RegisterPatient(ctx context.Context, patient *entities.Patient) error {
	if err := patient.Validate(); err != nil {
		return err
	}

	existing, _ := s.repo.GetByEmail(ctx, patient.Email)
	if existing != nil {
		return entities.ErrConflict("patient with this email already exists")
	}

	return s.repo.Create(ctx, patient)
}

func (s *PatientService) GetPatientProfile(ctx context.Context, patientID int) (*entities.Patient, error) {
	if patientID == 0 {
		return nil, entities.ErrInvalidInput("patient ID required")
	}

	return s.repo.GetByID(ctx, patientID)
}

func (s *PatientService) UpdatePatientInfo(ctx context.Context, patient *entities.Patient) error {
	if err := patient.Validate(); err != nil {
		return err
	}

	existing, _ := s.repo.GetByID(ctx, patient.PatientID)
	if existing == nil {
		return entities.ErrNotFound("patient not found")
	}

	return s.repo.Update(ctx, patient)
}

func (s *PatientService) DeletePatient(ctx context.Context, patientID int) error {
	if patientID == 0 {
		return entities.ErrInvalidInput("patient ID required")
	}

	existing, _ := s.repo.GetByID(ctx, patientID)
	if existing == nil {
		return entities.ErrNotFound("patient not found")
	}

	return s.repo.Delete(ctx, patientID)
}

func (s *PatientService) SearchPatients(ctx context.Context, query string) ([]entities.Patient, error) {
	if query == "" {
		return nil, entities.ErrInvalidInput("search query required")
	}

	return s.repo.SearchByName(ctx, query)
}

func (s *PatientService) GetPatientsByPage(ctx context.Context, limit, offset int) ([]entities.Patient, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	return s.repo.GetAll(ctx, limit, offset)
}

func (s *PatientService) ValidatePatientAge(ctx context.Context, patientID int) (bool, error) {
	patient, err := s.repo.GetByID(ctx, patientID)
	if err != nil {
		return false, err
	}

	return patient.IsAdult(), nil
}
