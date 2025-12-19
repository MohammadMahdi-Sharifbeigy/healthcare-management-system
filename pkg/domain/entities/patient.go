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
		return ErrInvalidInput("first name and last name required")
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

func (p *Patient) GetAge() int {
	return int(time.Since(p.DateOfBirth).Hours() / 24 / 365)
}

func (p *Patient) IsAdult() bool {
	return p.GetAge() >= 18
}

func (p *Patient) FullName() string {
	return p.FirstName + " " + p.LastName
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func isValidPhone(phone string) bool {
	return len(phone) >= 10
}
