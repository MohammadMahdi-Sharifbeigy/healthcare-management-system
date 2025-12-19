package usecase

import (
	"context"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/dto"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DiseaseUseCase struct {
	service *services.DiseaseService
}

func NewDiseaseUseCase(service *services.DiseaseService) *DiseaseUseCase {
	return &DiseaseUseCase{service: service}
}

// AddNewDisease creates new disease in catalog
func (uc *DiseaseUseCase) AddNewDisease(ctx context.Context, req *dto.CreateDiseaseRequest) (*dto.DiseaseResponse, error) {
	disease := &entities.Disease{
		DiseaseName: req.DiseaseName,
		Category:    req.Category,
		ICDCode:     req.ICDCode,
		Description: req.Description,
	}

	if err := uc.service.CreateDisease(ctx, disease); err != nil {
		return nil, err
	}

	return &dto.DiseaseResponse{
		DiseaseID:   disease.DiseaseID,
		DiseaseName: disease.DiseaseName,
		Category:    disease.Category,
		ICDCode:     disease.ICDCode,
		Description: disease.Description,
	}, nil
}

// GetDiseaseInfo retrieves disease details
func (uc *DiseaseUseCase) GetDiseaseInfo(ctx context.Context, diseaseID int) (*dto.DiseaseResponse, error) {
	disease, err := uc.service.GetDisease(ctx, diseaseID)
	if err != nil {
		return nil, err
	}

	return &dto.DiseaseResponse{
		DiseaseID:   disease.DiseaseID,
		DiseaseName: disease.DiseaseName,
		Category:    disease.Category,
		ICDCode:     disease.ICDCode,
		Description: disease.Description,
	}, nil
}

// SearchDiseaseByName searches disease by name
func (uc *DiseaseUseCase) SearchDiseaseByName(ctx context.Context, query string) ([]*dto.DiseaseListResponse, error) {
	diseases, err := uc.service.SearchByName(ctx, query)
	if err != nil {
		return nil, err
	}

	var responses []*dto.DiseaseListResponse
	for _, d := range diseases {
		responses = append(responses, &dto.DiseaseListResponse{
			DiseaseID:   d.DiseaseID,
			DiseaseName: d.DiseaseName,
			Category:    d.Category,
			ICDCode:     d.ICDCode,
		})
	}

	return responses, nil
}

// GetDiseaseByICDCode retrieves disease by ICD code
func (uc *DiseaseUseCase) GetDiseaseByICDCode(ctx context.Context, icdCode string) (*dto.DiseaseResponse, error) {
	disease, err := uc.service.GetByICDCode(ctx, icdCode)
	if err != nil {
		return nil, err
	}

	return &dto.DiseaseResponse{
		DiseaseID:   disease.DiseaseID,
		DiseaseName: disease.DiseaseName,
		Category:    disease.Category,
		ICDCode:     disease.ICDCode,
		Description: disease.Description,
	}, nil
}

// GetDiseasesByCategory retrieves diseases by category
func (uc *DiseaseUseCase) GetDiseasesByCategory(ctx context.Context, category string) ([]*dto.DiseaseListResponse, error) {
	diseases, err := uc.service.GetByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	var responses []*dto.DiseaseListResponse
	for _, d := range diseases {
		responses = append(responses, &dto.DiseaseListResponse{
			DiseaseID:   d.DiseaseID,
			DiseaseName: d.DiseaseName,
			Category:    d.Category,
			ICDCode:     d.ICDCode,
		})
	}

	return responses, nil
}

// GetAllDiseases retrieves all diseases with pagination
func (uc *DiseaseUseCase) GetAllDiseases(ctx context.Context, limit, offset int) ([]*dto.DiseaseListResponse, error) {
	diseases, err := uc.service.GetAllDiseases(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []*dto.DiseaseListResponse
	for _, d := range diseases {
		responses = append(responses, &dto.DiseaseListResponse{
			DiseaseID:   d.DiseaseID,
			DiseaseName: d.DiseaseName,
			Category:    d.Category,
			ICDCode:     d.ICDCode,
		})
	}

	return responses, nil
}
