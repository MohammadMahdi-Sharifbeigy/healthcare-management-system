#!/bin/bash

REPOS_DIR="pkg/domain/repositories"

# Create patient_repository.go
cat > "$REPOS_DIR/patient_repository.go" << 'REPO'
package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type PatientRepository interface {
	// CRUD operations
	Create(ctx context.Context, patient *entities.Patient) error
	GetByID(ctx context.Context, id int) (*entities.Patient, error)
	Update(ctx context.Context, patient *entities.Patient) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Patient, error)
	GetByEmail(ctx context.Context, email string) (*entities.Patient, error)
	GetByPhone(ctx context.Context, phone string) (*entities.Patient, error)
	SearchByName(ctx context.Context, name string) ([]entities.Patient, error)
}
REPO

# Create doctor_repository.go
cat > "$REPOS_DIR/doctor_repository.go" << 'REPO'
package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DoctorRepository interface {
	// CRUD operations
	Create(ctx context.Context, doctor *entities.Doctor) error
	GetByID(ctx context.Context, id int) (*entities.Doctor, error)
	Update(ctx context.Context, doctor *entities.Doctor) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Doctor, error)
	GetByEmail(ctx context.Context, email string) (*entities.Doctor, error)
	GetByLicenseNumber(ctx context.Context, licenseNumber string) (*entities.Doctor, error)
	GetBySpecialization(ctx context.Context, specialization string) ([]entities.Doctor, error)
	GetByDepartment(ctx context.Context, department string) ([]entities.Doctor, error)
}
REPO

# Create appointment_repository.go
cat > "$REPOS_DIR/appointment_repository.go" << 'REPO'
package repositories

import (
	"context"
	"time"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type AppointmentRepository interface {
	// CRUD operations
	Create(ctx context.Context, appointment *entities.Appointment) error
	GetByID(ctx context.Context, id int) (*entities.Appointment, error)
	Update(ctx context.Context, appointment *entities.Appointment) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Appointment, error)
	GetByPatientID(ctx context.Context, patientID int) ([]entities.Appointment, error)
	GetByDoctorID(ctx context.Context, doctorID int) ([]entities.Appointment, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]entities.Appointment, error)
	GetByStatus(ctx context.Context, status string) ([]entities.Appointment, error)
	CheckAvailability(ctx context.Context, doctorID int, appointmentDate time.Time) (bool, error)
}
REPO

# Create diagnosis_repository.go
cat > "$REPOS_DIR/diagnosis_repository.go" << 'REPO'
package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DiagnosisRepository interface {
	// CRUD operations
	Create(ctx context.Context, diagnosis *entities.Diagnosis) error
	GetByID(ctx context.Context, id int) (*entities.Diagnosis, error)
	Update(ctx context.Context, diagnosis *entities.Diagnosis) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Diagnosis, error)
	GetByPatientID(ctx context.Context, patientID int) ([]entities.Diagnosis, error)
	GetByDoctorID(ctx context.Context, doctorID int) ([]entities.Diagnosis, error)
	GetByDiseaseID(ctx context.Context, diseaseID int) ([]entities.Diagnosis, error)
	GetBySeverity(ctx context.Context, severity string) ([]entities.Diagnosis, error)
	GetSevereForPatient(ctx context.Context, patientID int) ([]entities.Diagnosis, error)
}
REPO

# Create disease_repository.go
cat > "$REPOS_DIR/disease_repository.go" << 'REPO'
package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DiseaseRepository interface {
	// CRUD operations
	Create(ctx context.Context, disease *entities.Disease) error
	GetByID(ctx context.Context, id int) (*entities.Disease, error)
	Update(ctx context.Context, disease *entities.Disease) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Disease, error)
	GetByName(ctx context.Context, name string) (*entities.Disease, error)
	GetByICDCode(ctx context.Context, icdCode string) (*entities.Disease, error)
	GetByCategory(ctx context.Context, category string) ([]entities.Disease, error)
	SearchByName(ctx context.Context, name string) ([]entities.Disease, error)
}
REPO

# Create medication_repository.go
cat > "$REPOS_DIR/medication_repository.go" << 'REPO'
package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type MedicationRepository interface {
	// CRUD operations
	Create(ctx context.Context, medication *entities.Medication) error
	GetByID(ctx context.Context, id int) (*entities.Medication, error)
	Update(ctx context.Context, medication *entities.Medication) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Medication, error)
	GetByName(ctx context.Context, name string) (*entities.Medication, error)
	GetByGenericName(ctx context.Context, genericName string) (*entities.Medication, error)
	SearchByName(ctx context.Context, name string) ([]entities.Medication, error)
	GetByForm(ctx context.Context, form string) ([]entities.Medication, error)
}
REPO

# Create prescription_repository.go
cat > "$REPOS_DIR/prescription_repository.go" << 'REPO'
package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type PrescriptionRepository interface {
	// CRUD operations
	Create(ctx context.Context, prescription *entities.Prescription) error
	GetByID(ctx context.Context, id int) (*entities.Prescription, error)
	Update(ctx context.Context, prescription *entities.Prescription) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Prescription, error)
	GetByPatientID(ctx context.Context, patientID int) ([]entities.Prescription, error)
	GetByDoctorID(ctx context.Context, doctorID int) ([]entities.Prescription, error)
	GetActivePrescriptionsByPatient(ctx context.Context, patientID int) ([]entities.Prescription, error)
	GetExpiringPrescriptionsByPatient(ctx context.Context, patientID int, days int) ([]entities.Prescription, error)
}
REPO

# Create prescription_medication_repository.go
cat > "$REPOS_DIR/prescription_medication_repository.go" << 'REPO'
package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type PrescriptionMedicationRepository interface {
	// CRUD operations
	Create(ctx context.Context, pm *entities.PrescriptionMedication) error
	GetByID(ctx context.Context, id int) (*entities.PrescriptionMedication, error)
	Update(ctx context.Context, pm *entities.PrescriptionMedication) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.PrescriptionMedication, error)
	GetByPrescriptionID(ctx context.Context, prescriptionID int) ([]entities.PrescriptionMedication, error)
	GetByMedicationID(ctx context.Context, medicationID int) ([]entities.PrescriptionMedication, error)
	DeleteByPrescriptionAndMedication(ctx context.Context, prescriptionID, medicationID int) error
}
REPO

# Create medical_test_repository.go
cat > "$REPOS_DIR/medical_test_repository.go" << 'REPO'
package repositories

import (
	"context"
	"time"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type MedicalTestRepository interface {
	// CRUD operations
	Create(ctx context.Context, test *entities.MedicalTest) error
	GetByID(ctx context.Context, id int) (*entities.MedicalTest, error)
	Update(ctx context.Context, test *entities.MedicalTest) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.MedicalTest, error)
	GetByPatientID(ctx context.Context, patientID int) ([]entities.MedicalTest, error)
	GetByTestCatalogID(ctx context.Context, testCatalogID int) ([]entities.MedicalTest, error)
	GetAbnormalResults(ctx context.Context) ([]entities.MedicalTest, error)
	GetCriticalResults(ctx context.Context) ([]entities.MedicalTest, error)
	GetByPatientAndTestType(ctx context.Context, patientID, testCatalogID int) ([]entities.MedicalTest, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]entities.MedicalTest, error)
	GetByResultStatus(ctx context.Context, status string) ([]entities.MedicalTest, error)
}
REPO

# Create test_catalog_repository.go
cat > "$REPOS_DIR/test_catalog_repository.go" << 'REPO'
package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type TestCatalogRepository interface {
	// CRUD operations
	Create(ctx context.Context, catalog *entities.TestCatalog) error
	GetByID(ctx context.Context, id int) (*entities.TestCatalog, error)
	Update(ctx context.Context, catalog *entities.TestCatalog) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.TestCatalog, error)
	GetByName(ctx context.Context, name string) (*entities.TestCatalog, error)
	GetByCategory(ctx context.Context, category string) ([]entities.TestCatalog, error)
	SearchByName(ctx context.Context, name string) ([]entities.TestCatalog, error)
}
REPO

# Create treatment_plan_repository.go
cat > "$REPOS_DIR/treatment_plan_repository.go" << 'REPO'
package repositories

import (
	"context"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type TreatmentPlanRepository interface {
	// CRUD operations
	Create(ctx context.Context, plan *entities.TreatmentPlan) error
	GetByID(ctx context.Context, id int) (*entities.TreatmentPlan, error)
	Update(ctx context.Context, plan *entities.TreatmentPlan) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.TreatmentPlan, error)
	GetByPatientID(ctx context.Context, patientID int) ([]entities.TreatmentPlan, error)
	GetByDoctorID(ctx context.Context, doctorID int) ([]entities.TreatmentPlan, error)
	GetActiveByPatient(ctx context.Context, patientID int) ([]entities.TreatmentPlan, error)
	GetByStatus(ctx context.Context, status string) ([]entities.TreatmentPlan, error)
	GetActivePlans(ctx context.Context) ([]entities.TreatmentPlan, error)
}
REPO

echo "✓ All 11 repository interfaces created successfully!"
echo "✓ Files created in: pkg/domain/repositories/"
echo ""
echo "Repositories created:"
echo "  1. patient_repository.go"
echo "  2. doctor_repository.go"
echo "  3. appointment_repository.go"
echo "  4. diagnosis_repository.go"
echo "  5. disease_repository.go"
echo "  6. medication_repository.go"
echo "  7. prescription_repository.go"
echo "  8. prescription_medication_repository.go"
echo "  9. medical_test_repository.go"
echo "  10. test_catalog_repository.go"
echo "  11. treatment_plan_repository.go"
echo ""
echo "Next: Implement repositories in pkg/infrastructure/database/postgres/"
