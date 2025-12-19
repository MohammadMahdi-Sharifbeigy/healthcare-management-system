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

type appointmentRepository struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) repositories.AppointmentRepository {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) Create(ctx context.Context, appointment *entities.Appointment) error {
	if err := appointment.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO appointment (patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate, duration_minutes, notes)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		appointment.PatientID,
		appointment.DoctorID,
		appointment.AppointmentDate,
		appointment.Reason,
		appointment.Status,
		appointment.BloodPressure,
		appointment.HeartRate,
		appointment.DurationMinutes,
		appointment.Notes,
	)

	if err != nil {
		return fmt.Errorf("failed to create appointment: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get appointment id: %w", err)
	}

	appointment.AppointmentID = int(id)
	return nil
}

func (r *appointmentRepository) GetByID(ctx context.Context, id int) (*entities.Appointment, error) {
	query := `
		SELECT appointment_id, patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate, duration_minutes, notes
		FROM appointment
		WHERE appointment_id = ?
	`

	appointment := &entities.Appointment{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&appointment.AppointmentID,
		&appointment.PatientID,
		&appointment.DoctorID,
		&appointment.AppointmentDate,
		&appointment.Reason,
		&appointment.Status,
		&appointment.BloodPressure,
		&appointment.HeartRate,
		&appointment.DurationMinutes,
		&appointment.Notes,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("appointment not found")
		}
		return nil, fmt.Errorf("failed to get appointment: %w", err)
	}

	return appointment, nil
}

func (r *appointmentRepository) Update(ctx context.Context, appointment *entities.Appointment) error {
	if err := appointment.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE appointment
		SET patient_id = ?, doctor_id = ?, appointment_date = ?, reason = ?, status = ?, blood_pressure = ?, heart_rate = ?, duration_minutes = ?, notes = ?
		WHERE appointment_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		appointment.PatientID,
		appointment.DoctorID,
		appointment.AppointmentDate,
		appointment.Reason,
		appointment.Status,
		appointment.BloodPressure,
		appointment.HeartRate,
		appointment.DurationMinutes,
		appointment.Notes,
		appointment.AppointmentID,
	)

	if err != nil {
		return fmt.Errorf("failed to update appointment: %w", err)
	}

	return nil
}

func (r *appointmentRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM appointment WHERE appointment_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete appointment: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("appointment not found")
	}

	return nil
}

func (r *appointmentRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Appointment, error) {
	query := `
		SELECT appointment_id, patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate, duration_minutes, notes
		FROM appointment
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get appointments: %w", err)
	}
	defer rows.Close()

	var appointments []entities.Appointment
	for rows.Next() {
		var appointment entities.Appointment
		if err := rows.Scan(
			&appointment.AppointmentID,
			&appointment.PatientID,
			&appointment.DoctorID,
			&appointment.AppointmentDate,
			&appointment.Reason,
			&appointment.Status,
			&appointment.BloodPressure,
			&appointment.HeartRate,
			&appointment.DurationMinutes,
			&appointment.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan appointment: %w", err)
		}
		appointments = append(appointments, appointment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating appointments: %w", err)
	}

	return appointments, nil
}

func (r *appointmentRepository) GetByPatientID(ctx context.Context, patientID int) ([]entities.Appointment, error) {
	query := `
		SELECT appointment_id, patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate, duration_minutes, notes
		FROM appointment
		WHERE patient_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get appointments: %w", err)
	}
	defer rows.Close()

	var appointments []entities.Appointment
	for rows.Next() {
		var appointment entities.Appointment
		if err := rows.Scan(
			&appointment.AppointmentID,
			&appointment.PatientID,
			&appointment.DoctorID,
			&appointment.AppointmentDate,
			&appointment.Reason,
			&appointment.Status,
			&appointment.BloodPressure,
			&appointment.HeartRate,
			&appointment.DurationMinutes,
			&appointment.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan appointment: %w", err)
		}
		appointments = append(appointments, appointment)
	}

	return appointments, rows.Err()
}

func (r *appointmentRepository) GetByDoctorID(ctx context.Context, doctorID int) ([]entities.Appointment, error) {
	query := `
		SELECT appointment_id, patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate, duration_minutes, notes
		FROM appointment
		WHERE doctor_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, doctorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get appointments: %w", err)
	}
	defer rows.Close()

	var appointments []entities.Appointment
	for rows.Next() {
		var appointment entities.Appointment
		if err := rows.Scan(
			&appointment.AppointmentID,
			&appointment.PatientID,
			&appointment.DoctorID,
			&appointment.AppointmentDate,
			&appointment.Reason,
			&appointment.Status,
			&appointment.BloodPressure,
			&appointment.HeartRate,
			&appointment.DurationMinutes,
			&appointment.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan appointment: %w", err)
		}
		appointments = append(appointments, appointment)
	}

	return appointments, rows.Err()
}

func (r *appointmentRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]entities.Appointment, error) {
	query := `
		SELECT appointment_id, patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate, duration_minutes, notes
		FROM appointment
		WHERE appointment_date BETWEEN ? AND ?
	`

	rows, err := r.db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get appointments: %w", err)
	}
	defer rows.Close()

	var appointments []entities.Appointment
	for rows.Next() {
		var appointment entities.Appointment
		if err := rows.Scan(
			&appointment.AppointmentID,
			&appointment.PatientID,
			&appointment.DoctorID,
			&appointment.AppointmentDate,
			&appointment.Reason,
			&appointment.Status,
			&appointment.BloodPressure,
			&appointment.HeartRate,
			&appointment.DurationMinutes,
			&appointment.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan appointment: %w", err)
		}
		appointments = append(appointments, appointment)
	}

	return appointments, rows.Err()
}

func (r *appointmentRepository) GetByStatus(ctx context.Context, status string) ([]entities.Appointment, error) {
	query := `
		SELECT appointment_id, patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate, duration_minutes, notes
		FROM appointment
		WHERE status = ?
	`

	rows, err := r.db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, fmt.Errorf("failed to get appointments: %w", err)
	}
	defer rows.Close()

	var appointments []entities.Appointment
	for rows.Next() {
		var appointment entities.Appointment
		if err := rows.Scan(
			&appointment.AppointmentID,
			&appointment.PatientID,
			&appointment.DoctorID,
			&appointment.AppointmentDate,
			&appointment.Reason,
			&appointment.Status,
			&appointment.BloodPressure,
			&appointment.HeartRate,
			&appointment.DurationMinutes,
			&appointment.Notes,
		); err != nil {
			return nil, fmt.Errorf("failed to scan appointment: %w", err)
		}
		appointments = append(appointments, appointment)
	}

	return appointments, rows.Err()
}

func (r *appointmentRepository) CheckAvailability(ctx context.Context, doctorID int, appointmentDate time.Time) (bool, error) {
	query := `
		SELECT COUNT(*) FROM appointment
		WHERE doctor_id = ? AND appointment_date = ? AND status != 'cancelled'
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, doctorID, appointmentDate).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check availability: %w", err)
	}

	return count == 0, nil
}
