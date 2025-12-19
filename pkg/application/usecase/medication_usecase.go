package usecase

import (
	"context"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/dto"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type MedicationUseCase struct {
	service *services.MedicationService
}

func NewMedicationUseCase(service *services.MedicationService) *MedicationUseCase {
	return &MedicationUseCase{service: service}
}

// AddMedicationToCatalog creates new medication in catalog
func (uc *MedicationUseCase) AddMedicationToCatalog(ctx context.Context, req *dto.CreateMedicationRequest) (*dto.MedicationResponse, error) {
	medication := &entities.Medication{
		MedicationName: req.MedicationName,
		GenericName:    req.GenericName,
		Form:           req.Form,
		Strength:       req.Strength,
		Manufacturer:   req.Manufacturer,
		SideEffects:    req.SideEffects,
	}

	if err := uc.service.CreateMedication(ctx, medication); err != nil {
		return nil, err
	}

	return &dto.MedicationResponse{
		MedicationID:   medication.MedicationID,
		MedicationName: medication.MedicationName,
		GenericName:    medication.GenericName,
		Form:           medication.Form,
		Strength:       medication.Strength,
		Manufacturer:   medication.Manufacturer,
		SideEffects:    medication.SideEffects,
	}, nil
}

// GetMedicationInfo retrieves medication details
func (uc *MedicationUseCase) GetMedicationInfo(ctx context.Context, medicationID int) (*dto.MedicationResponse, error) {
	medication, err := uc.service.GetMedication(ctx, medicationID)
	if err != nil {
		return nil, err
	}

	return &dto.MedicationResponse{
		MedicationID:   medication.MedicationID,
		MedicationName: medication.MedicationName,
		GenericName:    medication.GenericName,
		Form:           medication.Form,
		Strength:       medication.Strength,
		Manufacturer:   medication.Manufacturer,
		SideEffects:    medication.SideEffects,
	}, nil
}

// SearchMedicationByName searches medication by name
func (uc *MedicationUseCase) SearchMedicationByName(ctx context.Context, query string) ([]*dto.MedicationListResponse, error) {
	medications, err := uc.service.SearchByName(ctx, query)
	if err != nil {
		return nil, err
	}

	var responses []*dto.MedicationListResponse
	for _, m := range medications {
		responses = append(responses, &dto.MedicationListResponse{
			MedicationID:   m.MedicationID,
			MedicationName: m.MedicationName,
			GenericName:    m.GenericName,
			Form:           m.Form,
			Strength:       m.Strength,
		})
	}

	return responses, nil
}

// GetMedicationByGenericName retrieves medication by generic name
func (uc *MedicationUseCase) GetMedicationByGenericName(ctx context.Context, genericName string) ([]*dto.MedicationListResponse, error) {
	medications, err := uc.service.GetByGenericName(ctx, genericName)
	if err != nil {
		return nil, err
	}

	var responses []*dto.MedicationListResponse
	for _, m := range medications {
		responses = append(responses, &dto.MedicationListResponse{
			MedicationID:   m.MedicationID,
			MedicationName: m.MedicationName,
			GenericName:    m.GenericName,
			Form:           m.Form,
			Strength:       m.Strength,
		})
	}

	return responses, nil
}

// GetMedicationsByForm retrieves medications by form
func (uc *MedicationUseCase) GetMedicationsByForm(ctx context.Context, form string) ([]*dto.MedicationListResponse, error) {
	medications, err := uc.service.GetByForm(ctx, form)
	if err != nil {
		return nil, err
	}

	var responses []*dto.MedicationListResponse
	for _, m := range medications {
		responses = append(responses, &dto.MedicationListResponse{
			MedicationID:   m.MedicationID,
			MedicationName: m.MedicationName,
			GenericName:    m.GenericName,
			Form:           m.Form,
			Strength:       m.Strength,
		})
	}

	return responses, nil
}

// GetAllMedications retrieves all medications with pagination
func (uc *MedicationUseCase) GetAllMedications(ctx context.Context, limit, offset int) ([]*dto.MedicationListResponse, error) {
	medications, err := uc.service.GetAllMedications(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []*dto.MedicationListResponse
	for _, m := range medications {
		responses = append(responses, &dto.MedicationListResponse{
			MedicationID:   m.MedicationID,
			MedicationName: m.MedicationName,
			GenericName:    m.GenericName,
			Form:           m.Form,
			Strength:       m.Strength,
		})
	}

	return responses, nil
}
