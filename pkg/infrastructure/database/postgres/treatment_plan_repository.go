package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type treatmentPlanRepository struct {
	db *sql.DB
}

func NewTreatmentPlanRepository(db *sql.DB) repositories.TreatmentPlanRepository {
	return &treatmentPlanRepository{db: db}
}

func (r *treatmentPlanRepository) Create(ctx context.Context, plan *entities.TreatmentPlan) error {
	if err := plan.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO treatment_plan (patient_id, doctor_id, diagnosis, treatment_type, start_date, end_date, session_duration, status, goals, progress_notes)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		plan.PatientID,
		plan.DoctorID,
		plan.Diagnosis,
		plan.TreatmentType,
		plan.StartDate,
		plan.EndDate,
		plan.SessionDuration,
		plan.Status,
		plan.Goals,
		plan.ProgressNotes,
	)

	if err != nil {
		return fmt.Errorf("failed to create treatment plan: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get treatment plan id: %w", err)
	}

	plan.PlanID = int(id)
	return nil
}

func (r *treatmentPlanRepository) GetByID(ctx context.Context, id int) (*entities.TreatmentPlan, error) {
	query := `
		SELECT plan_id, patient_id, doctor_id, diagnosis, treatment_type, start_date, end_date, session_duration, status, goals, progress_notes
		FROM treatment_plan
		WHERE plan_id = ?
	`

	plan := &entities.TreatmentPlan{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&plan.PlanID,
		&plan.PatientID,
		&plan.DoctorID,
		&plan.Diagnosis,
		&plan.TreatmentType,
		&plan.StartDate,
		&plan.EndDate,
		&plan.SessionDuration,
		&plan.Status,
		&plan.Goals,
		&plan.ProgressNotes,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("treatment plan not found")
		}
		return nil, fmt.Errorf("failed to get treatment plan: %w", err)
	}

	return plan, nil
}

func (r *treatmentPlanRepository) Update(ctx context.Context, plan *entities.TreatmentPlan) error {
	if err := plan.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE treatment_plan
		SET patient_id = ?, doctor_id = ?, diagnosis = ?, treatment_type = ?, start_date = ?, end_date = ?, session_duration = ?, status = ?, goals = ?, progress_notes = ?
		WHERE plan_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		plan.PatientID,
		plan.DoctorID,
		plan.Diagnosis,
		plan.TreatmentType,
		plan.StartDate,
		plan.EndDate,
		plan.SessionDuration,
		plan.Status,
		plan.Goals,
		plan.ProgressNotes,
		plan.PlanID,
	)

	if err != nil {
		return fmt.Errorf("failed to update treatment plan: %w", err)
	}

	return nil
}

func (r *treatmentPlanRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM treatment_plan WHERE plan_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete treatment plan: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("treatment plan not found")
	}

	return nil
}

func (r *treatmentPlanRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.TreatmentPlan, error) {
	query := `
		SELECT plan_id, patient_id, doctor_id, diagnosis, treatment_type, start_date, end_date, session_duration, status, goals, progress_notes
		FROM treatment_plan
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get treatment plans: %w", err)
	}
	defer rows.Close()

	var plans []entities.TreatmentPlan
	for rows.Next() {
		var plan entities.TreatmentPlan
		if err := rows.Scan(
			&plan.PlanID,
			&plan.PatientID,
			&plan.DoctorID,
			&plan.Diagnosis,
			&plan.TreatmentType,
			&plan.StartDate,
			&plan.EndDate,
			&plan.SessionDuration,
			&plan.Status,
			&plan.Goals,
			&plan.ProgressNotes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan treatment plan: %w", err)
		}
		plans = append(plans, plan)
	}

	return plans, rows.Err()
}

func (r *treatmentPlanRepository) GetByPatientID(ctx context.Context, patientID int) ([]entities.TreatmentPlan, error) {
	query := `
		SELECT plan_id, patient_id, doctor_id, diagnosis, treatment_type, start_date, end_date, session_duration, status, goals, progress_notes
		FROM treatment_plan
		WHERE patient_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get treatment plans: %w", err)
	}
	defer rows.Close()

	var plans []entities.TreatmentPlan
	for rows.Next() {
		var plan entities.TreatmentPlan
		if err := rows.Scan(
			&plan.PlanID,
			&plan.PatientID,
			&plan.DoctorID,
			&plan.Diagnosis,
			&plan.TreatmentType,
			&plan.StartDate,
			&plan.EndDate,
			&plan.SessionDuration,
			&plan.Status,
			&plan.Goals,
			&plan.ProgressNotes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan treatment plan: %w", err)
		}
		plans = append(plans, plan)
	}

	return plans, rows.Err()
}

func (r *treatmentPlanRepository) GetByDoctorID(ctx context.Context, doctorID int) ([]entities.TreatmentPlan, error) {
	query := `
		SELECT plan_id, patient_id, doctor_id, diagnosis, treatment_type, start_date, end_date, session_duration, status, goals, progress_notes
		FROM treatment_plan
		WHERE doctor_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, doctorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get treatment plans: %w", err)
	}
	defer rows.Close()

	var plans []entities.TreatmentPlan
	for rows.Next() {
		var plan entities.TreatmentPlan
		if err := rows.Scan(
			&plan.PlanID,
			&plan.PatientID,
			&plan.DoctorID,
			&plan.Diagnosis,
			&plan.TreatmentType,
			&plan.StartDate,
			&plan.EndDate,
			&plan.SessionDuration,
			&plan.Status,
			&plan.Goals,
			&plan.ProgressNotes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan treatment plan: %w", err)
		}
		plans = append(plans, plan)
	}

	return plans, rows.Err()
}

func (r *treatmentPlanRepository) GetActiveByPatient(ctx context.Context, patientID int) ([]entities.TreatmentPlan, error) {
	query := `
		SELECT plan_id, patient_id, doctor_id, diagnosis, treatment_type, start_date, end_date, session_duration, status, goals, progress_notes
		FROM treatment_plan
		WHERE patient_id = ? AND status = 'active'
	`

	rows, err := r.db.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get treatment plans: %w", err)
	}
	defer rows.Close()

	var plans []entities.TreatmentPlan
	for rows.Next() {
		var plan entities.TreatmentPlan
		if err := rows.Scan(
			&plan.PlanID,
			&plan.PatientID,
			&plan.DoctorID,
			&plan.Diagnosis,
			&plan.TreatmentType,
			&plan.StartDate,
			&plan.EndDate,
			&plan.SessionDuration,
			&plan.Status,
			&plan.Goals,
			&plan.ProgressNotes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan treatment plan: %w", err)
		}
		plans = append(plans, plan)
	}

	return plans, rows.Err()
}

func (r *treatmentPlanRepository) GetByStatus(ctx context.Context, status string) ([]entities.TreatmentPlan, error) {
	query := `
		SELECT plan_id, patient_id, doctor_id, diagnosis, treatment_type, start_date, end_date, session_duration, status, goals, progress_notes
		FROM treatment_plan
		WHERE status = ?
	`

	rows, err := r.db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, fmt.Errorf("failed to get treatment plans: %w", err)
	}
	defer rows.Close()

	var plans []entities.TreatmentPlan
	for rows.Next() {
		var plan entities.TreatmentPlan
		if err := rows.Scan(
			&plan.PlanID,
			&plan.PatientID,
			&plan.DoctorID,
			&plan.Diagnosis,
			&plan.TreatmentType,
			&plan.StartDate,
			&plan.EndDate,
			&plan.SessionDuration,
			&plan.Status,
			&plan.Goals,
			&plan.ProgressNotes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan treatment plan: %w", err)
		}
		plans = append(plans, plan)
	}

	return plans, rows.Err()
}

func (r *treatmentPlanRepository) GetActivePlans(ctx context.Context) ([]entities.TreatmentPlan, error) {
	query := `
		SELECT plan_id, patient_id, doctor_id, diagnosis, treatment_type, start_date, end_date, session_duration, status, goals, progress_notes
		FROM treatment_plan
		WHERE status = 'active'
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get treatment plans: %w", err)
	}
	defer rows.Close()

	var plans []entities.TreatmentPlan
	for rows.Next() {
		var plan entities.TreatmentPlan
		if err := rows.Scan(
			&plan.PlanID,
			&plan.PatientID,
			&plan.DoctorID,
			&plan.Diagnosis,
			&plan.TreatmentType,
			&plan.StartDate,
			&plan.EndDate,
			&plan.SessionDuration,
			&plan.Status,
			&plan.Goals,
			&plan.ProgressNotes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan treatment plan: %w", err)
		}
		plans = append(plans, plan)
	}

	return plans, rows.Err()
}
