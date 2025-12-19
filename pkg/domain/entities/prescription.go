package entities

import (
	"time"
)

type Prescription struct {
	PrescriptionID int
	PatientID      int
	DoctorID       int
	PrescribedDate time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (p *Prescription) Validate() error {
	if p.PatientID == 0 || p.DoctorID == 0 {
		return ErrInvalidInput("patient and doctor required")
	}
	if p.PrescribedDate.IsZero() {
		return ErrInvalidInput("prescribed date required")
	}
	return nil
}
