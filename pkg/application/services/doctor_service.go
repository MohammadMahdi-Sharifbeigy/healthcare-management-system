package services

import (
	"context"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type DoctorService struct {
	repo repositories.DoctorRepository
}

func NewDoctorService(repo repositories.DoctorRepository) *DoctorService {
	return &DoctorService{repo: repo}
}

func (s *DoctorService) RegisterDoctor(ctx context.Context, doctor *entities.Doctor) error {
	if err := doctor.Validate(); err != nil {
		return err
	}

	existing, _ := s.repo.GetByEmail(ctx, doctor.Email)
	if existing != nil {
		return entities.ErrConflict("doctor with this email already exists")
	}

	existingLicense, _ := s.repo.GetByLicenseNumber(ctx, doctor.LicenseNumber)
	if existingLicense != nil {
		return entities.ErrConflict("doctor with this license number already exists")
	}

	return s.repo.Create(ctx, doctor)
}

func (s *DoctorService) GetDoctorProfile(ctx context.Context, doctorID int) (*entities.Doctor, error) {
	if doctorID == 0 {
		return nil, entities.ErrInvalidInput("doctor ID required")
	}

	return s.repo.GetByID(ctx, doctorID)
}

func (s *DoctorService) UpdateDoctorInfo(ctx context.Context, doctor *entities.Doctor) error {
	if err := doctor.Validate(); err != nil {
		return err
	}

	existing, _ := s.repo.GetByID(ctx, doctor.DoctorID)
	if existing == nil {
		return entities.ErrNotFound("doctor not found")
	}

	return s.repo.Update(ctx, doctor)
}

func (s *DoctorService) GetDoctorsBySpecialization(ctx context.Context, specialization string) ([]entities.Doctor, error) {
	if specialization == "" {
		return nil, entities.ErrInvalidInput("specialization required")
	}

	return s.repo.GetBySpecialization(ctx, specialization)
}

func (s *DoctorService) GetDoctorsByDepartment(ctx context.Context, department string) ([]entities.Doctor, error) {
	if department == "" {
		return nil, entities.ErrInvalidInput("department required")
	}

	return s.repo.GetByDepartment(ctx, department)
}

func (s *DoctorService) GetAllDoctors(ctx context.Context, limit, offset int) ([]entities.Doctor, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	return s.repo.GetAll(ctx, limit, offset)
}
