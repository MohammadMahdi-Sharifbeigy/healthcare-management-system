package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type diagnosisRepository struct {
	db *sql.DB
}

func NewDiagnosisRepository(db *sql.DB) repositories.DiagnosisRepository {
	return &diagnosisRepository{db: db}
}

func (r *diagnosisRepository) Create(ctx context.Context, diagnosis *entities.Diagnosis) error {
	if err := diagnosis.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO diagnosis (patient_id, doctor_id, disease_id, diagnosis_date, severity, notes)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		diagnosis.PatientID,
		diagnosis.DoctorID,
		diagnosis.DiseaseID,
		diagnosis.DiagnosisDate,
		diagnosis.Severity,
		diagnosis.Notes,
	)

	if err != nil {
		return fmt.Errorf("failed to create diagnosis: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get diagnosis id: %w", err)
	}

	diagnosis.DiagnosisID = int(id)
	return nil
}

func (r *diagnosisRepository) GetByID(ctx context.Context, id int) (*entities.Diagnosis, error) {
	query := `
		SELECT diagnosis_id, patient_id, doctor_id, disease_id, diagnosis_date, severity, notes
		FROM diagnosis
		WHERE diagnosis_id = ?
	`

	diagnosis := &entities.Diagnosis{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&diagnosis.DiagnosisID,
		&diagnosis.PatientID,
		&diagnosis.DoctorID,
		&diagnosis.DiseaseID,
		&diagnosis.DiagnosisDate,
		&diagnosis.Severity,
		&diagnosis.Notes,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("diagnosis not found")
		}
		return nil, fmt.Errorf("failed to get diagnosis: %w", err)
	}

	return diagnosis, nil
}

func (r *diagnosisRepository) Update(ctx context.Context, diagnosis *entities.Diagnosis) error {
	if err := diagnosis.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE diagnosis
		SET patient_id = ?, doctor_id = ?, disease_id = ?, diagnosis_date = ?, severity = ?, notes = ?
		WHERE diagnosis_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		diagnosis.PatientID,
		diagnosis.DoctorID,
		diagnosis.DiseaseID,
		diagnosis.DiagnosisDate,
		diagnosis.Severity,
		diagnosis.Notes,
		diagnosis.DiagnosisID,
	)

	if err != nil {
		return fmt.Errorf("failed to update diagnosis: %w", err)
	}

	return nil
}

func (r *diagnosisRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM diagnosis WHERE diagnosis_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete diagnosis: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("diagnosis not found")
	}

	return nil
}

func (r *diagnosisRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Diagnosis, error) {
	query := `
		SELECT diagnosis_id, patient_id, doctor_id, disease_id, diagnosis_date, severity, notes
		FROM diagnosis
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get diagnoses: %w", err)
	}
	defer rows.Close()

	var diagnoses []entities.Diagnosis
	for rows.Next() {
		var diagnosis entities.Diagnosis
		if err := rows.Scan(
			&diagnosis.DiagnosisID,
			&diagnosis.PatientID,
			&diagnosis.DoctorID,
			&diagnosis.DiseaseID,
			&diagnosis.DiagnosisDate,
			&diagnosis.Severity,
			&diagnosis.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan diagnosis: %w", err)
		}
		diagnoses = append(diagnoses, diagnosis)
	}

	return diagnoses, rows.Err()
}

func (r *diagnosisRepository) GetByPatientID(ctx context.Context, patientID int) ([]entities.Diagnosis, error) {
	query := `
		SELECT diagnosis_id, patient_id, doctor_id, disease_id, diagnosis_date, severity, notes
		FROM diagnosis
		WHERE patient_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get diagnoses: %w", err)
	}
	defer rows.Close()

	var diagnoses []entities.Diagnosis
	for rows.Next() {
		var diagnosis entities.Diagnosis
		if err := rows.Scan(
			&diagnosis.DiagnosisID,
			&diagnosis.PatientID,
			&diagnosis.DoctorID,
			&diagnosis.DiseaseID,
			&diagnosis.DiagnosisDate,
			&diagnosis.Severity,
			&diagnosis.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan diagnosis: %w", err)
		}
		diagnoses = append(diagnoses, diagnosis)
	}

	return diagnoses, rows.Err()
}

func (r *diagnosisRepository) GetByDoctorID(ctx context.Context, doctorID int) ([]entities.Diagnosis, error) {
	query := `
		SELECT diagnosis_id, patient_id, doctor_id, disease_id, diagnosis_date, severity, notes
		FROM diagnosis
		WHERE doctor_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, doctorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get diagnoses: %w", err)
	}
	defer rows.Close()

	var diagnoses []entities.Diagnosis
	for rows.Next() {
		var diagnosis entities.Diagnosis
		if err := rows.Scan(
			&diagnosis.DiagnosisID,
			&diagnosis.PatientID,
			&diagnosis.DoctorID,
			&diagnosis.DiseaseID,
			&diagnosis.DiagnosisDate,
			&diagnosis.Severity,
			&diagnosis.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan diagnosis: %w", err)
		}
		diagnoses = append(diagnoses, diagnosis)
	}

	return diagnoses, rows.Err()
}

func (r *diagnosisRepository) GetByDiseaseID(ctx context.Context, diseaseID int) ([]entities.Diagnosis, error) {
	query := `
		SELECT diagnosis_id, patient_id, doctor_id, disease_id, diagnosis_date, severity, notes
		FROM diagnosis
		WHERE disease_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, diseaseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get diagnoses: %w", err)
	}
	defer rows.Close()

	var diagnoses []entities.Diagnosis
	for rows.Next() {
		var diagnosis entities.Diagnosis
		if err := rows.Scan(
			&diagnosis.DiagnosisID,
			&diagnosis.PatientID,
			&diagnosis.DoctorID,
			&diagnosis.DiseaseID,
			&diagnosis.DiagnosisDate,
			&diagnosis.Severity,
			&diagnosis.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan diagnosis: %w", err)
		}
		diagnoses = append(diagnoses, diagnosis)
	}

	return diagnoses, rows.Err()
}

func (r *diagnosisRepository) GetBySeverity(ctx context.Context, severity string) ([]entities.Diagnosis, error) {
	query := `
		SELECT diagnosis_id, patient_id, doctor_id, disease_id, diagnosis_date, severity, notes
		FROM diagnosis
		WHERE severity = ?
	`

	rows, err := r.db.QueryContext(ctx, query, severity)
	if err != nil {
		return nil, fmt.Errorf("failed to get diagnoses: %w", err)
	}
	defer rows.Close()

	var diagnoses []entities.Diagnosis
	for rows.Next() {
		var diagnosis entities.Diagnosis
		if err := rows.Scan(
			&diagnosis.DiagnosisID,
			&diagnosis.PatientID,
			&diagnosis.DoctorID,
			&diagnosis.DiseaseID,
			&diagnosis.DiagnosisDate,
			&diagnosis.Severity,
			&diagnosis.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan diagnosis: %w", err)
		}
		diagnoses = append(diagnoses, diagnosis)
	}

	return diagnoses, rows.Err()
}

func (r *diagnosisRepository) GetSevereForPatient(ctx context.Context, patientID int) ([]entities.Diagnosis, error) {
	query := `
		SELECT diagnosis_id, patient_id, doctor_id, disease_id, diagnosis_date, severity, notes
		FROM diagnosis
		WHERE patient_id = ? AND severity = 'severe'
	`

	rows, err := r.db.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get severe diagnoses: %w", err)
	}
	defer rows.Close()

	var diagnoses []entities.Diagnosis
	for rows.Next() {
		var diagnosis entities.Diagnosis
		if err := rows.Scan(
			&diagnosis.DiagnosisID,
			&diagnosis.PatientID,
			&diagnosis.DoctorID,
			&diagnosis.DiseaseID,
			&diagnosis.DiagnosisDate,
			&diagnosis.Severity,
			&diagnosis.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan diagnosis: %w", err)
		}
		diagnoses = append(diagnoses, diagnosis)
	}

	return diagnoses, rows.Err()
}
