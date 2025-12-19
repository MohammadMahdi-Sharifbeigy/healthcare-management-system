package services

import (
	"context"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type DiseaseService struct {
	repo repositories.DiseaseRepository
}

func NewDiseaseService(repo repositories.DiseaseRepository) *DiseaseService {
	return &DiseaseService{repo: repo}
}

func (s *DiseaseService) CreateDisease(ctx context.Context, disease *entities.Disease) error {
	if err := disease.Validate(); err != nil {
		return err
	}

	return s.repo.Create(ctx, disease)
}

func (s *DiseaseService) GetDisease(ctx context.Context, diseaseID int) (*entities.Disease, error) {
	if diseaseID == 0 {
		return nil, entities.ErrInvalidInput("disease ID required")
	}

	return s.repo.GetByID(ctx, diseaseID)
}

func (s *DiseaseService) SearchByName(ctx context.Context, name string) ([]entities.Disease, error) {
	if name == "" {
		return nil, entities.ErrInvalidInput("disease name required")
	}

	return s.repo.SearchByName(ctx, name)
}

func (s *DiseaseService) GetByICDCode(ctx context.Context, icdCode string) (*entities.Disease, error) {
	if icdCode == "" {
		return nil, entities.ErrInvalidInput("ICD code required")
	}

	return s.repo.GetByICDCode(ctx, icdCode)
}

func (s *DiseaseService) GetByCategory(ctx context.Context, category string) ([]entities.Disease, error) {
	if category == "" {
		return nil, entities.ErrInvalidInput("category required")
	}

	return s.repo.GetByCategory(ctx, category)
}

func (s *DiseaseService) GetAllDiseases(ctx context.Context, limit, offset int) ([]entities.Disease, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	return s.repo.GetAll(ctx, limit, offset)
}

func (s *DiseaseService) UpdateDisease(ctx context.Context, disease *entities.Disease) error {
	if err := disease.Validate(); err != nil {
		return err
	}

	existing, err := s.repo.GetByID(ctx, disease.DiseaseID)
	if err != nil {
		return err
	}
	if existing == nil {
		return entities.ErrNotFound("disease not found")
	}

	return s.repo.Update(ctx, disease)
}
