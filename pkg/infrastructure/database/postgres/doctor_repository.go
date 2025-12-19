package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type doctorRepository struct {
	db *sql.DB
}

func NewDoctorRepository(db *sql.DB) repositories.DoctorRepository {
	return &doctorRepository{db: db}
}

func (r *doctorRepository) Create(ctx context.Context, doctor *entities.Doctor) error {
	if err := doctor.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO doctor (doctor_id, first_name, last_name, email, phone, license_number, specialization, department)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query,
		doctor.DoctorID,
		doctor.FirstName,
		doctor.LastName,
		doctor.Email,
		doctor.Phone,
		doctor.LicenseNumber,
		doctor.Specialization,
		doctor.Department,
	)

	if err != nil {
		return fmt.Errorf("failed to create doctor: %w", err)
	}

	return nil
}

func (r *doctorRepository) GetByID(ctx context.Context, id int) (*entities.Doctor, error) {
	query := `
		SELECT doctor_id, first_name, last_name, email, phone, license_number, specialization, department
		FROM doctor
		WHERE doctor_id = ?
	`

	doctor := &entities.Doctor{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&doctor.DoctorID,
		&doctor.FirstName,
		&doctor.LastName,
		&doctor.Email,
		&doctor.Phone,
		&doctor.LicenseNumber,
		&doctor.Specialization,
		&doctor.Department,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("doctor not found")
		}
		return nil, fmt.Errorf("failed to get doctor: %w", err)
	}

	return doctor, nil
}

func (r *doctorRepository) Update(ctx context.Context, doctor *entities.Doctor) error {
	if err := doctor.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE doctor
		SET first_name = ?, last_name = ?, email = ?, phone = ?, license_number = ?, specialization = ?, department = ?
		WHERE doctor_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		doctor.FirstName,
		doctor.LastName,
		doctor.Email,
		doctor.Phone,
		doctor.LicenseNumber,
		doctor.Specialization,
		doctor.Department,
		doctor.DoctorID,
	)

	if err != nil {
		return fmt.Errorf("failed to update doctor: %w", err)
	}

	return nil
}

func (r *doctorRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM doctor WHERE doctor_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete doctor: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("doctor not found")
	}

	return nil
}

func (r *doctorRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Doctor, error) {
	query := `
		SELECT doctor_id, first_name, last_name, email, phone, license_number, specialization, department
		FROM doctor
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get doctors: %w", err)
	}
	defer rows.Close()

	var doctors []entities.Doctor
	for rows.Next() {
		var doctor entities.Doctor
		if err := rows.Scan(
			&doctor.DoctorID,
			&doctor.FirstName,
			&doctor.LastName,
			&doctor.Email,
			&doctor.Phone,
			&doctor.LicenseNumber,
			&doctor.Specialization,
			&doctor.Department,
		); err != nil {
			return nil, fmt.Errorf("failed to scan doctor: %w", err)
		}
		doctors = append(doctors, doctor)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating doctors: %w", err)
	}

	return doctors, nil
}

func (r *doctorRepository) GetByEmail(ctx context.Context, email string) (*entities.Doctor, error) {
	query := `
		SELECT doctor_id, first_name, last_name, email, phone, license_number, specialization, department
		FROM doctor
		WHERE email = ?
	`

	doctor := &entities.Doctor{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&doctor.DoctorID,
		&doctor.FirstName,
		&doctor.LastName,
		&doctor.Email,
		&doctor.Phone,
		&doctor.LicenseNumber,
		&doctor.Specialization,
		&doctor.Department,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("doctor not found")
		}
		return nil, fmt.Errorf("failed to get doctor by email: %w", err)
	}

	return doctor, nil
}

func (r *doctorRepository) GetByLicenseNumber(ctx context.Context, licenseNumber string) (*entities.Doctor, error) {
	query := `
		SELECT doctor_id, first_name, last_name, email, phone, license_number, specialization, department
		FROM doctor
		WHERE license_number = ?
	`

	doctor := &entities.Doctor{}
	err := r.db.QueryRowContext(ctx, query, licenseNumber).Scan(
		&doctor.DoctorID,
		&doctor.FirstName,
		&doctor.LastName,
		&doctor.Email,
		&doctor.Phone,
		&doctor.LicenseNumber,
		&doctor.Specialization,
		&doctor.Department,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("doctor not found")
		}
		return nil, fmt.Errorf("failed to get doctor by license: %w", err)
	}

	return doctor, nil
}

func (r *doctorRepository) GetBySpecialization(ctx context.Context, specialization string) ([]entities.Doctor, error) {
	query := `
		SELECT doctor_id, first_name, last_name, email, phone, license_number, specialization, department
		FROM doctor
		WHERE specialization = ?
	`

	rows, err := r.db.QueryContext(ctx, query, specialization)
	if err != nil {
		return nil, fmt.Errorf("failed to get doctors by specialization: %w", err)
	}
	defer rows.Close()

	var doctors []entities.Doctor
	for rows.Next() {
		var doctor entities.Doctor
		if err := rows.Scan(
			&doctor.DoctorID,
			&doctor.FirstName,
			&doctor.LastName,
			&doctor.Email,
			&doctor.Phone,
			&doctor.LicenseNumber,
			&doctor.Specialization,
			&doctor.Department,
		); err != nil {
			return nil, fmt.Errorf("failed to scan doctor: %w", err)
		}
		doctors = append(doctors, doctor)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating doctors: %w", err)
	}

	return doctors, nil
}

func (r *doctorRepository) GetByDepartment(ctx context.Context, department string) ([]entities.Doctor, error) {
	query := `
		SELECT doctor_id, first_name, last_name, email, phone, license_number, specialization, department
		FROM doctor
		WHERE department = ?
	`

	rows, err := r.db.QueryContext(ctx, query, department)
	if err != nil {
		return nil, fmt.Errorf("failed to get doctors by department: %w", err)
	}
	defer rows.Close()

	var doctors []entities.Doctor
	for rows.Next() {
		var doctor entities.Doctor
		if err := rows.Scan(
			&doctor.DoctorID,
			&doctor.FirstName,
			&doctor.LastName,
			&doctor.Email,
			&doctor.Phone,
			&doctor.LicenseNumber,
			&doctor.Specialization,
			&doctor.Department,
		); err != nil {
			return nil, fmt.Errorf("failed to scan doctor: %w", err)
		}
		doctors = append(doctors, doctor)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating doctors: %w", err)
	}

	return doctors, nil
}
