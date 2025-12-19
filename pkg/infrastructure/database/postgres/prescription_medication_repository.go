package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type prescriptionMedicationRepository struct {
	db *sql.DB
}

func NewPrescriptionMedicationRepository(db *sql.DB) repositories.PrescriptionMedicationRepository {
	return &prescriptionMedicationRepository{db: db}
}

func (r *prescriptionMedicationRepository) Create(ctx context.Context, pm *entities.PrescriptionMedication) error {
	if err := pm.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO prescription_medication (prescription_id, medication_id, dosage, frequency, instructions, end_date)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		pm.PrescriptionID,
		pm.MedicationID,
		pm.Dosage,
		pm.Frequency,
		pm.Instructions,
		pm.EndDate,
	)

	if err != nil {
		return fmt.Errorf("failed to create prescription medication: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get prescription medication id: %w", err)
	}

	pm.PrescriptionMedicationID = int(id)
	return nil
}

func (r *prescriptionMedicationRepository) GetByID(ctx context.Context, id int) (*entities.PrescriptionMedication, error) {
	query := `
		SELECT prescription_medication_id, prescription_id, medication_id, dosage, frequency, instructions, end_date
		FROM prescription_medication
		WHERE prescription_medication_id = ?
	`

	pm := &entities.PrescriptionMedication{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&pm.PrescriptionMedicationID,
		&pm.PrescriptionID,
		&pm.MedicationID,
		&pm.Dosage,
		&pm.Frequency,
		&pm.Instructions,
		&pm.EndDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("prescription medication not found")
		}
		return nil, fmt.Errorf("failed to get prescription medication: %w", err)
	}

	return pm, nil
}

func (r *prescriptionMedicationRepository) Update(ctx context.Context, pm *entities.PrescriptionMedication) error {
	if err := pm.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE prescription_medication
		SET prescription_id = ?, medication_id = ?, dosage = ?, frequency = ?, instructions = ?, end_date = ?
		WHERE prescription_medication_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		pm.PrescriptionID,
		pm.MedicationID,
		pm.Dosage,
		pm.Frequency,
		pm.Instructions,
		pm.EndDate,
		pm.PrescriptionMedicationID,
	)

	if err != nil {
		return fmt.Errorf("failed to update prescription medication: %w", err)
	}

	return nil
}

func (r *prescriptionMedicationRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM prescription_medication WHERE prescription_medication_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete prescription medication: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("prescription medication not found")
	}

	return nil
}

func (r *prescriptionMedicationRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.PrescriptionMedication, error) {
	query := `
		SELECT prescription_medication_id, prescription_id, medication_id, dosage, frequency, instructions, end_date
		FROM prescription_medication
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get prescription medications: %w", err)
	}
	defer rows.Close()

	var pms []entities.PrescriptionMedication
	for rows.Next() {
		var pm entities.PrescriptionMedication
		if err := rows.Scan(
			&pm.PrescriptionMedicationID,
			&pm.PrescriptionID,
			&pm.MedicationID,
			&pm.Dosage,
			&pm.Frequency,
			&pm.Instructions,
			&pm.EndDate,
		); err != nil {
			return nil, fmt.Errorf("failed to scan prescription medication: %w", err)
		}
		pms = append(pms, pm)
	}

	return pms, rows.Err()
}

func (r *prescriptionMedicationRepository) GetByPrescriptionID(ctx context.Context, prescriptionID int) ([]entities.PrescriptionMedication, error) {
	query := `
		SELECT prescription_medication_id, prescription_id, medication_id, dosage, frequency, instructions, end_date
		FROM prescription_medication
		WHERE prescription_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, prescriptionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get prescription medications: %w", err)
	}
	defer rows.Close()

	var pms []entities.PrescriptionMedication
	for rows.Next() {
		var pm entities.PrescriptionMedication
		if err := rows.Scan(
			&pm.PrescriptionMedicationID,
			&pm.PrescriptionID,
			&pm.MedicationID,
			&pm.Dosage,
			&pm.Frequency,
			&pm.Instructions,
			&pm.EndDate,
		); err != nil {
			return nil, fmt.Errorf("failed to scan prescription medication: %w", err)
		}
		pms = append(pms, pm)
	}

	return pms, rows.Err()
}

func (r *prescriptionMedicationRepository) GetByMedicationID(ctx context.Context, medicationID int) ([]entities.PrescriptionMedication, error) {
	query := `
		SELECT prescription_medication_id, prescription_id, medication_id, dosage, frequency, instructions, end_date
		FROM prescription_medication
		WHERE medication_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, medicationID)
	if err != nil {
		return nil, fmt.Errorf("failed to get prescription medications: %w", err)
	}
	defer rows.Close()

	var pms []entities.PrescriptionMedication
	for rows.Next() {
		var pm entities.PrescriptionMedication
		if err := rows.Scan(
			&pm.PrescriptionMedicationID,
			&pm.PrescriptionID,
			&pm.MedicationID,
			&pm.Dosage,
			&pm.Frequency,
			&pm.Instructions,
			&pm.EndDate,
		); err != nil {
			return nil, fmt.Errorf("failed to scan prescription medication: %w", err)
		}
		pms = append(pms, pm)
	}

	return pms, rows.Err()
}

func (r *prescriptionMedicationRepository) DeleteByPrescriptionAndMedication(ctx context.Context, prescriptionID, medicationID int) error {
	query := `DELETE FROM prescription_medication WHERE prescription_id = ? AND medication_id = ?`

	_, err := r.db.ExecContext(ctx, query, prescriptionID, medicationID)
	if err != nil {
		return fmt.Errorf("failed to delete prescription medication: %w", err)
	}

	return nil
}
