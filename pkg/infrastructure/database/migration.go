package database

import (
	"database/sql"
	"fmt"
)

// RunMigrations executes all database migrations
func RunMigrations(db *sql.DB) error {
	// Create migrations table if it doesn't exist
	if err := createMigrationsTable(db); err != nil {
		return err
	}

	migrations := []Migration{
		{
			ID:      "001_create_disease_table",
			Version: 1,
			Up:      createDiseaseTable,
			Down:    dropDiseaseTable,
		},
		{
			ID:      "002_create_doctor_table",
			Version: 2,
			Up:      createDoctorTable,
			Down:    dropDoctorTable,
		},
		{
			ID:      "003_create_patient_table",
			Version: 3,
			Up:      createPatientTable,
			Down:    dropPatientTable,
		},
		{
			ID:      "004_create_medication_table",
			Version: 4,
			Up:      createMedicationTable,
			Down:    dropMedicationTable,
		},
		{
			ID:      "005_create_test_catalog_table",
			Version: 5,
			Up:      createTestCatalogTable,
			Down:    dropTestCatalogTable,
		},
		{
			ID:      "006_create_diagnosis_table",
			Version: 6,
			Up:      createDiagnosisTable,
			Down:    dropDiagnosisTable,
		},
		{
			ID:      "007_create_appointment_table",
			Version: 7,
			Up:      createAppointmentTable,
			Down:    dropAppointmentTable,
		},
		{
			ID:      "008_create_medical_test_table",
			Version: 8,
			Up:      createMedicalTestTable,
			Down:    dropMedicalTestTable,
		},
		{
			ID:      "009_create_prescription_table",
			Version: 9,
			Up:      createPrescriptionTable,
			Down:    dropPrescriptionTable,
		},
		{
			ID:      "010_create_treatment_plan_table",
			Version: 10,
			Up:      createTreatmentPlanTable,
			Down:    dropTreatmentPlanTable,
		},
		{
			ID:      "011_create_prescription_medication_table",
			Version: 11,
			Up:      createPrescriptionMedicationTable,
			Down:    dropPrescriptionMedicationTable,
		},
		{
			ID:      "012_create_indexes",
			Version: 12,
			Up:      createIndexes,
			Down:    dropIndexes,
		},
	}

	// Run each migration
	for _, migration := range migrations {
		if err := runMigration(db, migration); err != nil {
			return err
		}
	}

	fmt.Println("✅ All migrations completed successfully")
	return nil
}

type Migration struct {
	ID      string
	Version int
	Up      func(*sql.DB) error
	Down    func(*sql.DB) error
}

// createMigrationsTable creates schema_migrations table
func createMigrationsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			id SERIAL PRIMARY KEY,
			version INT UNIQUE NOT NULL,
			name VARCHAR(255) NOT NULL,
			executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(query)
	return err
}

// runMigration checks if migration was run and executes if needed
func runMigration(db *sql.DB, m Migration) error {
	var executed bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)", m.Version).Scan(&executed)
	if err != nil {
		return err
	}

	if executed {
		fmt.Printf("⏭️  Migration %s already executed\n", m.ID)
		return nil
	}

	fmt.Printf("🔄 Running migration %s\n", m.ID)

	if err := m.Up(db); err != nil {
		return fmt.Errorf("migration %s failed: %w", m.ID, err)
	}

	// Record migration
	_, err = db.Exec("INSERT INTO schema_migrations (version, name) VALUES ($1, $2)", m.Version, m.ID)
	if err != nil {
		return fmt.Errorf("failed to record migration %s: %w", m.ID, err)
	}

	fmt.Printf("✅ Migration %s completed\n", m.ID)
	return nil
}

// ============ UP MIGRATIONS ============

func createDiseaseTable(db *sql.DB) error {
	query := `
		CREATE TABLE disease (
			disease_id SERIAL PRIMARY KEY,
			disease_name VARCHAR(150) NOT NULL,
			category VARCHAR(100) NOT NULL,
			icd_code VARCHAR(20) UNIQUE NOT NULL,
			description TEXT
		);
	`
	_, err := db.Exec(query)
	return err
}

func createDoctorTable(db *sql.DB) error {
	query := `
		CREATE TABLE doctor (
			doctor_id INT PRIMARY KEY,
			first_name VARCHAR(100) NOT NULL,
			last_name VARCHAR(100) NOT NULL,
			email VARCHAR(150) UNIQUE NOT NULL,
			phone VARCHAR(20) NOT NULL,
			license_number VARCHAR(50) UNIQUE NOT NULL,
			specialization VARCHAR(100) NOT NULL,
			department VARCHAR(100) NOT NULL
		);
	`
	_, err := db.Exec(query)
	return err
}

func createPatientTable(db *sql.DB) error {
	query := `
		CREATE TABLE patient (
			patient_id INT PRIMARY KEY,
			first_name VARCHAR(100) NOT NULL,
			last_name VARCHAR(100) NOT NULL,
			date_of_birth DATE NOT NULL,
			gender VARCHAR(10) NOT NULL,
			email VARCHAR(150) UNIQUE NOT NULL,
			phone VARCHAR(20) NOT NULL,
			address VARCHAR(255) NOT NULL,
			emergency_contact VARCHAR(150) NOT NULL,
			blood_type VARCHAR(10) NOT NULL,
			medical_history TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(query)
	return err
}

func createMedicationTable(db *sql.DB) error {
	query := `
		CREATE TABLE medication (
			medication_id SERIAL PRIMARY KEY,
			medication_name VARCHAR(150) NOT NULL,
			generic_name VARCHAR(150) NOT NULL,
			form VARCHAR(50) NOT NULL,
			strength VARCHAR(50) NOT NULL,
			manufacturer VARCHAR(150) NOT NULL,
			side_effects TEXT
		);
	`
	_, err := db.Exec(query)
	return err
}

func createTestCatalogTable(db *sql.DB) error {
	query := `
		CREATE TABLE test_catalog (
			test_catalog_id SERIAL PRIMARY KEY,
			test_name VARCHAR(150) NOT NULL,
			test_category VARCHAR(100) NOT NULL,
			description TEXT,
			normal_range VARCHAR(100)
		);
	`
	_, err := db.Exec(query)
	return err
}

func createDiagnosisTable(db *sql.DB) error {
	query := `
		CREATE TABLE diagnosis (
			diagnosis_id SERIAL PRIMARY KEY,
			patient_id INT NOT NULL,
			disease_id INT NOT NULL,
			doctor_id INT NOT NULL,
			diagnosis_date DATE NOT NULL,
			severity VARCHAR(50) NOT NULL,
			notes TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (patient_id) REFERENCES patient(patient_id),
			FOREIGN KEY (disease_id) REFERENCES disease(disease_id),
			FOREIGN KEY (doctor_id) REFERENCES doctor(doctor_id),
			UNIQUE(patient_id, disease_id, doctor_id, diagnosis_date)
		);
	`
	_, err := db.Exec(query)
	return err
}

func createAppointmentTable(db *sql.DB) error {
	query := `
		CREATE TABLE appointment (
			appointment_id SERIAL PRIMARY KEY,
			patient_id INT NOT NULL,
			doctor_id INT NOT NULL,
			appointment_date TIMESTAMP NOT NULL,
			reason VARCHAR(255) NOT NULL,
			status VARCHAR(50) NOT NULL,
			blood_pressure VARCHAR(25),
			heart_rate VARCHAR(25),
			duration_minutes INT,
			notes TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (patient_id) REFERENCES patient(patient_id),
			FOREIGN KEY (doctor_id) REFERENCES doctor(doctor_id)
		);
	`
	_, err := db.Exec(query)
	return err
}

func createMedicalTestTable(db *sql.DB) error {
	query := `
		CREATE TABLE medical_test (
			test_id SERIAL PRIMARY KEY,
			patient_id INT NOT NULL,
			test_catalog_id INT NOT NULL,
			test_date DATE NOT NULL,
			result_value VARCHAR(100),
			normal_range VARCHAR(100),
			findings TEXT,
			result_status VARCHAR(50),
			image_url VARCHAR(500),
			interpreted_by INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (patient_id) REFERENCES patient(patient_id),
			FOREIGN KEY (test_catalog_id) REFERENCES test_catalog(test_catalog_id),
			FOREIGN KEY (interpreted_by) REFERENCES doctor(doctor_id)
		);
	`
	_, err := db.Exec(query)
	return err
}

func createPrescriptionTable(db *sql.DB) error {
	query := `
		CREATE TABLE prescription (
			prescription_id SERIAL PRIMARY KEY,
			patient_id INT NOT NULL,
			doctor_id INT NOT NULL,
			prescribed_date DATE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (patient_id) REFERENCES patient(patient_id),
			FOREIGN KEY (doctor_id) REFERENCES doctor(doctor_id)
		);
	`
	_, err := db.Exec(query)
	return err
}

func createTreatmentPlanTable(db *sql.DB) error {
	query := `
		CREATE TABLE treatment_plan (
			plan_id SERIAL PRIMARY KEY,
			patient_id INT NOT NULL,
			doctor_id INT NOT NULL,
			diagnosis VARCHAR(100) NOT NULL,
			treatment_type VARCHAR(150) NOT NULL,
			start_date DATE NOT NULL,
			end_date DATE,
			session_duration INT,
			status VARCHAR(50) NOT NULL,
			goals TEXT NOT NULL,
			progress_notes TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (patient_id) REFERENCES patient(patient_id),
			FOREIGN KEY (doctor_id) REFERENCES doctor(doctor_id)
		);
	`
	_, err := db.Exec(query)
	return err
}

func createPrescriptionMedicationTable(db *sql.DB) error {
	query := `
		CREATE TABLE prescription_medication (
			prescription_medication_id SERIAL PRIMARY KEY,
			prescription_id INT NOT NULL,
			medication_id INT NOT NULL,
			end_date DATE NOT NULL,
			frequency VARCHAR(50) NOT NULL,
			dosage VARCHAR(10) NOT NULL,
			instructions TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (prescription_id) REFERENCES prescription(prescription_id),
			FOREIGN KEY (medication_id) REFERENCES medication(medication_id),
			UNIQUE(prescription_id, medication_id)
		);
	`
	_, err := db.Exec(query)
	return err
}

func createIndexes(db *sql.DB) error {
	indexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_patient_email ON patient(email);",
		"CREATE INDEX IF NOT EXISTS idx_patient_phone ON patient(phone);",
		"CREATE INDEX IF NOT EXISTS idx_patient_date_of_birth ON patient(date_of_birth);",
		"CREATE INDEX IF NOT EXISTS idx_doctor_email ON doctor(email);",
		"CREATE INDEX IF NOT EXISTS idx_doctor_license ON doctor(license_number);",
		"CREATE INDEX IF NOT EXISTS idx_doctor_specialization ON doctor(specialization);",
		"CREATE INDEX IF NOT EXISTS idx_diagnosis_patient ON diagnosis(patient_id);",
		"CREATE INDEX IF NOT EXISTS idx_diagnosis_disease ON diagnosis(disease_id);",
		"CREATE INDEX IF NOT EXISTS idx_diagnosis_date ON diagnosis(diagnosis_date);",
		"CREATE INDEX IF NOT EXISTS idx_appointment_patient ON appointment(patient_id);",
		"CREATE INDEX IF NOT EXISTS idx_appointment_doctor ON appointment(doctor_id);",
		"CREATE INDEX IF NOT EXISTS idx_appointment_date ON appointment(appointment_date);",
		"CREATE INDEX IF NOT EXISTS idx_medical_test_patient ON medical_test(patient_id);",
		"CREATE INDEX IF NOT EXISTS idx_medical_test_catalog ON medical_test(test_catalog_id);",
		"CREATE INDEX IF NOT EXISTS idx_medical_test_date ON medical_test(test_date);",
		"CREATE INDEX IF NOT EXISTS idx_prescription_patient ON prescription(patient_id);",
		"CREATE INDEX IF NOT EXISTS idx_prescription_doctor ON prescription(doctor_id);",
		"CREATE INDEX IF NOT EXISTS idx_treatment_plan_patient ON treatment_plan(patient_id);",
		"CREATE INDEX IF NOT EXISTS idx_treatment_plan_doctor ON treatment_plan(doctor_id);",
		"CREATE INDEX IF NOT EXISTS idx_treatment_plan_status ON treatment_plan(status);",
		"CREATE INDEX IF NOT EXISTS idx_prescription_med_medication ON prescription_medication(medication_id);",
	}

	for _, idx := range indexes {
		if _, err := db.Exec(idx); err != nil {
			return err
		}
	}

	return nil
}

// ============ DOWN MIGRATIONS ============

func dropDiseaseTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS disease CASCADE;")
	return err
}

func dropDoctorTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS doctor CASCADE;")
	return err
}

func dropPatientTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS patient CASCADE;")
	return err
}

func dropMedicationTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS medication CASCADE;")
	return err
}

func dropTestCatalogTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS test_catalog CASCADE;")
	return err
}

func dropDiagnosisTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS diagnosis CASCADE;")
	return err
}

func dropAppointmentTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS appointment CASCADE;")
	return err
}

func dropMedicalTestTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS medical_test CASCADE;")
	return err
}

func dropPrescriptionTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS prescription CASCADE;")
	return err
}

func dropTreatmentPlanTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS treatment_plan CASCADE;")
	return err
}

func dropPrescriptionMedicationTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS prescription_medication CASCADE;")
	return err
}

func dropIndexes(db *sql.DB) error {
	indexes := []string{
		"DROP INDEX IF EXISTS idx_patient_email;",
		"DROP INDEX IF EXISTS idx_patient_phone;",
		"DROP INDEX IF EXISTS idx_patient_date_of_birth;",
		"DROP INDEX IF EXISTS idx_doctor_email;",
		"DROP INDEX IF EXISTS idx_doctor_license;",
		"DROP INDEX IF EXISTS idx_doctor_specialization;",
		"DROP INDEX IF EXISTS idx_diagnosis_patient;",
		"DROP INDEX IF EXISTS idx_diagnosis_disease;",
		"DROP INDEX IF EXISTS idx_diagnosis_date;",
		"DROP INDEX IF EXISTS idx_appointment_patient;",
		"DROP INDEX IF EXISTS idx_appointment_doctor;",
		"DROP INDEX IF EXISTS idx_appointment_date;",
		"DROP INDEX IF EXISTS idx_medical_test_patient;",
		"DROP INDEX IF EXISTS idx_medical_test_catalog;",
		"DROP INDEX IF EXISTS idx_medical_test_date;",
		"DROP INDEX IF EXISTS idx_prescription_patient;",
		"DROP INDEX IF EXISTS idx_prescription_doctor;",
		"DROP INDEX IF EXISTS idx_treatment_plan_patient;",
		"DROP INDEX IF EXISTS idx_treatment_plan_doctor;",
		"DROP INDEX IF EXISTS idx_treatment_plan_status;",
		"DROP INDEX IF EXISTS idx_prescription_med_medication;",
	}

	for _, idx := range indexes {
		if _, err := db.Exec(idx); err != nil {
			return err
		}
	}

	return nil
}
