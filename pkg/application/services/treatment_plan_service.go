package services

import (
	"context"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type TreatmentPlanService struct {
	treatmentPlanRepo repositories.TreatmentPlanRepository
	patientRepo       repositories.PatientRepository
	doctorRepo        repositories.DoctorRepository
}

func NewTreatmentPlanService(
	treatmentPlanRepo repositories.TreatmentPlanRepository,
	patientRepo repositories.PatientRepository,
	doctorRepo repositories.DoctorRepository,
) *TreatmentPlanService {
	return &TreatmentPlanService{
		treatmentPlanRepo: treatmentPlanRepo,
		patientRepo:       patientRepo,
		doctorRepo:        doctorRepo,
	}
}

func (s *TreatmentPlanService) CreateTreatmentPlan(ctx context.Context, plan *entities.TreatmentPlan) error {
	if err := plan.Validate(); err != nil {
		return err
	}

	// Verify patient exists
	if _, err := s.patientRepo.GetByID(ctx, plan.PatientID); err != nil {
		return entities.ErrNotFound("patient not found")
	}

	// Verify doctor exists
	if _, err := s.doctorRepo.GetByID(ctx, plan.DoctorID); err != nil {
		return entities.ErrNotFound("doctor not found")
	}

	return s.treatmentPlanRepo.Create(ctx, plan)
}

func (s *TreatmentPlanService) GetTreatmentPlan(ctx context.Context, planID int) (*entities.TreatmentPlan, error) {
	if planID == 0 {
		return nil, entities.ErrInvalidInput("plan ID required")
	}

	return s.treatmentPlanRepo.GetByID(ctx, planID)
}

func (s *TreatmentPlanService) UpdateTreatmentPlan(ctx context.Context, plan *entities.TreatmentPlan) error {
	if err := plan.Validate(); err != nil {
		return err
	}

	existing, err := s.treatmentPlanRepo.GetByID(ctx, plan.PlanID)
	if err != nil {
		return err
	}
	if existing == nil {
		return entities.ErrNotFound("treatment plan not found")
	}

	return s.treatmentPlanRepo.Update(ctx, plan)
}

func (s *TreatmentPlanService) UpdatePlanStatus(ctx context.Context, planID int, status string) error {
	if planID == 0 {
		return entities.ErrInvalidInput("plan ID required")
	}

	plan, err := s.treatmentPlanRepo.GetByID(ctx, planID)
	if err != nil {
		return err
	}

	validStatuses := map[string]bool{"active": true, "completed": true, "cancelled": true}
	if !validStatuses[status] {
		return entities.ErrInvalidInput("invalid status")
	}

	plan.Status = status
	return s.treatmentPlanRepo.Update(ctx, plan)
}

func (s *TreatmentPlanService) GetPatientPlans(ctx context.Context, patientID int) ([]entities.TreatmentPlan, error) {
	if patientID == 0 {
		return nil, entities.ErrInvalidInput("patient ID required")
	}

	return s.treatmentPlanRepo.GetByPatientID(ctx, patientID)
}

func (s *TreatmentPlanService) GetActivePlans(ctx context.Context, patientID int) ([]entities.TreatmentPlan, error) {
	if patientID == 0 {
		return nil, entities.ErrInvalidInput("patient ID required")
	}

	return s.treatmentPlanRepo.GetActiveByPatient(ctx, patientID)
}

func (s *TreatmentPlanService) GetDoctorPlans(ctx context.Context, doctorID int) ([]entities.TreatmentPlan, error) {
	if doctorID == 0 {
		return nil, entities.ErrInvalidInput("doctor ID required")
	}

	return s.treatmentPlanRepo.GetByDoctorID(ctx, doctorID)
}

func (s *TreatmentPlanService) GetAllActivePlans(ctx context.Context) ([]entities.TreatmentPlan, error) {
	return s.treatmentPlanRepo.GetActivePlans(ctx)
}

func (s *TreatmentPlanService) CalculateProgress(ctx context.Context, planID int, completedSessions int) (float64, error) {
	if planID == 0 {
		return 0, entities.ErrInvalidInput("plan ID required")
	}

	plan, err := s.treatmentPlanRepo.GetByID(ctx, planID)
	if err != nil {
		return 0, err
	}

	return plan.GetProgress(completedSessions), nil
}

func (s *TreatmentPlanService) DeleteTreatmentPlan(ctx context.Context, planID int) error {
	if planID == 0 {
		return entities.ErrInvalidInput("plan ID required")
	}

	return s.treatmentPlanRepo.Delete(ctx, planID)
}
