package entities

import (
	"regexp"
	"time"
)

type Patient struct {
	PatientID        int
	FirstName        string
	LastName         string
	DateOfBirth      time.Time
	Gender           string
	Email            string
	Phone            string
	Address          string
	EmergencyContact string
	BloodType        string
	MedicalHistory   string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (p *Patient) Validate() error {
	if p.FirstName == "" || p.LastName == "" {
		return ErrInvalidInput("first and last name required")
	}
	if p.DateOfBirth.IsZero() {
		return ErrInvalidInput("date of birth required")
	}
	if !isValidEmail(p.Email) {
		return ErrInvalidInput("invalid email format")
	}
	if !isValidPhone(p.Phone) {
		return ErrInvalidInput("invalid phone format")
	}
	if p.BloodType == "" {
		return ErrInvalidInput("blood type required")
	}
	return nil
}

func (p *Patient) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Patient) GetAge() int {
	return int(time.Since(p.DateOfBirth).Hours() / 24 / 365)
}

func (p *Patient) IsAdult() bool {
	return p.GetAge() >= 18
}

func isValidEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailPattern, email)
	return match
}

func isValidPhone(phone string) bool {
	// Simple validation: at least 10 digits
	digits := 0
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			digits++
		}
	}
	return digits >= 10
}
