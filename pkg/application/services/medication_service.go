package services

import (
	"context"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type MedicationService struct {
	repo repositories.MedicationRepository
}

func NewMedicationService(repo repositories.MedicationRepository) *MedicationService {
	return &MedicationService{repo: repo}
}

func (s *MedicationService) CreateMedication(ctx context.Context, medication *entities.Medication) error {
	if err := medication.Validate(); err != nil {
		return err
	}

	return s.repo.Create(ctx, medication)
}

func (s *MedicationService) GetMedication(ctx context.Context, medicationID int) (*entities.Medication, error) {
	if medicationID == 0 {
		return nil, entities.ErrInvalidInput("medication ID required")
	}

	return s.repo.GetByID(ctx, medicationID)
}

func (s *MedicationService) SearchByName(ctx context.Context, name string) ([]entities.Medication, error) {
	if name == "" {
		return nil, entities.ErrInvalidInput("medication name required")
	}

	return s.repo.SearchByName(ctx, name)
}

func (s *MedicationService) GetByGenericName(ctx context.Context, genericName string) ([]entities.Medication, error) {
	if genericName == "" {
		return nil, entities.ErrInvalidInput("generic name required")
	}

	return s.repo.GetByGenericName(ctx, genericName)
}

func (s *MedicationService) GetByForm(ctx context.Context, form string) ([]entities.Medication, error) {
	if form == "" {
		return nil, entities.ErrInvalidInput("form required")
	}

	return s.repo.GetByForm(ctx, form)
}

func (s *MedicationService) GetAllMedications(ctx context.Context, limit, offset int) ([]entities.Medication, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	return s.repo.GetAll(ctx, limit, offset)
}

func (s *MedicationService) UpdateMedication(ctx context.Context, medication *entities.Medication) error {
	if err := medication.Validate(); err != nil {
		return err
	}

	existing, err := s.repo.GetByID(ctx, medication.MedicationID)
	if err != nil {
		return err
	}
	if existing == nil {
		return entities.ErrNotFound("medication not found")
	}

	return s.repo.Update(ctx, medication)
}
