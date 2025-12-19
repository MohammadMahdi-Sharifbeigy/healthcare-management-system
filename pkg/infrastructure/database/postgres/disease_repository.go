package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type diseaseRepository struct {
	db *sql.DB
}

func NewDiseaseRepository(db *sql.DB) repositories.DiseaseRepository {
	return &diseaseRepository{db: db}
}

func (r *diseaseRepository) Create(ctx context.Context, disease *entities.Disease) error {
	if err := disease.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO disease (disease_name, category, icd_code, description)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		disease.DiseaseName,
		disease.Category,
		disease.ICDCode,
		disease.Description,
	)

	if err != nil {
		return fmt.Errorf("failed to create disease: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get disease id: %w", err)
	}

	disease.DiseaseID = int(id)
	return nil
}

func (r *diseaseRepository) GetByID(ctx context.Context, id int) (*entities.Disease, error) {
	query := `
		SELECT disease_id, disease_name, category, icd_code, description
		FROM disease
		WHERE disease_id = ?
	`

	disease := &entities.Disease{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&disease.DiseaseID,
		&disease.DiseaseName,
		&disease.Category,
		&disease.ICDCode,
		&disease.Description,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("disease not found")
		}
		return nil, fmt.Errorf("failed to get disease: %w", err)
	}

	return disease, nil
}

func (r *diseaseRepository) Update(ctx context.Context, disease *entities.Disease) error {
	if err := disease.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE disease
		SET disease_name = ?, category = ?, icd_code = ?, description = ?
		WHERE disease_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		disease.DiseaseName,
		disease.Category,
		disease.ICDCode,
		disease.Description,
		disease.DiseaseID,
	)

	if err != nil {
		return fmt.Errorf("failed to update disease: %w", err)
	}

	return nil
}

func (r *diseaseRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM disease WHERE disease_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete disease: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("disease not found")
	}

	return nil
}

func (r *diseaseRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.Disease, error) {
	query := `
		SELECT disease_id, disease_name, category, icd_code, description
		FROM disease
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get diseases: %w", err)
	}
	defer rows.Close()

	var diseases []entities.Disease
	for rows.Next() {
		var disease entities.Disease
		if err := rows.Scan(
			&disease.DiseaseID,
			&disease.DiseaseName,
			&disease.Category,
			&disease.ICDCode,
			&disease.Description,
		); err != nil {
			return nil, fmt.Errorf("failed to scan disease: %w", err)
		}
		diseases = append(diseases, disease)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating diseases: %w", err)
	}

	return diseases, nil
}

func (r *diseaseRepository) GetByName(ctx context.Context, name string) (*entities.Disease, error) {
	query := `
		SELECT disease_id, disease_name, category, icd_code, description
		FROM disease
		WHERE disease_name = ?
	`

	disease := &entities.Disease{}
	err := r.db.QueryRowContext(ctx, query, name).Scan(
		&disease.DiseaseID,
		&disease.DiseaseName,
		&disease.Category,
		&disease.ICDCode,
		&disease.Description,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("disease not found")
		}
		return nil, fmt.Errorf("failed to get disease by name: %w", err)
	}

	return disease, nil
}

func (r *diseaseRepository) GetByICDCode(ctx context.Context, icdCode string) (*entities.Disease, error) {
	query := `
		SELECT disease_id, disease_name, category, icd_code, description
		FROM disease
		WHERE icd_code = ?
	`

	disease := &entities.Disease{}
	err := r.db.QueryRowContext(ctx, query, icdCode).Scan(
		&disease.DiseaseID,
		&disease.DiseaseName,
		&disease.Category,
		&disease.ICDCode,
		&disease.Description,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("disease not found")
		}
		return nil, fmt.Errorf("failed to get disease by ICD code: %w", err)
	}

	return disease, nil
}

func (r *diseaseRepository) GetByCategory(ctx context.Context, category string) ([]entities.Disease, error) {
	query := `
		SELECT disease_id, disease_name, category, icd_code, description
		FROM disease
		WHERE category = ?
	`

	rows, err := r.db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, fmt.Errorf("failed to get diseases by category: %w", err)
	}
	defer rows.Close()

	var diseases []entities.Disease
	for rows.Next() {
		var disease entities.Disease
		if err := rows.Scan(
			&disease.DiseaseID,
			&disease.DiseaseName,
			&disease.Category,
			&disease.ICDCode,
			&disease.Description,
		); err != nil {
			return nil, fmt.Errorf("failed to scan disease: %w", err)
		}
		diseases = append(diseases, disease)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating diseases: %w", err)
	}

	return diseases, nil
}

func (r *diseaseRepository) SearchByName(ctx context.Context, name string) ([]entities.Disease, error) {
	query := `
		SELECT disease_id, disease_name, category, icd_code, description
		FROM disease
		WHERE disease_name LIKE ?
	`

	searchTerm := "%" + name + "%"
	rows, err := r.db.QueryContext(ctx, query, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search diseases: %w", err)
	}
	defer rows.Close()

	var diseases []entities.Disease
	for rows.Next() {
		var disease entities.Disease
		if err := rows.Scan(
			&disease.DiseaseID,
			&disease.DiseaseName,
			&disease.Category,
			&disease.ICDCode,
			&disease.Description,
		); err != nil {
			return nil, fmt.Errorf("failed to scan disease: %w", err)
		}
		diseases = append(diseases, disease)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating diseases: %w", err)
	}

	return diseases, nil
}
