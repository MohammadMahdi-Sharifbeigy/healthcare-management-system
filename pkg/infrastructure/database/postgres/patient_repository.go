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

type patientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) repositories.PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) Create(ctx context.Context, patient *entities.Patient) error {
	if err := patient.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO patient (first_name, last_name, date_of_birth, gender, email, phone, address, emergency_contact, blood_type, medical_history)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		patient.FirstName,
		patient.LastName,
		patient.DateOfBirth,
		patient.Gender,
		patient.Email,
		patient.Phone,
		patient.Address,
		patient.EmergencyContact,
		patient.BloodType,
		patient.MedicalHistory,
	)

	if err != nil {
		return fmt.Errorf("failed to create patient: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get patient id: %w", err)
	}

	patient.PatientID = int(id)
	return nil
}

func (r *patientRepository) GetByID(ctx context.Context, id int) (*entities.Patient, error) {
	query := `
		SELECT patient_id, first_name, last_name, date_of_birth, gender, email, phone, address, emergency_contact, blood_type, medical_history
		FROM patient
		WHERE patient_id = ?
	`

	patient := &entities.Patient{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&patient.PatientID,
		&patient.FirstName,
		&patient.LastName,
		&patient.DateOfBirth,
		&patient.Gender,
		&patient.Email,
		&patient.Phone,
		&patient.Address,
		&patient.EmergencyContact,
		&patient.BloodType,
		&patient.MedicalHistory,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("patient not found")
		}
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}

	return patient, nil
}

func (r *patientRepository) Update(ctx context.Context, patient *entities.Patient) error {
	if err := patient.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE patient
		SET first_name = ?, last_name = ?, date_of_birth = ?, gender = ?, email = ?, phone = ?, address = ?, emergency_contact = ?, blood_type = ?, medical_history = ?
		WHERE patient_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		patient.FirstName,
		patient.LastName,
		patient.DateOfBirth,
		patient.Gender,
		patient.Email,
		patient.Phone,
		patient.Address,
		patient.EmergencyContact,
		patient.BloodType,
		patient.MedicalHistory,
		patient.PatientID,
	)

	if err != nil {
		return fmt.Errorf("failed to update patient: %w", err)
	}

	return nil
}

func (r *patientRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM patient WHERE patient_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete patient: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("patient not found")
	}

	return nil
}

func (r *patientRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Patient, error) {
	query := `
		SELECT patient_id, first_name, last_name, date_of_birth, gender, email, phone, address, emergency_contact, blood_type, medical_history
		FROM patient
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get patients: %w", err)
	}
	defer rows.Close()

	var patients []entities.Patient
	for rows.Next() {
		var patient entities.Patient
		if err := rows.Scan(
			&patient.PatientID,
			&patient.FirstName,
			&patient.LastName,
			&patient.DateOfBirth,
			&patient.Gender,
			&patient.Email,
			&patient.Phone,
			&patient.Address,
			&patient.EmergencyContact,
			&patient.BloodType,
			&patient.MedicalHistory,
		); err != nil {
			return nil, fmt.Errorf("failed to scan patient: %w", err)
		}
		patients = append(patients, patient)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating patients: %w", err)
	}

	return patients, nil
}

func (r *patientRepository) GetByEmail(ctx context.Context, email string) (*entities.Patient, error) {
	query := `
		SELECT patient_id, first_name, last_name, date_of_birth, gender, email, phone, address, emergency_contact, blood_type, medical_history
		FROM patient
		WHERE email = ?
	`

	patient := &entities.Patient{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&patient.PatientID,
		&patient.FirstName,
		&patient.LastName,
		&patient.DateOfBirth,
		&patient.Gender,
		&patient.Email,
		&patient.Phone,
		&patient.Address,
		&patient.EmergencyContact,
		&patient.BloodType,
		&patient.MedicalHistory,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("patient not found")
		}
		return nil, fmt.Errorf("failed to get patient by email: %w", err)
	}

	return patient, nil
}

func (r *patientRepository) GetByPhone(ctx context.Context, phone string) (*entities.Patient, error) {
	query := `
		SELECT patient_id, first_name, last_name, date_of_birth, gender, email, phone, address, emergency_contact, blood_type, medical_history
		FROM patient
		WHERE phone = ?
	`

	patient := &entities.Patient{}
	err := r.db.QueryRowContext(ctx, query, phone).Scan(
		&patient.PatientID,
		&patient.FirstName,
		&patient.LastName,
		&patient.DateOfBirth,
		&patient.Gender,
		&patient.Email,
		&patient.Phone,
		&patient.Address,
		&patient.EmergencyContact,
		&patient.BloodType,
		&patient.MedicalHistory,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("patient not found")
		}
		return nil, fmt.Errorf("failed to get patient by phone: %w", err)
	}

	return patient, nil
}

func (r *patientRepository) SearchByName(ctx context.Context, name string) ([]entities.Patient, error) {
	query := `
		SELECT patient_id, first_name, last_name, date_of_birth, gender, email, phone, address, emergency_contact, blood_type, medical_history
		FROM patient
		WHERE first_name LIKE ? OR last_name LIKE ?
	`

	searchTerm := "%" + name + "%"
	rows, err := r.db.QueryContext(ctx, query, searchTerm, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search patients: %w", err)
	}
	defer rows.Close()

	var patients []entities.Patient
	for rows.Next() {
		var patient entities.Patient
		if err := rows.Scan(
			&patient.PatientID,
			&patient.FirstName,
			&patient.LastName,
			&patient.DateOfBirth,
			&patient.Gender,
			&patient.Email,
			&patient.Phone,
			&patient.Address,
			&patient.EmergencyContact,
			&patient.BloodType,
			&patient.MedicalHistory,
		); err != nil {
			return nil, fmt.Errorf("failed to scan patient: %w", err)
		}
		patients = append(patients, patient)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating patients: %w", err)
	}

	return patients, nil
}
