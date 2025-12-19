package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type testCatalogRepository struct {
	db *sql.DB
}

func NewTestCatalogRepository(db *sql.DB) repositories.TestCatalogRepository {
	return &testCatalogRepository{db: db}
}

func (r *testCatalogRepository) Create(ctx context.Context, catalog *entities.TestCatalog) error {
	if err := catalog.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO test_catalog (test_name, test_category, description, normal_range)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		catalog.TestName,
		catalog.TestCategory,
		catalog.Description,
		catalog.NormalRange,
	)

	if err != nil {
		return fmt.Errorf("failed to create test catalog: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get test catalog id: %w", err)
	}

	catalog.TestCatalogID = int(id)
	return nil
}

func (r *testCatalogRepository) GetByID(ctx context.Context, id int) (*entities.TestCatalog, error) {
	query := `
		SELECT test_catalog_id, test_name, test_category, description, normal_range
		FROM test_catalog
		WHERE test_catalog_id = ?
	`

	catalog := &entities.TestCatalog{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&catalog.TestCatalogID,
		&catalog.TestName,
		&catalog.TestCategory,
		&catalog.Description,
		&catalog.NormalRange,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("test catalog not found")
		}
		return nil, fmt.Errorf("failed to get test catalog: %w", err)
	}

	return catalog, nil
}

func (r *testCatalogRepository) Update(ctx context.Context, catalog *entities.TestCatalog) error {
	if err := catalog.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE test_catalog
		SET test_name = ?, test_category = ?, description = ?, normal_range = ?
		WHERE test_catalog_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		catalog.TestName,
		catalog.TestCategory,
		catalog.Description,
		catalog.NormalRange,
		catalog.TestCatalogID,
	)

	if err != nil {
		return fmt.Errorf("failed to update test catalog: %w", err)
	}

	return nil
}

func (r *testCatalogRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM test_catalog WHERE test_catalog_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete test catalog: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("test catalog not found")
	}

	return nil
}

func (r *testCatalogRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.TestCatalog, error) {
	query := `
		SELECT test_catalog_id, test_name, test_category, description, normal_range
		FROM test_catalog
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get test catalogs: %w", err)
	}
	defer rows.Close()

	var catalogs []entities.TestCatalog
	for rows.Next() {
		var catalog entities.TestCatalog
		if err := rows.Scan(
			&catalog.TestCatalogID,
			&catalog.TestName,
			&catalog.TestCategory,
			&catalog.Description,
			&catalog.NormalRange,
		); err != nil {
			return nil, fmt.Errorf("failed to scan test catalog: %w", err)
		}
		catalogs = append(catalogs, catalog)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating test catalogs: %w", err)
	}

	return catalogs, nil
}

func (r *testCatalogRepository) GetByName(ctx context.Context, name string) (*entities.TestCatalog, error) {
	query := `
		SELECT test_catalog_id, test_name, test_category, description, normal_range
		FROM test_catalog
		WHERE test_name = ?
	`

	catalog := &entities.TestCatalog{}
	err := r.db.QueryRowContext(ctx, query, name).Scan(
		&catalog.TestCatalogID,
		&catalog.TestName,
		&catalog.TestCategory,
		&catalog.Description,
		&catalog.NormalRange,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("test catalog not found")
		}
		return nil, fmt.Errorf("failed to get test catalog by name: %w", err)
	}

	return catalog, nil
}

func (r *testCatalogRepository) GetByCategory(ctx context.Context, category string) ([]entities.TestCatalog, error) {
	query := `
		SELECT test_catalog_id, test_name, test_category, description, normal_range
		FROM test_catalog
		WHERE test_category = ?
	`

	rows, err := r.db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, fmt.Errorf("failed to get test catalogs by category: %w", err)
	}
	defer rows.Close()

	var catalogs []entities.TestCatalog
	for rows.Next() {
		var catalog entities.TestCatalog
		if err := rows.Scan(
			&catalog.TestCatalogID,
			&catalog.TestName,
			&catalog.TestCategory,
			&catalog.Description,
			&catalog.NormalRange,
		); err != nil {
			return nil, fmt.Errorf("failed to scan test catalog: %w", err)
		}
		catalogs = append(catalogs, catalog)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating test catalogs: %w", err)
	}

	return catalogs, nil
}

func (r *testCatalogRepository) SearchByName(ctx context.Context, name string) ([]entities.TestCatalog, error) {
	query := `
		SELECT test_catalog_id, test_name, test_category, description, normal_range
		FROM test_catalog
		WHERE test_name LIKE ?
	`

	searchTerm := "%" + name + "%"
	rows, err := r.db.QueryContext(ctx, query, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search test catalogs: %w", err)
	}
	defer rows.Close()

	var catalogs []entities.TestCatalog
	for rows.Next() {
		var catalog entities.TestCatalog
		if err := rows.Scan(
			&catalog.TestCatalogID,
			&catalog.TestName,
			&catalog.TestCategory,
			&catalog.Description,
			&catalog.NormalRange,
		); err != nil {
			return nil, fmt.Errorf("failed to scan test catalog: %w", err)
		}
		catalogs = append(catalogs, catalog)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating test catalogs: %w", err)
	}

	return catalogs, nil
}
