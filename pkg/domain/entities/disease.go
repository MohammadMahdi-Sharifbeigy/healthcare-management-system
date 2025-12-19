package entities

import (
	"time"
)

type Disease struct {
	DiseaseID   int
	DiseaseName string
	Category    string
	ICDCode     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (d *Disease) Validate() error {
	if d.DiseaseName == "" {
		return ErrInvalidInput("disease name required")
	}
	if d.ICDCode == "" {
		return ErrInvalidInput("ICD code required")
	}
	if d.Category == "" {
		return ErrInvalidInput("category required")
	}
	return nil
}
