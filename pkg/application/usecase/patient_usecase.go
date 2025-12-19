package usecase

import (
	"context"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/dto"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type PatientUseCase struct {
	service *services.PatientService
}

func NewPatientUseCase(service *services.PatientService) *PatientUseCase {
	return &PatientUseCase{service: service}
}

// RegisterNewPatient handles patient registration
func (uc *PatientUseCase) RegisterNewPatient(ctx context.Context, req *dto.CreatePatientRequest) (*dto.PatientResponse, error) {
	patient := &entities.Patient{
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Email:            req.Email,
		Phone:            req.Phone,
		Address:          req.Address,
		EmergencyContact: req.EmergencyContact,
		BloodType:        req.BloodType,
		MedicalHistory:   req.MedicalHistory,
		Gender:           req.Gender,
	}

	if err := uc.service.RegisterPatient(ctx, patient); err != nil {
		return nil, err
	}

	return &dto.PatientResponse{
		PatientID:        patient.PatientID,
		FirstName:        patient.FirstName,
		LastName:         patient.LastName,
		Email:            patient.Email,
		Phone:            patient.Phone,
		Address:          patient.Address,
		EmergencyContact: patient.EmergencyContact,
		BloodType:        patient.BloodType,
		MedicalHistory:   patient.MedicalHistory,
		Gender:           patient.Gender,
		CreatedAt:        patient.CreatedAt,
		UpdatedAt:        patient.UpdatedAt,
	}, nil
}

// GetPatientProfile retrieves patient details
func (uc *PatientUseCase) GetPatientProfile(ctx context.Context, patientID int) (*dto.PatientResponse, error) {
	patient, err := uc.service.GetPatientProfile(ctx, patientID)
	if err != nil {
		return nil, err
	}

	return &dto.PatientResponse{
		PatientID:        patient.PatientID,
		FirstName:        patient.FirstName,
		LastName:         patient.LastName,
		DateOfBirth:      patient.DateOfBirth,
		Gender:           patient.Gender,
		Email:            patient.Email,
		Phone:            patient.Phone,
		Address:          patient.Address,
		EmergencyContact: patient.EmergencyContact,
		BloodType:        patient.BloodType,
		MedicalHistory:   patient.MedicalHistory,
		CreatedAt:        patient.CreatedAt,
		UpdatedAt:        patient.UpdatedAt,
	}, nil
}

// UpdatePatientInfo updates patient details
func (uc *PatientUseCase) UpdatePatientInfo(ctx context.Context, patientID int, req *dto.UpdatePatientRequest) (*dto.PatientResponse, error) {
	patient, err := uc.service.GetPatientProfile(ctx, patientID)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" {
		patient.FirstName = req.FirstName
	}
	if req.LastName != "" {
		patient.LastName = req.LastName
	}
	if req.Phone != "" {
		patient.Phone = req.Phone
	}
	if req.Address != "" {
		patient.Address = req.Address
	}
	if req.EmergencyContact != "" {
		patient.EmergencyContact = req.EmergencyContact
	}
	if req.MedicalHistory != "" {
		patient.MedicalHistory = req.MedicalHistory
	}

	if err := uc.service.UpdatePatientInfo(ctx, patient); err != nil {
		return nil, err
	}

	return &dto.PatientResponse{
		PatientID:        patient.PatientID,
		FirstName:        patient.FirstName,
		LastName:         patient.LastName,
		Email:            patient.Email,
		Phone:            patient.Phone,
		Address:          patient.Address,
		EmergencyContact: patient.EmergencyContact,
		BloodType:        patient.BloodType,
		MedicalHistory:   patient.MedicalHistory,
		CreatedAt:        patient.CreatedAt,
		UpdatedAt:        patient.UpdatedAt,
	}, nil
}

// DeletePatient soft deletes a patient
func (uc *PatientUseCase) DeletePatient(ctx context.Context, patientID int) error {
	return uc.service.DeletePatient(ctx, patientID)
}

// SearchPatients searches for patients by name or email
func (uc *PatientUseCase) SearchPatients(ctx context.Context, query string) ([]*dto.PatientListResponse, error) {
	patients, err := uc.service.SearchPatients(ctx, query)
	if err != nil {
		return nil, err
	}

	var responses []*dto.PatientListResponse
	for _, p := range patients {
		responses = append(responses, &dto.PatientListResponse{
			PatientID: p.PatientID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			Email:     p.Email,
			Phone:     p.Phone,
			BloodType: p.BloodType,
		})
	}

	return responses, nil
}

// GetPatientsByPage retrieves paginated patient list
func (uc *PatientUseCase) GetPatientsByPage(ctx context.Context, limit, offset int) ([]*dto.PatientListResponse, error) {
	patients, err := uc.service.GetPatientsByPage(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []*dto.PatientListResponse
	for _, p := range patients {
		responses = append(responses, &dto.PatientListResponse{
			PatientID: p.PatientID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			Email:     p.Email,
			Phone:     p.Phone,
			BloodType: p.BloodType,
		})
	}

	return responses, nil
}
