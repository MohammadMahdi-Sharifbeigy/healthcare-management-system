package entities

import "time"

type Doctor struct {
	DoctorID       int
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	LicenseNumber  string
	Specialization string
	Department     string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (d *Doctor) Validate() error {
	if d.FirstName == "" || d.LastName == "" {
		return ErrInvalidInput("first and last name required")
	}
	if !isValidEmail(d.Email) {
		return ErrInvalidInput("invalid email format")
	}
	if !isValidPhone(d.Phone) {
		return ErrInvalidInput("invalid phone format")
	}
	if d.LicenseNumber == "" {
		return ErrInvalidInput("license number required")
	}
	if d.Specialization == "" {
		return ErrInvalidInput("specialization required")
	}
	if d.Department == "" {
		return ErrInvalidInput("department required")
	}
	return nil
}

func (d *Doctor) FullName() string {
	return d.FirstName + " " + d.LastName
}

func (d *Doctor) GetTitle() string {
	return "Dr. " + d.FullName()
}
