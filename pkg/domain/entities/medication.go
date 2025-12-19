package entities

import (
	"time"
)

type Medication struct {
	MedicationID int
	MedicationName string
	GenericName  string
	Form         string // tablet, capsule, injection, etc
	Strength     string
	Manufacturer string
	SideEffects  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *Medication) Validate() error {
	if m.MedicationName == "" {
		return ErrInvalidInput("medication name required")
	}
	if m.GenericName == "" {
		return ErrInvalidInput("generic name required")
	}
	if m.Form == "" {
		return ErrInvalidInput("form required")
	}
	if m.Strength == "" {
		return ErrInvalidInput("strength required")
	}
	return nil
}

func (m *Medication) FullName() string {
	return m.MedicationName + " (" + m.GenericName + ") " + m.Strength + " " + m.Form
}
