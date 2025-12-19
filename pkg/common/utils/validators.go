package utils

import (
	"regexp"
	"strings"
	"time"
)

// ValidateEmail checks if email format is valid
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}
	
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

// ValidatePhone checks if phone format is valid (basic validation)
func ValidatePhone(phone string) bool {
	if len(phone) < 10 || len(phone) > 15 {
		return false
	}
	
	// Remove common separators
	cleaned := strings.NewReplacer(
		"-", "",
		" ", "",
		"(", "",
		")", "",
		"+", "",
	).Replace(phone)
	
	// Check if only digits remain
	for _, ch := range cleaned {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	
	return len(cleaned) >= 10
}

// ValidateBloodType checks if blood type is valid
func ValidateBloodType(bloodType string) bool {
	validTypes := map[string]bool{
		"O+":  true,
		"O-":  true,
		"A+":  true,
		"A-":  true,
		"B+":  true,
		"B-":  true,
		"AB+": true,
		"AB-": true,
	}
	return validTypes[bloodType]
}

// ValidateGender checks if gender is valid
func ValidateGender(gender string) bool {
	validGenders := map[string]bool{
		"Male":   true,
		"Female": true,
		"Other":  true,
	}
	return validGenders[gender]
}

// ValidateDateOfBirth validates date of birth
func ValidateDateOfBirth(dateOfBirth string) bool {
	parsedDate, err := time.Parse("2006-01-02", dateOfBirth)
	if err != nil {
		return false
	}
	
	// Check if date is in the past
	if parsedDate.After(time.Now()) {
		return false
	}
	
	// Check if person is at least 0 years old (not negative)
	age := time.Now().Year() - parsedDate.Year()
	if age < 0 || age > 150 {
		return false
	}
	
	return true
}

// ValidateAppointmentDate validates appointment date
func ValidateAppointmentDate(appointmentDate string) bool {
	parsedDate, err := time.Parse(time.RFC3339, appointmentDate)
	if err != nil {
		return false
	}
	
	// Appointment must be in the future
	if parsedDate.Before(time.Now()) {
		return false
	}
	
	// Appointment should not be more than 2 years in future
	twoYearsFromNow := time.Now().AddDate(2, 0, 0)
	if parsedDate.After(twoYearsFromNow) {
		return false
	}
	
	return true
}

// ValidateDateRange validates that start date is before end date
func ValidateDateRange(startDate, endDate string) bool {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return false
	}
	
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return false
	}
	
	return start.Before(end) || start.Equal(end)
}

// ValidateICD10Code checks if ICD-10 code format is valid
func ValidateICD10Code(code string) bool {
	if len(code) < 3 || len(code) > 7 {
		return false
	}
	
	// Format: Letter + 2 digits [. + 1-2 alphanumeric]
	pattern := `^[A-Z]{1}[0-9]{2}(\.[A-Z0-9]{1,2})?$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(code)
}

// ValidateDoctorID checks if doctor ID format is valid (8 digits)
func ValidateDoctorID(doctorID int) bool {
	return doctorID >= 10000000 && doctorID <= 99999999
}

// ValidateSeverity checks if severity level is valid
func ValidateSeverity(severity string) bool {
	validSeverities := map[string]bool{
		"mild":     true,
		"moderate": true,
		"severe":   true,
	}
	return validSeverities[strings.ToLower(severity)]
}

// ValidateAppointmentStatus checks if status is valid
func ValidateAppointmentStatus(status string) bool {
	validStatuses := map[string]bool{
		"confirmed": true,
		"completed": true,
		"cancelled": true,
	}
	return validStatuses[strings.ToLower(status)]
}

// ValidateTestResultStatus checks if result status is valid
func ValidateTestResultStatus(status string) bool {
	validStatuses := map[string]bool{
		"normal":    true,
		"abnormal":  true,
		"critical":  true,
	}
	return validStatuses[strings.ToLower(status)]
}

// ValidateMedicationForm checks if medication form is valid
func ValidateMedicationForm(form string) bool {
	validForms := map[string]bool{
		"tablet":     true,
		"capsule":    true,
		"injection":  true,
		"liquid":     true,
	}
	return validForms[strings.ToLower(form)]
}

// ValidateTreatmentType checks if treatment type is valid
func ValidateTreatmentType(treatmentType string) bool {
	validTypes := map[string]bool{
		"pharmacotherapy": true,
		"physiotherapy":   true,
		"psychotherapy":   true,
		"occupational":    true,
		"speech":          true,
		"rehabilitation":  true,
	}
	return validTypes[strings.ToLower(treatmentType)]
}

// ValidatePlanStatus checks if plan status is valid
func ValidatePlanStatus(status string) bool {
	validStatuses := map[string]bool{
		"active":    true,
		"completed": true,
		"cancelled": true,
	}
	return validStatuses[strings.ToLower(status)]
}

// ValidateStringLength checks if string length is within bounds
func ValidateStringLength(str string, min, max int) bool {
	length := len(str)
	return length >= min && length <= max
}

// ValidateIntRange checks if integer is within range
func ValidateIntRange(value, min, max int) bool {
	return value >= min && value <= max
}

// ValidateURL checks if URL format is valid
func ValidateURL(url string) bool {
	if url == "" {
		return false
	}
	
	pattern := `^(https?://)?([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.[a-zA-Z]{2,}(/.*)?$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}

// ValidateNotEmpty checks if string is not empty
func ValidateNotEmpty(str string) bool {
	return strings.TrimSpace(str) != ""
}

// ValidatePaginationParams checks if pagination parameters are valid
func ValidatePaginationParams(limit, offset int) bool {
	return limit > 0 && limit <= 100 && offset >= 0
}

// ValidateSessionDuration checks if session duration is valid
func ValidateSessionDuration(duration int) bool {
	return duration > 0 && duration <= 480 // Max 8 hours = 480 minutes
}

// ValidateFrequency checks if medication frequency is valid
func ValidateFrequency(frequency string) bool {
	validFrequencies := []string{
		"once daily",
		"twice daily",
		"three times daily",
		"four times daily",
		"every 4 hours",
		"every 6 hours",
		"every 8 hours",
		"every 12 hours",
		"as needed",
	}
	
	for _, f := range validFrequencies {
		if strings.EqualFold(frequency, f) {
			return true
		}
	}
	return false
}
