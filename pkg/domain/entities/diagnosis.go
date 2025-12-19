package entities

import (
	"time"
)

type Diagnosis struct {
	DiagnosisID int
	PatientID   int
	DoctorID    int
	DiseaseID   int
	DiagnosisDate time.Time
	Severity    string // mild, moderate, severe
	Notes       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (d *Diagnosis) Validate() error {
	if d.PatientID == 0 || d.DoctorID == 0 || d.DiseaseID == 0 {
		return ErrInvalidInput("patient, doctor, and disease required")
	}
	if d.DiagnosisDate.IsZero() {
		return ErrInvalidInput("diagnosis date required")
	}
	if d.Severity == "" {
		return ErrInvalidInput("severity required")
	}
	validSeverities := map[string]bool{"mild": true, "moderate": true, "severe": true}
	if !validSeverities[d.Severity] {
		return ErrInvalidInput("invalid severity")
	}
	return nil
}

func (d *Diagnosis) IsSevere() bool {
	return d.Severity == "severe"
}

func (d *Diagnosis) IsCritical() bool {
	return d.IsSevere()
}
