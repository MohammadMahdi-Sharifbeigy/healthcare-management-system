package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type medicationRepository struct {
	db *sql.DB
}

func NewMedicationRepository(db *sql.DB) repositories.MedicationRepository {
	return &medicationRepository{db: db}
}

func (r *medicationRepository) Create(ctx context.Context, medication *entities.Medication) error {
	if err := medication.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO medication (medication_name, generic_name, form, strength, manufacturer, side_effects)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		medication.MedicationName,
		medication.GenericName,
		medication.Form,
		medication.Strength,
		medication.Manufacturer,
		medication.SideEffects,
	)

	if err != nil {
		return fmt.Errorf("failed to create medication: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get medication id: %w", err)
	}

	medication.MedicationID = int(id)
	return nil
}

func (r *medicationRepository) GetByID(ctx context.Context, id int) (*entities.Medication, error) {
	query := `
		SELECT medication_id, medication_name, generic_name, form, strength, manufacturer, side_effects
		FROM medication
		WHERE medication_id = ?
	`

	medication := &entities.Medication{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&medication.MedicationID,
		&medication.MedicationName,
		&medication.GenericName,
		&medication.Form,
		&medication.Strength,
		&medication.Manufacturer,
		&medication.SideEffects,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("medication not found")
		}
		return nil, fmt.Errorf("failed to get medication: %w", err)
	}

	return medication, nil
}

func (r *medicationRepository) Update(ctx context.Context, medication *entities.Medication) error {
	if err := medication.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE medication
		SET medication_name = ?, generic_name = ?, form = ?, strength = ?, manufacturer = ?, side_effects = ?
		WHERE medication_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		medication.MedicationName,
		medication.GenericName,
		medication.Form,
		medication.Strength,
		medication.Manufacturer,
		medication.SideEffects,
		medication.MedicationID,
	)

	if err != nil {
		return fmt.Errorf("failed to update medication: %w", err)
	}

	return nil
}

func (r *medicationRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM medication WHERE medication_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete medication: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("medication not found")
	}

	return nil
}

func (r *medicationRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Medication, error) {
	query := `
		SELECT medication_id, medication_name, generic_name, form, strength, manufacturer, side_effects
		FROM medication
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get medications: %w", err)
	}
	defer rows.Close()

	var medications []entities.Medication
	for rows.Next() {
		var medication entities.Medication
		if err := rows.Scan(
			&medication.MedicationID,
			&medication.MedicationName,
			&medication.GenericName,
			&medication.Form,
			&medication.Strength,
			&medication.Manufacturer,
			&medication.SideEffects,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medication: %w", err)
		}
		medications = append(medications, medication)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating medications: %w", err)
	}

	return medications, nil
}

func (r *medicationRepository) GetByName(ctx context.Context, name string) (*entities.Medication, error) {
	query := `
		SELECT medication_id, medication_name, generic_name, form, strength, manufacturer, side_effects
		FROM medication
		WHERE medication_name = ?
	`

	medication := &entities.Medication{}
	err := r.db.QueryRowContext(ctx, query, name).Scan(
		&medication.MedicationID,
		&medication.MedicationName,
		&medication.GenericName,
		&medication.Form,
		&medication.Strength,
		&medication.Manufacturer,
		&medication.SideEffects,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("medication not found")
		}
		return nil, fmt.Errorf("failed to get medication by name: %w", err)
	}

	return medication, nil
}

func (r *medicationRepository) GetByGenericName(ctx context.Context, genericName string) (*entities.Medication, error) {
	query := `
		SELECT medication_id, medication_name, generic_name, form, strength, manufacturer, side_effects
		FROM medication
		WHERE generic_name = ?
	`

	medication := &entities.Medication{}
	err := r.db.QueryRowContext(ctx, query, genericName).Scan(
		&medication.MedicationID,
		&medication.MedicationName,
		&medication.GenericName,
		&medication.Form,
		&medication.Strength,
		&medication.Manufacturer,
		&medication.SideEffects,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("medication not found")
		}
		return nil, fmt.Errorf("failed to get medication by generic name: %w", err)
	}

	return medication, nil
}

func (r *medicationRepository) SearchByName(ctx context.Context, name string) ([]entities.Medication, error) {
	query := `
		SELECT medication_id, medication_name, generic_name, form, strength, manufacturer, side_effects
		FROM medication
		WHERE medication_name LIKE ? OR generic_name LIKE ?
	`

	searchTerm := "%" + name + "%"
	rows, err := r.db.QueryContext(ctx, query, searchTerm, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search medications: %w", err)
	}
	defer rows.Close()

	var medications []entities.Medication
	for rows.Next() {
		var medication entities.Medication
		if err := rows.Scan(
			&medication.MedicationID,
			&medication.MedicationName,
			&medication.GenericName,
			&medication.Form,
			&medication.Strength,
			&medication.Manufacturer,
			&medication.SideEffects,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medication: %w", err)
		}
		medications = append(medications, medication)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating medications: %w", err)
	}

	return medications, nil
}

func (r *medicationRepository) GetByForm(ctx context.Context, form string) ([]entities.Medication, error) {
	query := `
		SELECT medication_id, medication_name, generic_name, form, strength, manufacturer, side_effects
		FROM medication
		WHERE form = ?
	`

	rows, err := r.db.QueryContext(ctx, query, form)
	if err != nil {
		return nil, fmt.Errorf("failed to get medications by form: %w", err)
	}
	defer rows.Close()

	var medications []entities.Medication
	for rows.Next() {
		var medication entities.Medication
		if err := rows.Scan(
			&medication.MedicationID,
			&medication.MedicationName,
			&medication.GenericName,
			&medication.Form,
			&medication.Strength,
			&medication.Manufacturer,
			&medication.SideEffects,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medication: %w", err)
		}
		medications = append(medications, medication)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating medications: %w", err)
	}

	return medications, nil
}
