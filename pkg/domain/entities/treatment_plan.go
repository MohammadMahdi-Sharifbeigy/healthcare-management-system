package entities

import (
	"time"
)

type TreatmentPlan struct {
	PlanID          int
	PatientID       int
	DoctorID        int
	DiagnosisID     int
	Diagnosis       string
	TreatmentType   string // physiotherapy, medication, etc
	StartDate       time.Time
	EndDate         time.Time
	SessionDuration int
	Status          string // active, completed, cancelled
	Goals           string
	ProgressNotes   string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (tp *TreatmentPlan) Validate() error {
	if tp.PatientID == 0 || tp.DoctorID == 0 {
		return ErrInvalidInput("patient and doctor required")
	}
	if tp.StartDate.IsZero() {
		return ErrInvalidInput("start date required")
	}
	if tp.TreatmentType == "" {
		return ErrInvalidInput("treatment type required")
	}
	if tp.Goals == "" {
		return ErrInvalidInput("goals required")
	}
	return nil
}

func (tp *TreatmentPlan) IsActive() bool {
	return tp.Status == "active"
}

func (tp *TreatmentPlan) IsCompleted() bool {
	return tp.Status == "completed"
}

func (tp *TreatmentPlan) DaysRemaining() int {
	remaining := int(time.Until(tp.EndDate).Hours() / 24)
	if remaining < 0 {
		return 0
	}
	return remaining
}
