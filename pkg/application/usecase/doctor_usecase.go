package usecase

import (
	"context"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/dto"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DoctorUseCase struct {
	service *services.DoctorService
}

func NewDoctorUseCase(service *services.DoctorService) *DoctorUseCase {
	return &DoctorUseCase{service: service}
}

// RegisterNewDoctor creates new doctor profile
func (uc *DoctorUseCase) RegisterNewDoctor(ctx context.Context, req *dto.CreateDoctorRequest) (*dto.DoctorResponse, error) {
	doctor := &entities.Doctor{
		DoctorID:       req.DoctorID,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		LicenseNumber:  req.LicenseNumber,
		Specialization: req.Specialization,
		Department:     req.Department,
	}

	if err := uc.service.RegisterDoctor(ctx, doctor); err != nil {
		return nil, err
	}

	return &dto.DoctorResponse{
		DoctorID:       doctor.DoctorID,
		FirstName:      doctor.FirstName,
		LastName:       doctor.LastName,
		Email:          doctor.Email,
		Phone:          doctor.Phone,
		LicenseNumber:  doctor.LicenseNumber,
		Specialization: doctor.Specialization,
		Department:     doctor.Department,
	}, nil
}

// GetDoctorProfile retrieves doctor details
func (uc *DoctorUseCase) GetDoctorProfile(ctx context.Context, doctorID int) (*dto.DoctorResponse, error) {
	doctor, err := uc.service.GetDoctorProfile(ctx, doctorID)
	if err != nil {
		return nil, err
	}

	return &dto.DoctorResponse{
		DoctorID:       doctor.DoctorID,
		FirstName:      doctor.FirstName,
		LastName:       doctor.LastName,
		Email:          doctor.Email,
		Phone:          doctor.Phone,
		LicenseNumber:  doctor.LicenseNumber,
		Specialization: doctor.Specialization,
		Department:     doctor.Department,
	}, nil
}

// UpdateDoctorInfo updates doctor information
func (uc *DoctorUseCase) UpdateDoctorInfo(ctx context.Context, doctorID int, req *dto.UpdateDoctorRequest) (*dto.DoctorResponse, error) {
	doctor, err := uc.service.GetDoctorProfile(ctx, doctorID)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" {
		doctor.FirstName = req.FirstName
	}
	if req.LastName != "" {
		doctor.LastName = req.LastName
	}
	if req.Email != "" {
		doctor.Email = req.Email
	}
	if req.Phone != "" {
		doctor.Phone = req.Phone
	}
	if req.Specialization != "" {
		doctor.Specialization = req.Specialization
	}
	if req.Department != "" {
		doctor.Department = req.Department
	}

	if err := uc.service.UpdateDoctorInfo(ctx, doctor); err != nil {
		return nil, err
	}

	return &dto.DoctorResponse{
		DoctorID:       doctor.DoctorID,
		FirstName:      doctor.FirstName,
		LastName:       doctor.LastName,
		Email:          doctor.Email,
		Phone:          doctor.Phone,
		LicenseNumber:  doctor.LicenseNumber,
		Specialization: doctor.Specialization,
		Department:     doctor.Department,
	}, nil
}

// GetDoctorsBySpecialty retrieves doctors by specialization
func (uc *DoctorUseCase) GetDoctorsBySpecialty(ctx context.Context, specialization string) ([]*dto.DoctorListResponse, error) {
	doctors, err := uc.service.GetDoctorsBySpecialization(ctx, specialization)
	if err != nil {
		return nil, err
	}

	var responses []*dto.DoctorListResponse
	for _, d := range doctors {
		responses = append(responses, &dto.DoctorListResponse{
			DoctorID:       d.DoctorID,
			FirstName:      d.FirstName,
			LastName:       d.LastName,
			Specialization: d.Specialization,
			Department:     d.Department,
			Email:          d.Email,
		})
	}

	return responses, nil
}

// GetDoctorsByDept retrieves doctors by department
func (uc *DoctorUseCase) GetDoctorsByDept(ctx context.Context, department string) ([]*dto.DoctorListResponse, error) {
	doctors, err := uc.service.GetDoctorsByDepartment(ctx, department)
	if err != nil {
		return nil, err
	}

	var responses []*dto.DoctorListResponse
	for _, d := range doctors {
		responses = append(responses, &dto.DoctorListResponse{
			DoctorID:       d.DoctorID,
			FirstName:      d.FirstName,
			LastName:       d.LastName,
			Specialization: d.Specialization,
			Department:     d.Department,
			Email:          d.Email,
		})
	}

	return responses, nil
}

// GetAllDoctors retrieves all doctors with pagination
func (uc *DoctorUseCase) GetAllDoctors(ctx context.Context, limit, offset int) ([]*dto.DoctorListResponse, error) {
	doctors, err := uc.service.GetAllDoctors(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []*dto.DoctorListResponse
	for _, d := range doctors {
		responses = append(responses, &dto.DoctorListResponse{
			DoctorID:       d.DoctorID,
			FirstName:      d.FirstName,
			LastName:       d.LastName,
			Specialization: d.Specialization,
			Department:     d.Department,
			Email:          d.Email,
		})
	}

	return responses, nil
}
