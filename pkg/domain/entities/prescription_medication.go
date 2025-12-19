package entities

import (
	"time"
)

type PrescriptionMedication struct {
	PrescriptionMedicationID int
	PrescriptionID           int
	MedicationID             int
	Dosage                   string
	Frequency                string
	Instructions             string
	EndDate                  time.Time
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

func (pm *PrescriptionMedication) Validate() error {
	if pm.PrescriptionID == 0 || pm.MedicationID == 0 {
		return ErrInvalidInput("prescription and medication required")
	}
	if pm.Dosage == "" {
		return ErrInvalidInput("dosage required")
	}
	if pm.Frequency == "" {
		return ErrInvalidInput("frequency required")
	}
	if pm.EndDate.IsZero() {
		return ErrInvalidInput("end date required")
	}
	return nil
}

func (pm *PrescriptionMedication) IsExpired() bool {
	return time.Now().After(pm.EndDate)
}

func (pm *PrescriptionMedication) IsActive() bool {
	return !pm.IsExpired()
}
