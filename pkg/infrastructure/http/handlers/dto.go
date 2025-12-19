package handlers

import "time"

// Request DTOs
type PatientCreateRequest struct {
	FirstName        string `json:"first_name" binding:"required,min=2,max=100"`
	LastName         string `json:"last_name" binding:"required,min=2,max=100"`
	DateOfBirth      string `json:"date_of_birth" binding:"required"` // YYYY-MM-DD
	Gender           string `json:"gender" binding:"required,oneof=Male Female Other"`
	Email            string `json:"email" binding:"required,email"`
	Phone            string `json:"phone" binding:"required,min=10"`
	Address          string `json:"address" binding:"required"`
	EmergencyContact string `json:"emergency_contact" binding:"required"`
	BloodType        string `json:"blood_type" binding:"required,oneof=O+ O- A+ A- B+ B- AB+ AB-"`
	MedicalHistory   string `json:"medical_history"`
}

type PatientUpdateRequest struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	EmergencyContact string `json:"emergency_contact"`
	MedicalHistory   string `json:"medical_history"`
}

type DoctorCreateRequest struct {
	DoctorID       int    `json:"doctor_id" binding:"required,min=10000000,max=99999999"`
	FirstName      string `json:"first_name" binding:"required,min=2,max=100"`
	LastName       string `json:"last_name" binding:"required,min=2,max=100"`
	Email          string `json:"email" binding:"required,email"`
	Phone          string `json:"phone" binding:"required,min=10"`
	LicenseNumber  string `json:"license_number" binding:"required"`
	Specialization string `json:"specialization" binding:"required"`
	Department     string `json:"department" binding:"required"`
}

type AppointmentCreateRequest struct {
	PatientID       int    `json:"patient_id" binding:"required,gt=0"`
	DoctorID        int    `json:"doctor_id" binding:"required,gt=0"`
	AppointmentDate string `json:"appointment_date" binding:"required"` // RFC3339
	Reason          string `json:"reason" binding:"required,min=5"`
	Status          string `json:"status" binding:"required,oneof=confirmed completed cancelled"`
}

type AppointmentUpdateRequest struct {
	Reason          string `json:"reason"`
	Status          string `json:"status"`
	BloodPressure   string `json:"blood_pressure"` // Format: "120/80"
	HeartRate       string `json:"heart_rate"`
	DurationMinutes int    `json:"duration_minutes"`
	Notes           string `json:"notes"`
}

type AppointmentRescheduleRequest struct {
	NewDate string `json:"new_date" binding:"required"` // RFC3339
}

type DiagnosisCreateRequest struct {
	PatientID     int    `json:"patient_id" binding:"required,gt=0"`
	DoctorID      int    `json:"doctor_id" binding:"required,gt=0"`
	DiseaseID     int    `json:"disease_id" binding:"required,gt=0"`
	DiagnosisDate string `json:"diagnosis_date" binding:"required"` // YYYY-MM-DD
	Severity      string `json:"severity" binding:"required,oneof=mild moderate severe"`
	Notes         string `json:"notes"`
}

type PrescriptionCreateRequest struct {
	PatientID      int    `json:"patient_id" binding:"required,gt=0"`
	DoctorID       int    `json:"doctor_id" binding:"required,gt=0"`
	PrescribedDate string `json:"prescribed_date" binding:"required"` // YYYY-MM-DD
}

type PrescriptionMedicationRequest struct {
	MedicationID int    `json:"medication_id" binding:"required,gt=0"`
	Dosage       string `json:"dosage" binding:"required"`
	Frequency    string `json:"frequency" binding:"required"`
	Instructions string `json:"instructions" binding:"required"`
	EndDate      string `json:"end_date" binding:"required"` // YYYY-MM-DD
}

type MedicalTestCreateRequest struct {
	PatientID       int    `json:"patient_id" binding:"required,gt=0"`
	TestCatalogID   int    `json:"test_catalog_id" binding:"required,gt=0"`
	TestDate        string `json:"test_date" binding:"required"` // YYYY-MM-DD
	ResultValue     string `json:"result_value"`
	NormalRange     string `json:"normal_range"`
	Findings        string `json:"findings"`
	ResultStatus    string `json:"result_status" binding:"required,oneof=normal abnormal critical"`
	ImageURL        string `json:"image_url"`
	InterpretedByID int    `json:"interpreted_by_id"`
}

type MedicalTestUpdateRequest struct {
	ResultValue  string `json:"result_value"`
	NormalRange  string `json:"normal_range"`
	Findings     string `json:"findings"`
	ResultStatus string `json:"result_status" binding:"oneof=normal abnormal critical"`
}

type TreatmentPlanCreateRequest struct {
	PatientID       int    `json:"patient_id" binding:"required,gt=0"`
	DoctorID        int    `json:"doctor_id" binding:"required,gt=0"`
	Diagnosis       string `json:"diagnosis" binding:"required"`
	TreatmentType   string `json:"treatment_type" binding:"required"`
	StartDate       string `json:"start_date" binding:"required"` // YYYY-MM-DD
	EndDate         string `json:"end_date"`                      // YYYY-MM-DD
	SessionDuration int    `json:"session_duration" binding:"required,gt=0"`
	Goals           string `json:"goals" binding:"required"`
}

type TreatmentPlanUpdateRequest struct {
	Diagnosis       string `json:"diagnosis"`
	TreatmentType   string `json:"treatment_type"`
	EndDate         string `json:"end_date"`
	SessionDuration int    `json:"session_duration"`
	Goals           string `json:"goals"`
	ProgressNotes   string `json:"progress_notes"`
}

type DiseaseCreateRequest struct {
	DiseaseName string `json:"disease_name" binding:"required"`
	Category    string `json:"category" binding:"required"`
	ICDCode     string `json:"icd_code" binding:"required"`
	Description string `json:"description"`
}

type MedicationCreateRequest struct {
	MedicationName string `json:"medication_name" binding:"required"`
	GenericName    string `json:"generic_name" binding:"required"`
	Form           string `json:"form" binding:"required,oneof=tablet capsule injection liquid"`
	Strength       string `json:"strength" binding:"required"`
	Manufacturer   string `json:"manufacturer" binding:"required"`
	SideEffects    string `json:"side_effects"`
}

type TestCatalogCreateRequest struct {
	TestName     string `json:"test_name" binding:"required"`
	TestCategory string `json:"test_category" binding:"required,oneof=imaging laboratory cognitive"`
	Description  string `json:"description"`
	NormalRange  string `json:"normal_range"`
}

// Response DTOs
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type SuccessResponse struct {
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	Timestamp time.Time   `json:"timestamp"`
}

type PaginationResponse struct {
	Data   interface{} `json:"data"`
	Total  int         `json:"total"`
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
}

// Status update requests
type StatusUpdateRequest struct {
	Status string `json:"status" binding:"required"`
}

type ProgressUpdateRequest struct {
	CompletedSessions int `json:"completed_sessions" binding:"required,gte=0"`
}
