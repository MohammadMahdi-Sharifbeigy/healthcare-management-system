package entities

import (
	"time"
)

type Appointment struct {
	AppointmentID int
	PatientID     int
	DoctorID      int
	AppointmentDate time.Time
	Reason        string
	Status        string // scheduled, completed, cancelled
	BloodPressure string
	HeartRate     string
	DurationMinutes int
	Notes         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (a *Appointment) Validate() error {
	if a.PatientID == 0 || a.DoctorID == 0 {
		return ErrInvalidInput("patient and doctor required")
	}
	if a.AppointmentDate.IsZero() {
		return ErrInvalidInput("appointment date required")
	}
	if a.Reason == "" {
		return ErrInvalidInput("reason required")
	}
	return nil
}

func (a *Appointment) IsCompleted() bool {
	return a.Status == "completed"
}

func (a *Appointment) IsCancelled() bool {
	return a.Status == "cancelled"
}

func (a *Appointment) IsScheduled() bool {
	return a.Status == "scheduled"
}
