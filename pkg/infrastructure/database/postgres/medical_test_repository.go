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

type medicalTestRepository struct {
	db *sql.DB
}

func NewMedicalTestRepository(db *sql.DB) repositories.MedicalTestRepository {
	return &medicalTestRepository{db: db}
}

func (r *medicalTestRepository) Create(ctx context.Context, test *entities.MedicalTest) error {
	if err := test.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO medical_test (patient_id, test_catalog_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		test.PatientID,
		test.TestCatalogID,
		test.TestDate,
		test.ResultValue,
		test.NormalRange,
		test.Findings,
		test.ResultStatus,
		test.ImageURL,
		test.InterpretedBy,
	)

	if err != nil {
		return fmt.Errorf("failed to create medical test: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get medical test id: %w", err)
	}

	test.TestID = int(id)
	return nil
}

func (r *medicalTestRepository) GetByID(ctx context.Context, id int) (*entities.MedicalTest, error) {
	query := `
		SELECT test_id, patient_id, test_catalog_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by
		FROM medical_test
		WHERE test_id = ?
	`

	test := &entities.MedicalTest{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&test.TestID,
		&test.PatientID,
		&test.TestCatalogID,
		&test.TestDate,
		&test.ResultValue,
		&test.NormalRange,
		&test.Findings,
		&test.ResultStatus,
		&test.ImageURL,
		&test.InterpretedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound("medical test not found")
		}
		return nil, fmt.Errorf("failed to get medical test: %w", err)
	}

	return test, nil
}

func (r *medicalTestRepository) Update(ctx context.Context, test *entities.MedicalTest) error {
	if err := test.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE medical_test
		SET patient_id = ?, test_catalog_id = ?, test_date = ?, result_value = ?, normal_range = ?, findings = ?, result_status = ?, image_url = ?, interpreted_by = ?
		WHERE test_id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		test.PatientID,
		test.TestCatalogID,
		test.TestDate,
		test.ResultValue,
		test.NormalRange,
		test.Findings,
		test.ResultStatus,
		test.ImageURL,
		test.InterpretedBy,
		test.TestID,
	)

	if err != nil {
		return fmt.Errorf("failed to update medical test: %w", err)
	}

	return nil
}

func (r *medicalTestRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM medical_test WHERE test_id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete medical test: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return entities.ErrNotFound("medical test not found")
	}

	return nil
}

func (r *medicalTestRepository) GetAll(ctx context.Context, limit, offset int) ([]entities.MedicalTest, error) {
	query := `
		SELECT test_id, patient_id, test_catalog_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by
		FROM medical_test
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get medical tests: %w", err)
	}
	defer rows.Close()

	var tests []entities.MedicalTest
	for rows.Next() {
		var test entities.MedicalTest
		if err := rows.Scan(
			&test.TestID,
			&test.PatientID,
			&test.TestCatalogID,
			&test.TestDate,
			&test.ResultValue,
			&test.NormalRange,
			&test.Findings,
			&test.ResultStatus,
			&test.ImageURL,
			&test.InterpretedBy,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medical test: %w", err)
		}
		tests = append(tests, test)
	}

	return tests, rows.Err()
}

func (r *medicalTestRepository) GetByPatientID(ctx context.Context, patientID int) ([]entities.MedicalTest, error) {
	query := `
		SELECT test_id, patient_id, test_catalog_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by
		FROM medical_test
		WHERE patient_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get medical tests: %w", err)
	}
	defer rows.Close()

	var tests []entities.MedicalTest
	for rows.Next() {
		var test entities.MedicalTest
		if err := rows.Scan(
			&test.TestID,
			&test.PatientID,
			&test.TestCatalogID,
			&test.TestDate,
			&test.ResultValue,
			&test.NormalRange,
			&test.Findings,
			&test.ResultStatus,
			&test.ImageURL,
			&test.InterpretedBy,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medical test: %w", err)
		}
		tests = append(tests, test)
	}

	return tests, rows.Err()
}

func (r *medicalTestRepository) GetByTestCatalogID(ctx context.Context, testCatalogID int) ([]entities.MedicalTest, error) {
	query := `
		SELECT test_id, patient_id, test_catalog_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by
		FROM medical_test
		WHERE test_catalog_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, testCatalogID)
	if err != nil {
		return nil, fmt.Errorf("failed to get medical tests: %w", err)
	}
	defer rows.Close()

	var tests []entities.MedicalTest
	for rows.Next() {
		var test entities.MedicalTest
		if err := rows.Scan(
			&test.TestID,
			&test.PatientID,
			&test.TestCatalogID,
			&test.TestDate,
			&test.ResultValue,
			&test.NormalRange,
			&test.Findings,
			&test.ResultStatus,
			&test.ImageURL,
			&test.InterpretedBy,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medical test: %w", err)
		}
		tests = append(tests, test)
	}

	return tests, rows.Err()
}

func (r *medicalTestRepository) GetAbnormalResults(ctx context.Context) ([]entities.MedicalTest, error) {
	query := `
		SELECT test_id, patient_id, test_catalog_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by
		FROM medical_test
		WHERE result_status IN ('abnormal', 'critical')
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get abnormal results: %w", err)
	}
	defer rows.Close()

	var tests []entities.MedicalTest
	for rows.Next() {
		var test entities.MedicalTest
		if err := rows.Scan(
			&test.TestID,
			&test.PatientID,
			&test.TestCatalogID,
			&test.TestDate,
			&test.ResultValue,
			&test.NormalRange,
			&test.Findings,
			&test.ResultStatus,
			&test.ImageURL,
			&test.InterpretedBy,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medical test: %w", err)
		}
		tests = append(tests, test)
	}

	return tests, rows.Err()
}

func (r *medicalTestRepository) GetCriticalResults(ctx context.Context) ([]entities.MedicalTest, error) {
	query := `
		SELECT test_id, patient_id, test_catalog_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by
		FROM medical_test
		WHERE result_status = 'critical'
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get critical results: %w", err)
	}
	defer rows.Close()

	var tests []entities.MedicalTest
	for rows.Next() {
		var test entities.MedicalTest
		if err := rows.Scan(
			&test.TestID,
			&test.PatientID,
			&test.TestCatalogID,
			&test.TestDate,
			&test.ResultValue,
			&test.NormalRange,
			&test.Findings,
			&test.ResultStatus,
			&test.ImageURL,
			&test.InterpretedBy,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medical test: %w", err)
		}
		tests = append(tests, test)
	}

	return tests, rows.Err()
}

func (r *medicalTestRepository) GetByPatientAndTestType(ctx context.Context, patientID, testCatalogID int) ([]entities.MedicalTest, error) {
	query := `
		SELECT test_id, patient_id, test_catalog_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by
		FROM medical_test
		WHERE patient_id = ? AND test_catalog_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, patientID, testCatalogID)
	if err != nil {
		return nil, fmt.Errorf("failed to get medical tests: %w", err)
	}
	defer rows.Close()

	var tests []entities.MedicalTest
	for rows.Next() {
		var test entities.MedicalTest
		if err := rows.Scan(
			&test.TestID,
			&test.PatientID,
			&test.TestCatalogID,
			&test.TestDate,
			&test.ResultValue,
			&test.NormalRange,
			&test.Findings,
			&test.ResultStatus,
			&test.ImageURL,
			&test.InterpretedBy,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medical test: %w", err)
		}
		tests = append(tests, test)
	}

	return tests, rows.Err()
}

func (r *medicalTestRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]entities.MedicalTest, error) {
	query := `
		SELECT test_id, patient_id, test_catalog_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by
		FROM medical_test
		WHERE test_date BETWEEN ? AND ?
	`

	rows, err := r.db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get medical tests: %w", err)
	}
	defer rows.Close()

	var tests []entities.MedicalTest
	for rows.Next() {
		var test entities.MedicalTest
		if err := rows.Scan(
			&test.TestID,
			&test.PatientID,
			&test.TestCatalogID,
			&test.TestDate,
			&test.ResultValue,
			&test.NormalRange,
			&test.Findings,
			&test.ResultStatus,
			&test.ImageURL,
			&test.InterpretedBy,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medical test: %w", err)
		}
		tests = append(tests, test)
	}

	return tests, rows.Err()
}

func (r *medicalTestRepository) GetByResultStatus(ctx context.Context, status string) ([]entities.MedicalTest, error) {
	query := `
		SELECT test_id, patient_id, test_catalog_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by
		FROM medical_test
		WHERE result_status = ?
	`

	rows, err := r.db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, fmt.Errorf("failed to get medical tests: %w", err)
	}
	defer rows.Close()

	var tests []entities.MedicalTest
	for rows.Next() {
		var test entities.MedicalTest
		if err := rows.Scan(
			&test.TestID,
			&test.PatientID,
			&test.TestCatalogID,
			&test.TestDate,
			&test.ResultValue,
			&test.NormalRange,
			&test.Findings,
			&test.ResultStatus,
			&test.ImageURL,
			&test.InterpretedBy,
		); err != nil {
			return nil, fmt.Errorf("failed to scan medical test: %w", err)
		}
		tests = append(tests, test)
	}

	return tests, rows.Err()
}
