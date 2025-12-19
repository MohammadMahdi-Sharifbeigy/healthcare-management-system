package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type prescriptionRepository struct {
	db *sql.DB
}

func NewPrescriptionRepository(db *sql.DB) repositories.PrescriptionRepository {
	return &prescriptionRepository{db: db}
}

func (r *prescriptionRepository) Create(ctx context.Context, prescription *entities.Prescription) error {
	if err := prescription.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO prescription (patient_id, doctor_id, prescribed_date)
		VALUES (?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		prescription.PatientID,
		prescription.DoctorID,
		prescription.PrescribedDate,
	)

	if err != nil {
		return fmt.Errorf("failed to create prescription: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get prescription id: %w", err)
	}

	prescription.PrescriptionID = int(id)
	return nil
}

func (r *prescriptionRepository) GetByID(ctx context.Context, id int) (*entities.Prescription, error) {
	query := `
		SELECT prescription_id, patient_id, doctor_id, prescribed_date
		FROM prescription
		WHERE prescription_id = ?
	`

	prescription := &entities.Prescription{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&prescription.PrescriptionID,
		&prescription.PatientID,
		&prescription.DoctorID,
		&prescription.PrescribedDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("prescription not found")
		}
		return nil, fmt.Errorf("failed to get prescription: %w", err)
	}

	return prescription, nil
}

func (r *prescriptionRepository) Update(ctx context.Context, prescription *entities.Prescription) error {
	if err := prescription.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE prescription
		SET patient_id = ?, doctor_id = ?, prescribed_date = ?
		WHERE prescription_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		prescription.PatientID,
		prescription.DoctorID,
		prescription.PrescribedDate,
		prescription.PrescriptionID,
	)

	if err != nil {
		return fmt.Errorf("failed to update prescription: %w", err)
	}

	return nil
}

func (r *prescriptionRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM prescription WHERE prescription_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete prescription: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("prescription not found")
	}

	return nil
}

func (r *prescriptionRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Prescription, error) {
	query := `
		SELECT prescription_id, patient_id, doctor_id, prescribed_date
		FROM prescription
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get prescriptions: %w", err)
	}
	defer rows.Close()

	var prescriptions []entities.Prescription
	for rows.Next() {
		var prescription entities.Prescription
		if err := rows.Scan(
			&prescription.PrescriptionID,
			&prescription.PatientID,
			&prescription.DoctorID,
			&prescription.PrescribedDate,
		); err != nil {
			return nil, fmt.Errorf("failed to scan prescription: %w", err)
		}
		prescriptions = append(prescriptions, prescription)
	}

	return prescriptions, rows.Err()
}

func (r *prescriptionRepository) GetByPatientID(ctx context.Context, patientID int) ([]entities.Prescription, error) {
	query := `
		SELECT prescription_id, patient_id, doctor_id, prescribed_date
		FROM prescription
		WHERE patient_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get prescriptions: %w", err)
	}
	defer rows.Close()

	var prescriptions []entities.Prescription
	for rows.Next() {
		var prescription entities.Prescription
		if err := rows.Scan(
			&prescription.PrescriptionID,
			&prescription.PatientID,
			&prescription.DoctorID,
			&prescription.PrescribedDate,
		); err != nil {
			return nil, fmt.Errorf("failed to scan prescription: %w", err)
		}
		prescriptions = append(prescriptions, prescription)
	}

	return prescriptions, rows.Err()
}

func (r *prescriptionRepository) GetByDoctorID(ctx context.Context, doctorID int) ([]entities.Prescription, error) {
	query := `
		SELECT prescription_id, patient_id, doctor_id, prescribed_date
		FROM prescription
		WHERE doctor_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, doctorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get prescriptions: %w", err)
	}
	defer rows.Close()

	var prescriptions []entities.Prescription
	for rows.Next() {
		var prescription entities.Prescription
		if err := rows.Scan(
			&prescription.PrescriptionID,
			&prescription.PatientID,
			&prescription.DoctorID,
			&prescription.PrescribedDate,
		); err != nil {
			return nil, fmt.Errorf("failed to scan prescription: %w", err)
		}
		prescriptions = append(prescriptions, prescription)
	}

	return prescriptions, rows.Err()
}

func (r *prescriptionRepository) GetActivePrescriptionsByPatient(ctx context.Context, patientID int) ([]entities.Prescription, error) {
	query := `
		SELECT p.prescription_id, p.patient_id, p.doctor_id, p.prescribed_date
		FROM prescription p
		JOIN prescription_medication pm ON p.prescription_id = pm.prescription_id
		WHERE p.patient_id = ? AND pm.end_date >= ?
		GROUP BY p.prescription_id
	`

	rows, err := r.db.QueryContext(ctx, query, patientID, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to get prescriptions: %w", err)
	}
	defer rows.Close()

	var prescriptions []entities.Prescription
	for rows.Next() {
		var prescription entities.Prescription
		if err := rows.Scan(
			&prescription.PrescriptionID,
			&prescription.PatientID,
			&prescription.DoctorID,
			&prescription.PrescribedDate,
		); err != nil {
			return nil, fmt.Errorf("failed to scan prescription: %w", err)
		}
		prescriptions = append(prescriptions, prescription)
	}

	return prescriptions, rows.Err()
}

func (r *prescriptionRepository) GetExpiringPrescriptionsByPatient(ctx context.Context, patientID int, days int) ([]entities.Prescription, error) {
	query := `
		SELECT p.prescription_id, p.patient_id, p.doctor_id, p.prescribed_date
		FROM prescription p
		JOIN prescription_medication pm ON p.prescription_id = pm.prescription_id
		WHERE p.patient_id = ? AND pm.end_date <= ? + INTERVAL ? DAY AND pm.end_date >= ?
		GROUP BY p.prescription_id
	`

	rows, err := r.db.QueryContext(ctx, query, patientID, time.Now(), days, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to get prescriptions: %w", err)
	}
	defer rows.Close()

	var prescriptions []entities.Prescription
	for rows.Next() {
		var prescription entities.Prescription
		if err := rows.Scan(
			&prescription.PrescriptionID,
			&prescription.PatientID,
			&prescription.DoctorID,
			&prescription.PrescribedDate,
		); err != nil {
			return nil, fmt.Errorf("failed to scan prescription: %w", err)
		}
		prescriptions = append(prescriptions, prescription)
	}

	return prescriptions, rows.Err()
}
