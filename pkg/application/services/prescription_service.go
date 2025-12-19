package services

import (
	"context"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type PrescriptionService struct {
	prescriptionRepo          repositories.PrescriptionRepository
	prescriptionMedicationRepo repositories.PrescriptionMedicationRepository
	patientRepo               repositories.PatientRepository
	doctorRepo                repositories.DoctorRepository
	medicationRepo            repositories.MedicationRepository
}

func NewPrescriptionService(
	prescriptionRepo repositories.PrescriptionRepository,
	prescriptionMedicationRepo repositories.PrescriptionMedicationRepository,
	patientRepo repositories.PatientRepository,
	doctorRepo repositories.DoctorRepository,
	medicationRepo repositories.MedicationRepository,
) *PrescriptionService {
	return &PrescriptionService{
		prescriptionRepo:           prescriptionRepo,
		prescriptionMedicationRepo: prescriptionMedicationRepo,
		patientRepo:                patientRepo,
		doctorRepo:                 doctorRepo,
		medicationRepo:             medicationRepo,
	}
}

func (s *PrescriptionService) CreatePrescription(ctx context.Context, prescription *entities.Prescription) error {
	if err := prescription.Validate(); err != nil {
		return err
	}

	// Verify patient exists
	if _, err := s.patientRepo.GetByID(ctx, prescription.PatientID); err != nil {
		return entities.ErrNotFound("patient not found")
	}

	// Verify doctor exists
	if _, err := s.doctorRepo.GetByID(ctx, prescription.DoctorID); err != nil {
		return entities.ErrNotFound("doctor not found")
	}

	return s.prescriptionRepo.Create(ctx, prescription)
}

func (s *PrescriptionService) GetPrescription(ctx context.Context, prescriptionID int) (*entities.Prescription, error) {
	if prescriptionID == 0 {
		return nil, entities.ErrInvalidInput("prescription ID required")
	}

	return s.prescriptionRepo.GetByID(ctx, prescriptionID)
}

func (s *PrescriptionService) UpdatePrescription(ctx context.Context, prescription *entities.Prescription) error {
	if err := prescription.Validate(); err != nil {
		return err
	}

	existing, err := s.prescriptionRepo.GetByID(ctx, prescription.PrescriptionID)
	if err != nil {
		return err
	}
	if existing == nil {
		return entities.ErrNotFound("prescription not found")
	}

	return s.prescriptionRepo.Update(ctx, prescription)
}

func (s *PrescriptionService) AddMedicationToPrescription(ctx context.Context, pm *entities.PrescriptionMedication) error {
	if err := pm.Validate(); err != nil {
		return err
	}

	// Verify prescription exists
	if _, err := s.prescriptionRepo.GetByID(ctx, pm.PrescriptionID); err != nil {
		return entities.ErrNotFound("prescription not found")
	}

	// Verify medication exists
	if _, err := s.medicationRepo.GetByID(ctx, pm.MedicationID); err != nil {
		return entities.ErrNotFound("medication not found")
	}

	return s.prescriptionMedicationRepo.Create(ctx, pm)
}

func (s *PrescriptionService) RemoveMedicationFromPrescription(ctx context.Context, prescriptionID, medicationID int) error {
	if prescriptionID == 0 || medicationID == 0 {
		return entities.ErrInvalidInput("prescription and medication IDs required")
	}

	return s.prescriptionMedicationRepo.DeleteByPrescriptionAndMedication(ctx, prescriptionID, medicationID)
}

func (s *PrescriptionService) GetActivePrescriptions(ctx context.Context, patientID int) ([]entities.Prescription, error) {
	if patientID == 0 {
		return nil, entities.ErrInvalidInput("patient ID required")
	}

	return s.prescriptionRepo.GetActivePrescriptionsByPatient(ctx, patientID)
}

func (s *PrescriptionService) GetExpiringPrescriptions(ctx context.Context, patientID int, daysAhead int) ([]entities.Prescription, error) {
	if patientID == 0 {
		return nil, entities.ErrInvalidInput("patient ID required")
	}
	if daysAhead <= 0 {
		daysAhead = 7
	}

	return s.prescriptionRepo.GetExpiringPrescriptionsByPatient(ctx, patientID, daysAhead)
}

func (s *PrescriptionService) GetPrescriptionMedications(ctx context.Context, prescriptionID int) ([]entities.PrescriptionMedication, error) {
	if prescriptionID == 0 {
		return nil, entities.ErrInvalidInput("prescription ID required")
	}

	return s.prescriptionMedicationRepo.GetByPrescriptionID(ctx, prescriptionID)
}
