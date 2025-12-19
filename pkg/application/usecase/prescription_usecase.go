package usecase

import (
	"context"
	"time"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/dto"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type PrescriptionUseCase struct {
	service *services.PrescriptionService
}

func NewPrescriptionUseCase(service *services.PrescriptionService) *PrescriptionUseCase {
	return &PrescriptionUseCase{service: service}
}

// CreatePrescriptionWithMedications creates prescription and adds medications
func (uc *PrescriptionUseCase) CreatePrescriptionWithMedications(ctx context.Context, req *dto.CreatePrescriptionRequest) (*dto.PrescriptionResponse, error) {
	prescribedDate, _ := time.Parse("2006-01-02", req.PrescribedDate)

	prescription := &entities.Prescription{
		PatientID:      req.PatientID,
		DoctorID:       req.DoctorID,
		PrescribedDate: prescribedDate,
	}

	if err := uc.service.CreatePrescription(ctx, prescription); err != nil {
		return nil, err
	}

	return &dto.PrescriptionResponse{
		PrescriptionID: prescription.PrescriptionID,
		PatientID:      prescription.PatientID,
		DoctorID:       prescription.DoctorID,
		PrescribedDate: prescription.PrescribedDate,
		CreatedAt:      prescription.CreatedAt,
		UpdatedAt:      prescription.UpdatedAt,
	}, nil
}

// AddMedicationToPrescription adds medication to prescription
func (uc *PrescriptionUseCase) AddMedicationToPrescription(ctx context.Context, prescriptionID int, req *dto.AddMedicationRequest) error {
	endDate, _ := time.Parse("2006-01-02", req.EndDate)

	pm := &entities.PrescriptionMedication{
		PrescriptionID: prescriptionID,
		MedicationID:   req.MedicationID,
		Dosage:         req.Dosage,
		Frequency:      req.Frequency,
		Instructions:   req.Instructions,
		EndDate:        endDate,
	}

	return uc.service.AddMedicationToPrescription(ctx, pm)
}

// RemoveMedicationFromPrescription removes medication from prescription
func (uc *PrescriptionUseCase) RemoveMedicationFromPrescription(ctx context.Context, prescriptionID, medicationID int) error {
	return uc.service.RemoveMedicationFromPrescription(ctx, prescriptionID, medicationID)
}

// GetActiveMedicationsForPatient retrieves current active medications
func (uc *PrescriptionUseCase) GetActiveMedicationsForPatient(ctx context.Context, patientID int) ([]*dto.MedicationInPrescription, error) {
	prescriptions, err := uc.service.GetActivePrescriptions(ctx, patientID)
	if err != nil {
		return nil, err
	}

	var medications []*dto.MedicationInPrescription
	for _, p := range prescriptions {
		meds, _ := uc.service.GetPrescriptionMedications(ctx, p.PrescriptionID)
		for _, med := range meds {
			medications = append(medications, &dto.MedicationInPrescription{
				MedicationID: med.MedicationID,
				Dosage:       med.Dosage,
				Frequency:    med.Frequency,
				Instructions: med.Instructions,
				EndDate:      med.EndDate,
				Status:       "active",
			})
		}
	}

	return medications, nil
}

// GetExpiringMedicationsAlert retrieves expiring prescriptions
func (uc *PrescriptionUseCase) GetExpiringMedicationsAlert(ctx context.Context, patientID, daysAhead int) ([]*dto.MedicationInPrescription, error) {
	prescriptions, err := uc.service.GetExpiringPrescriptions(ctx, patientID, daysAhead)
	if err != nil {
		return nil, err
	}

	var medications []*dto.MedicationInPrescription
	for _, p := range prescriptions {
		meds, _ := uc.service.GetPrescriptionMedications(ctx, p.PrescriptionID)
		for _, med := range meds {
			medications = append(medications, &dto.MedicationInPrescription{
				MedicationID:  med.MedicationID,
				Dosage:        med.Dosage,
				Frequency:     med.Frequency,
				Instructions:  med.Instructions,
				EndDate:       med.EndDate,
				DaysRemaining: int(med.EndDate.Sub(time.Now()).Hours() / 24),
				Status:        "expiring",
			})
		}
	}

	return medications, nil
}
