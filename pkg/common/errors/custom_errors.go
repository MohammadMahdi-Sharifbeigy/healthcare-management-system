package errors

import (
	"fmt"
	"net/http"
)

// ErrorType defines type of error
type ErrorType string

const (
	ValidationError    ErrorType = "VALIDATION_ERROR"
	NotFoundError      ErrorType = "NOT_FOUND_ERROR"
	ConflictError      ErrorType = "CONFLICT_ERROR"
	UnauthorizedError  ErrorType = "UNAUTHORIZED_ERROR"
	ForbiddenError     ErrorType = "FORBIDDEN_ERROR"
	InternalError      ErrorType = "INTERNAL_ERROR"
	BadRequestError    ErrorType = "BAD_REQUEST_ERROR"
	DatabaseError      ErrorType = "DATABASE_ERROR"
	DuplicateError     ErrorType = "DUPLICATE_ERROR"
	InvalidStateError  ErrorType = "INVALID_STATE_ERROR"
)

// AppError represents application error with context
type AppError struct {
	Type       ErrorType
	Message    string
	StatusCode int
	Details    map[string]interface{}
	Err        error
}

// NewAppError creates new application error
func NewAppError(errorType ErrorType, message string) *AppError {
	return &AppError{
		Type:       errorType,
		Message:    message,
		Details:    make(map[string]interface{}),
		StatusCode: getStatusCode(errorType),
	}
}

// WithStatusCode sets custom status code
func (e *AppError) WithStatusCode(code int) *AppError {
	e.StatusCode = code
	return e
}

// WithDetails adds error details
func (e *AppError) WithDetails(key string, value interface{}) *AppError {
	e.Details[key] = value
	return e
}

// WithError wraps underlying error
func (e *AppError) WithError(err error) *AppError {
	e.Err = err
	return e
}

// Error implements error interface
func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// PatientNotFoundError returns patient not found error
func PatientNotFoundError(patientID int) *AppError {
	return NewAppError(NotFoundError, fmt.Sprintf("Patient with ID %d not found", patientID)).
		WithDetails("patient_id", patientID).
		WithStatusCode(http.StatusNotFound)
}

// DoctorNotFoundError returns doctor not found error
func DoctorNotFoundError(doctorID int) *AppError {
	return NewAppError(NotFoundError, fmt.Sprintf("Doctor with ID %d not found", doctorID)).
		WithDetails("doctor_id", doctorID).
		WithStatusCode(http.StatusNotFound)
}

// AppointmentNotFoundError returns appointment not found error
func AppointmentNotFoundError(appointmentID int) *AppError {
	return NewAppError(NotFoundError, fmt.Sprintf("Appointment with ID %d not found", appointmentID)).
		WithDetails("appointment_id", appointmentID).
		WithStatusCode(http.StatusNotFound)
}

// DiagnosisNotFoundError returns diagnosis not found error
func DiagnosisNotFoundError(diagnosisID int) *AppError {
	return NewAppError(NotFoundError, fmt.Sprintf("Diagnosis with ID %d not found", diagnosisID)).
		WithDetails("diagnosis_id", diagnosisID).
		WithStatusCode(http.StatusNotFound)
}

// DiseaseNotFoundError returns disease not found error
func DiseaseNotFoundError(diseaseID int) *AppError {
	return NewAppError(NotFoundError, fmt.Sprintf("Disease with ID %d not found", diseaseID)).
		WithDetails("disease_id", diseaseID).
		WithStatusCode(http.StatusNotFound)
}

// MedicationNotFoundError returns medication not found error
func MedicationNotFoundError(medicationID int) *AppError {
	return NewAppError(NotFoundError, fmt.Sprintf("Medication with ID %d not found", medicationID)).
		WithDetails("medication_id", medicationID).
		WithStatusCode(http.StatusNotFound)
}

// PrescriptionNotFoundError returns prescription not found error
func PrescriptionNotFoundError(prescriptionID int) *AppError {
	return NewAppError(NotFoundError, fmt.Sprintf("Prescription with ID %d not found", prescriptionID)).
		WithDetails("prescription_id", prescriptionID).
		WithStatusCode(http.StatusNotFound)
}

// MedicalTestNotFoundError returns medical test not found error
func MedicalTestNotFoundError(testID int) *AppError {
	return NewAppError(NotFoundError, fmt.Sprintf("Medical test with ID %d not found", testID)).
		WithDetails("test_id", testID).
		WithStatusCode(http.StatusNotFound)
}

// TreatmentPlanNotFoundError returns treatment plan not found error
func TreatmentPlanNotFoundError(planID int) *AppError {
	return NewAppError(NotFoundError, fmt.Sprintf("Treatment plan with ID %d not found", planID)).
		WithDetails("plan_id", planID).
		WithStatusCode(http.StatusNotFound)
}

// DuplicateEmailError returns duplicate email error
func DuplicateEmailError(email string) *AppError {
	return NewAppError(ConflictError, fmt.Sprintf("Email %s already exists", email)).
		WithDetails("email", email).
		WithStatusCode(http.StatusConflict)
}

// DuplicateLicenseError returns duplicate license error
func DuplicateLicenseError(licenseNumber string) *AppError {
	return NewAppError(ConflictError, fmt.Sprintf("License number %s already exists", licenseNumber)).
		WithDetails("license_number", licenseNumber).
		WithStatusCode(http.StatusConflict)
}

// DuplicateDiagnosisError returns duplicate diagnosis error
func DuplicateDiagnosisError(patientID, diseaseID int) *AppError {
	return NewAppError(ConflictError, "This diagnosis already exists for this patient").
		WithDetails("patient_id", patientID).
		WithDetails("disease_id", diseaseID).
		WithStatusCode(http.StatusConflict)
}

// InvalidEmailError returns invalid email error
func InvalidEmailError(email string) *AppError {
	return NewAppError(ValidationError, fmt.Sprintf("Invalid email format: %s", email)).
		WithDetails("field", "email").
		WithStatusCode(http.StatusBadRequest)
}

// InvalidPhoneError returns invalid phone error
func InvalidPhoneError(phone string) *AppError {
	return NewAppError(ValidationError, fmt.Sprintf("Invalid phone format: %s", phone)).
		WithDetails("field", "phone").
		WithStatusCode(http.StatusBadRequest)
}

// InvalidBloodTypeError returns invalid blood type error
func InvalidBloodTypeError(bloodType string) *AppError {
	return NewAppError(ValidationError, fmt.Sprintf("Invalid blood type: %s", bloodType)).
		WithDetails("field", "blood_type").
		WithStatusCode(http.StatusBadRequest)
}

// InvalidGenderError returns invalid gender error
func InvalidGenderError(gender string) *AppError {
	return NewAppError(ValidationError, fmt.Sprintf("Invalid gender: %s", gender)).
		WithDetails("field", "gender").
		WithStatusCode(http.StatusBadRequest)
}

// InvalidSeverityError returns invalid severity error
func InvalidSeverityError(severity string) *AppError {
	return NewAppError(ValidationError, fmt.Sprintf("Invalid severity: %s", severity)).
		WithDetails("field", "severity").
		WithStatusCode(http.StatusBadRequest)
}

// InvalidDateError returns invalid date error
func InvalidDateError(field, dateStr string) *AppError {
	return NewAppError(ValidationError, fmt.Sprintf("Invalid date format for %s: %s", field, dateStr)).
		WithDetails("field", field).
		WithStatusCode(http.StatusBadRequest)
}

// AppointmentConflictError returns appointment conflict error
func AppointmentConflictError(doctorID int, dateTime string) *AppError {
	return NewAppError(ConflictError, "Doctor is not available at this time").
		WithDetails("doctor_id", doctorID).
		WithDetails("requested_time", dateTime).
		WithStatusCode(http.StatusConflict)
}

// DrugInteractionError returns drug interaction error
func DrugInteractionError(drug1, drug2, level string) *AppError {
	return NewAppError(InvalidStateError, fmt.Sprintf("Drug interaction found: %s and %s (%s)", drug1, drug2, level)).
		WithDetails("drug1", drug1).
		WithDetails("drug2", drug2).
		WithDetails("interaction_level", level).
		WithStatusCode(http.StatusBadRequest)
}

// InvalidVitalsError returns invalid vitals error
func InvalidVitalsError(field, value string) *AppError {
	return NewAppError(ValidationError, fmt.Sprintf("Invalid vital sign value for %s: %s", field, value)).
		WithDetails("field", field).
		WithStatusCode(http.StatusBadRequest)
}

// DatabaseError returns database error
func DatabaseError(operation string, err error) *AppError {
	return NewAppError(DatabaseError, fmt.Sprintf("Database error during %s", operation)).
		WithError(err).
		WithStatusCode(http.StatusInternalServerError)
}

// UnauthorizedError returns unauthorized error
func UnauthorizedAccessError(reason string) *AppError {
	return NewAppError(UnauthorizedError, "Unauthorized access").
		WithDetails("reason", reason).
		WithStatusCode(http.StatusUnauthorized)
}

// ForbiddenError returns forbidden error
func ForbiddenAccessError(resource string) *AppError {
	return NewAppError(ForbiddenError, fmt.Sprintf("Access to %s is forbidden", resource)).
		WithDetails("resource", resource).
		WithStatusCode(http.StatusForbidden)
}

// InvalidPaginationError returns invalid pagination error
func InvalidPaginationError(limit, offset int) *AppError {
	return NewAppError(ValidationError, "Invalid pagination parameters").
		WithDetails("limit", limit).
		WithDetails("offset", offset).
		WithStatusCode(http.StatusBadRequest)
}

// InvalidInputError returns generic invalid input error
func InvalidInputError(field, reason string) *AppError {
	return NewAppError(ValidationError, fmt.Sprintf("Invalid input for %s: %s", field, reason)).
		WithDetails("field", field).
		WithStatusCode(http.StatusBadRequest)
}

// MissingFieldError returns missing required field error
func MissingFieldError(field string) *AppError {
	return NewAppError(ValidationError, fmt.Sprintf("Missing required field: %s", field)).
		WithDetails("field", field).
		WithStatusCode(http.StatusBadRequest)
}

// InternalServerError returns internal server error
func InternalServerError(message string) *AppError {
	return NewAppError(InternalError, message).
		WithStatusCode(http.StatusInternalServerError)
}

// getStatusCode returns HTTP status code for error type
func getStatusCode(errorType ErrorType) int {
	switch errorType {
	case ValidationError:
		return http.StatusBadRequest
	case NotFoundError:
		return http.StatusNotFound
	case ConflictError:
		return http.StatusConflict
	case UnauthorizedError:
		return http.StatusUnauthorized
	case ForbiddenError:
		return http.StatusForbidden
	case DatabaseError:
		return http.StatusInternalServerError
	case InternalError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
