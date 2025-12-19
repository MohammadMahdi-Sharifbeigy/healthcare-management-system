package usecase

import (
	"context"
	"time"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/dto"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type TreatmentPlanUseCase struct {
	service *services.TreatmentPlanService
}

func NewTreatmentPlanUseCase(service *services.TreatmentPlanService) *TreatmentPlanUseCase {
	return &TreatmentPlanUseCase{service: service}
}

// CreateTreatmentPlan creates new treatment plan for patient
func (uc *TreatmentPlanUseCase) CreateTreatmentPlan(ctx context.Context, req *dto.CreateTreatmentPlanRequest) (*dto.TreatmentPlanResponse, error) {
	startDate, _ := time.Parse("2006-01-02", req.StartDate)

	var endDate *time.Time
	if req.EndDate != "" {
		ed, _ := time.Parse("2006-01-02", req.EndDate)
		endDate = &ed
	}

	plan := &entities.TreatmentPlan{
		PatientID:       req.PatientID,
		DoctorID:        req.DoctorID,
		Diagnosis:       req.Diagnosis,
		TreatmentType:   req.TreatmentType,
		StartDate:       startDate,
		EndDate:         endDate,
		SessionDuration: req.SessionDuration,
		Status:          "active",
		Goals:           req.Goals,
	}

	if err := uc.service.CreateTreatmentPlan(ctx, plan); err != nil {
		return nil, err
	}

	return &dto.TreatmentPlanResponse{
		PlanID:          plan.PlanID,
		PatientID:       plan.PatientID,
		DoctorID:        plan.DoctorID,
		Diagnosis:       plan.Diagnosis,
		TreatmentType:   plan.TreatmentType,
		StartDate:       plan.StartDate,
		EndDate:         plan.EndDate,
		SessionDuration: plan.SessionDuration,
		Status:          plan.Status,
		Goals:           plan.Goals,
		CreatedAt:       plan.CreatedAt,
		UpdatedAt:       plan.UpdatedAt,
	}, nil
}

// GetTreatmentPlanDetails retrieves plan details
func (uc *TreatmentPlanUseCase) GetTreatmentPlanDetails(ctx context.Context, planID int) (*dto.TreatmentPlanResponse, error) {
	plan, err := uc.service.GetTreatmentPlan(ctx, planID)
	if err != nil {
		return nil, err
	}

	return &dto.TreatmentPlanResponse{
		PlanID:          plan.PlanID,
		PatientID:       plan.PatientID,
		DoctorID:        plan.DoctorID,
		Diagnosis:       plan.Diagnosis,
		TreatmentType:   plan.TreatmentType,
		StartDate:       plan.StartDate,
		EndDate:         plan.EndDate,
		SessionDuration: plan.SessionDuration,
		Status:          plan.Status,
		Goals:           plan.Goals,
		ProgressNotes:   plan.ProgressNotes,
		CreatedAt:       plan.CreatedAt,
		UpdatedAt:       plan.UpdatedAt,
	}, nil
}

// UpdateTreatmentPlan updates plan details
func (uc *TreatmentPlanUseCase) UpdateTreatmentPlan(ctx context.Context, planID int, req *dto.UpdateTreatmentPlanRequest) (*dto.TreatmentPlanResponse, error) {
	plan, err := uc.service.GetTreatmentPlan(ctx, planID)
	if err != nil {
		return nil, err
	}

	if req.Diagnosis != "" {
		plan.Diagnosis = req.Diagnosis
	}
	if req.TreatmentType != "" {
		plan.TreatmentType = req.TreatmentType
	}
	if req.EndDate != "" {
		ed, _ := time.Parse("2006-01-02", req.EndDate)
		plan.EndDate = &ed
	}
	if req.SessionDuration > 0 {
		plan.SessionDuration = req.SessionDuration
	}
	if req.Goals != "" {
		plan.Goals = req.Goals
	}
	if req.ProgressNotes != "" {
		plan.ProgressNotes = req.ProgressNotes
	}

	if err := uc.service.UpdateTreatmentPlan(ctx, plan); err != nil {
		return nil, err
	}

	return &dto.TreatmentPlanResponse{
		PlanID:          plan.PlanID,
		PatientID:       plan.PatientID,
		DoctorID:        plan.DoctorID,
		Diagnosis:       plan.Diagnosis,
		TreatmentType:   plan.TreatmentType,
		StartDate:       plan.StartDate,
		EndDate:         plan.EndDate,
		SessionDuration: plan.SessionDuration,
		Status:          plan.Status,
		Goals:           plan.Goals,
		ProgressNotes:   plan.ProgressNotes,
		CreatedAt:       plan.CreatedAt,
		UpdatedAt:       plan.UpdatedAt,
	}, nil
}

// UpdatePlanStatus changes plan status
func (uc *TreatmentPlanUseCase) UpdatePlanStatus(ctx context.Context, planID int, status string) error {
	return uc.service.UpdatePlanStatus(ctx, planID, status)
}

// GetPatientPlans retrieves all treatment plans for patient
func (uc *TreatmentPlanUseCase) GetPatientPlans(ctx context.Context, patientID int) ([]*dto.TreatmentPlanListResponse, error) {
	plans, err := uc.service.GetPatientPlans(ctx, patientID)
	if err != nil {
		return nil, err
	}

	var responses []*dto.TreatmentPlanListResponse
	for _, p := range plans {
		responses = append(responses, &dto.TreatmentPlanListResponse{
			PlanID:        p.PlanID,
			TreatmentType: p.TreatmentType,
			StartDate:     p.StartDate,
			Status:        p.Status,
		})
	}

	return responses, nil
}

// GetActivePlans retrieves active treatment plans
func (uc *TreatmentPlanUseCase) GetActivePlans(ctx context.Context, patientID int) ([]*dto.TreatmentPlanListResponse, error) {
	plans, err := uc.service.GetActivePlans(ctx, patientID)
	if err != nil {
		return nil, err
	}

	var responses []*dto.TreatmentPlanListResponse
	for _, p := range plans {
		responses = append(responses, &dto.TreatmentPlanListResponse{
			PlanID:        p.PlanID,
			TreatmentType: p.TreatmentType,
			StartDate:     p.StartDate,
			Status:        p.Status,
		})
	}

	return responses, nil
}

// GetPlanProgress calculates plan completion percentage
func (uc *TreatmentPlanUseCase) GetPlanProgress(ctx context.Context, planID int, completedSessions int) (float64, error) {
	return uc.service.CalculateProgress(ctx, planID, completedSessions)
}
