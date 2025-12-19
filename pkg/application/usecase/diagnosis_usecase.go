package usecase

import (
	"context"
	"time"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/dto"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DiagnosisUseCase struct {
	service *services.DiagnosisService
}

func NewDiagnosisUseCase(service *services.DiagnosisService) *DiagnosisUseCase {
	return &DiagnosisUseCase{service: service}
}

// CreateNewDiagnosis creates diagnosis for patient
func (uc *DiagnosisUseCase) CreateNewDiagnosis(ctx context.Context, req *dto.CreateDiagnosisRequest) (*dto.DiagnosisResponse, error) {
	diagnosisDate, _ := time.Parse("2006-01-02", req.DiagnosisDate)

	diagnosis := &entities.Diagnosis{
		PatientID:     req.PatientID,
		DoctorID:      req.DoctorID,
		DiseaseID:     req.DiseaseID,
		DiagnosisDate: diagnosisDate,
		Severity:      req.Severity,
		Notes:         req.Notes,
	}

	if err := uc.service.CreateDiagnosis(ctx, diagnosis); err != nil {
		return nil, err
	}

	return &dto.DiagnosisResponse{
		DiagnosisID:   diagnosis.DiagnosisID,
		PatientID:     diagnosis.PatientID,
		DoctorID:      diagnosis.DoctorID,
		DiseaseID:     diagnosis.DiseaseID,
		DiagnosisDate: diagnosis.DiagnosisDate,
		Severity:      diagnosis.Severity,
		Notes:         diagnosis.Notes,
		CreatedAt:     diagnosis.CreatedAt,
		UpdatedAt:     diagnosis.UpdatedAt,
	}, nil
}

// GetDiagnosisDetails retrieves diagnosis information
func (uc *DiagnosisUseCase) GetDiagnosisDetails(ctx context.Context, diagnosisID int) (*dto.DiagnosisResponse, error) {
	diagnosis, err := uc.service.GetDiagnosis(ctx, diagnosisID)
	if err != nil {
		return nil, err
	}

	return &dto.DiagnosisResponse{
		DiagnosisID:   diagnosis.DiagnosisID,
		PatientID:     diagnosis.PatientID,
		DoctorID:      diagnosis.DoctorID,
		DiseaseID:     diagnosis.DiseaseID,
		DiagnosisDate: diagnosis.DiagnosisDate,
		Severity:      diagnosis.Severity,
		Notes:         diagnosis.Notes,
		CreatedAt:     diagnosis.CreatedAt,
		UpdatedAt:     diagnosis.UpdatedAt,
	}, nil
}

// UpdateDiagnosisInfo updates diagnosis details
func (uc *DiagnosisUseCase) UpdateDiagnosisInfo(ctx context.Context, diagnosisID int, req *dto.UpdateDiagnosisRequest) (*dto.DiagnosisResponse, error) {
	diagnosis, err := uc.service.GetDiagnosis(ctx, diagnosisID)
	if err != nil {
		return nil, err
	}

	if req.Severity != "" {
		diagnosis.Severity = req.Severity
	}
	if req.Notes != "" {
		diagnosis.Notes = req.Notes
	}

	if err := uc.service.UpdateDiagnosis(ctx, diagnosis); err != nil {
		return nil, err
	}

	return &dto.DiagnosisResponse{
		DiagnosisID:   diagnosis.DiagnosisID,
		PatientID:     diagnosis.PatientID,
		DoctorID:      diagnosis.DoctorID,
		DiseaseID:     diagnosis.DiseaseID,
		DiagnosisDate: diagnosis.DiagnosisDate,
		Severity:      diagnosis.Severity,
		Notes:         diagnosis.Notes,
		CreatedAt:     diagnosis.CreatedAt,
		UpdatedAt:     diagnosis.UpdatedAt,
	}, nil
}

// RemoveDiagnosis deletes a diagnosis
func (uc *DiagnosisUseCase) RemoveDiagnosis(ctx context.Context, diagnosisID int) error {
	return uc.service.DeleteDiagnosis(ctx, diagnosisID)
}

// GetPatientDiagnoses retrieves all diagnoses for a patient
func (uc *DiagnosisUseCase) GetPatientDiagnoses(ctx context.Context, patientID int) ([]*dto.DiagnosisListResponse, error) {
	diagnoses, err := uc.service.GetPatientDiagnoses(ctx, patientID)
	if err != nil {
		return nil, err
	}

	var responses []*dto.DiagnosisListResponse
	for _, d := range diagnoses {
		responses = append(responses, &dto.DiagnosisListResponse{
			DiagnosisID:   d.DiagnosisID,
			DiagnosisDate: d.DiagnosisDate,
			Severity:      d.Severity,
		})
	}

	return responses, nil
}

// GetSevereDiagnosesAlert retrieves severe diagnoses
func (uc *DiagnosisUseCase) GetSevereDiagnosesAlert(ctx context.Context, patientID int) ([]*dto.DiagnosisResponse, error) {
	diagnoses, err := uc.service.GetSevereDiagnoses(ctx, patientID)
	if err != nil {
		return nil, err
	}

	var responses []*dto.DiagnosisResponse
	for _, d := range diagnoses {
		responses = append(responses, &dto.DiagnosisResponse{
			DiagnosisID:   d.DiagnosisID,
			PatientID:     d.PatientID,
			DoctorID:      d.DoctorID,
			DiseaseID:     d.DiseaseID,
			DiagnosisDate: d.DiagnosisDate,
			Severity:      d.Severity,
			Notes:         d.Notes,
			CreatedAt:     d.CreatedAt,
			UpdatedAt:     d.UpdatedAt,
		})
	}

	return responses, nil
}
