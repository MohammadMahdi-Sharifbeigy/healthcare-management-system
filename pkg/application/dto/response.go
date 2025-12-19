package dto

import "time"

// Base Response Types
type SuccessResponse struct {
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	Timestamp time.Time   `json:"timestamp"`
}

type ErrorResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	Code      int    `json:"code"`
	RequestID string `json:"request_id,omitempty"`
}

type PaginatedResponse struct {
	Data   interface{} `json:"data"`
	Total  int         `json:"total"`
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    int64     `json:"uptime_seconds"`
	Version   string    `json:"version"`
}

// Patient Response DTOs
type PatientResponse struct {
	PatientID        int       `json:"patient_id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	Gender           string    `json:"gender"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	Address          string    `json:"address"`
	EmergencyContact string    `json:"emergency_contact"`
	BloodType        string    `json:"blood_type"`
	MedicalHistory   string    `json:"medical_history"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type PatientListResponse struct {
	PatientID   int       `json:"patient_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	BloodType   string    `json:"blood_type"`
	LastVisit   *time.Time `json:"last_visit"`
}

// Doctor Response DTOs
type DoctorResponse struct {
	DoctorID       int    `json:"doctor_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	LicenseNumber  string `json:"license_number"`
	Specialization string `json:"specialization"`
	Department     string `json:"department"`
}

type DoctorListResponse struct {
	DoctorID       int    `json:"doctor_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Specialization string `json:"specialization"`
	Department     string `json:"department"`
	Email          string `json:"email"`
}

// Appointment Response DTOs
type AppointmentResponse struct {
	AppointmentID   int        `json:"appointment_id"`
	PatientID       int        `json:"patient_id"`
	DoctorID        int        `json:"doctor_id"`
	AppointmentDate time.Time  `json:"appointment_date"`
	Reason          string     `json:"reason"`
	Status          string     `json:"status"`
	BloodPressure   string     `json:"blood_pressure"`
	HeartRate       string     `json:"heart_rate"`
	DurationMinutes *int       `json:"duration_minutes"`
	Notes           string     `json:"notes"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type AppointmentListResponse struct {
	AppointmentID   int       `json:"appointment_id"`
	PatientName     string    `json:"patient_name"`
	DoctorName      string    `json:"doctor_name"`
	AppointmentDate time.Time `json:"appointment_date"`
	Reason          string    `json:"reason"`
	Status          string    `json:"status"`
}

// Diagnosis Response DTOs
type DiagnosisResponse struct {
	DiagnosisID   int       `json:"diagnosis_id"`
	PatientID     int       `json:"patient_id"`
	DoctorID      int       `json:"doctor_id"`
	DiseaseID     int       `json:"disease_id"`
	DiseaseName   string    `json:"disease_name"`
	DiagnosisDate time.Time `json:"diagnosis_date"`
	Severity      string    `json:"severity"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type DiagnosisListResponse struct {
	DiagnosisID   int       `json:"diagnosis_id"`
	DiseaseName   string    `json:"disease_name"`
	DiagnosisDate time.Time `json:"diagnosis_date"`
	Severity      string    `json:"severity"`
	DoctorName    string    `json:"doctor_name"`
}

// Prescription Response DTOs
type PrescriptionResponse struct {
	PrescriptionID int       `json:"prescription_id"`
	PatientID      int       `json:"patient_id"`
	DoctorID       int       `json:"doctor_id"`
	PrescribedDate time.Time `json:"prescribed_date"`
	Medications    []MedicationInPrescription `json:"medications"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type MedicationInPrescription struct {
	MedicationID              int       `json:"medication_id"`
	MedicationName            string    `json:"medication_name"`
	GenericName               string    `json:"generic_name"`
	Dosage                    string    `json:"dosage"`
	Frequency                 string    `json:"frequency"`
	Instructions              string    `json:"instructions"`
	EndDate                   time.Time `json:"end_date"`
	DaysRemaining             int       `json:"days_remaining"`
	Status                    string    `json:"status"`
}

type PrescriptionListResponse struct {
	PrescriptionID int       `json:"prescription_id"`
	PatientName    string    `json:"patient_name"`
	DoctorName     string    `json:"doctor_name"`
	PrescribedDate time.Time `json:"prescribed_date"`
	MedicationCount int      `json:"medication_count"`
}

// Medical Test Response DTOs
type MedicalTestResponse struct {
	TestID          int        `json:"test_id"`
	PatientID       int        `json:"patient_id"`
	TestCatalogID   int        `json:"test_catalog_id"`
	TestName        string     `json:"test_name"`
	TestDate        time.Time  `json:"test_date"`
	ResultValue     string     `json:"result_value"`
	NormalRange     string     `json:"normal_range"`
	Findings        string     `json:"findings"`
	ResultStatus    string     `json:"result_status"`
	ImageURL        string     `json:"image_url"`
	InterpretedBy   *string    `json:"interpreted_by"`
	InterpretedDate *time.Time `json:"interpreted_date"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type MedicalTestListResponse struct {
	TestID       int       `json:"test_id"`
	TestName     string    `json:"test_name"`
	TestDate     time.Time `json:"test_date"`
	ResultStatus string    `json:"result_status"`
	PatientName  string    `json:"patient_name"`
}

// Treatment Plan Response DTOs
type TreatmentPlanResponse struct {
	PlanID          int        `json:"plan_id"`
	PatientID       int        `json:"patient_id"`
	DoctorID        int        `json:"doctor_id"`
	Diagnosis       string     `json:"diagnosis"`
	TreatmentType   string     `json:"treatment_type"`
	StartDate       time.Time  `json:"start_date"`
	EndDate         *time.Time `json:"end_date"`
	SessionDuration int        `json:"session_duration"`
	Status          string     `json:"status"`
	Goals           string     `json:"goals"`
	ProgressNotes   string     `json:"progress_notes"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type TreatmentPlanListResponse struct {
	PlanID        int       `json:"plan_id"`
	PatientName   string    `json:"patient_name"`
	TreatmentType string    `json:"treatment_type"`
	StartDate     time.Time `json:"start_date"`
	Status        string    `json:"status"`
	Progress      float64   `json:"progress_percentage"`
}

// Disease Response DTOs
type DiseaseResponse struct {
	DiseaseID   int    `json:"disease_id"`
	DiseaseName string `json:"disease_name"`
	Category    string `json:"category"`
	ICDCode     string `json:"icd_code"`
	Description string `json:"description"`
}

type DiseaseListResponse struct {
	DiseaseID   int    `json:"disease_id"`
	DiseaseName string `json:"disease_name"`
	Category    string `json:"category"`
	ICDCode     string `json:"icd_code"`
}

// Medication Response DTOs
type MedicationResponse struct {
	MedicationID   int    `json:"medication_id"`
	MedicationName string `json:"medication_name"`
	GenericName    string `json:"generic_name"`
	Form           string `json:"form"`
	Strength       string `json:"strength"`
	Manufacturer   string `json:"manufacturer"`
	SideEffects    string `json:"side_effects"`
}

type MedicationListResponse struct {
	MedicationID   int    `json:"medication_id"`
	MedicationName string `json:"medication_name"`
	GenericName    string `json:"generic_name"`
	Form           string `json:"form"`
	Strength       string `json:"strength"`
}

// Test Catalog Response DTOs
type TestCatalogResponse struct {
	TestCatalogID int    `json:"test_catalog_id"`
	TestName      string `json:"test_name"`
	TestCategory  string `json:"test_category"`
	Description   string `json:"description"`
	NormalRange   string `json:"normal_range"`
}

type TestCatalogListResponse struct {
	TestCatalogID int    `json:"test_catalog_id"`
	TestName      string `json:"test_name"`
	TestCategory  string `json:"test_category"`
}

// Statistics Response DTOs
type StatisticsResponse struct {
	TotalPatients      int `json:"total_patients"`
	TotalDoctors       int `json:"total_doctors"`
	TotalAppointments  int `json:"total_appointments"`
	TotalDiagnoses     int `json:"total_diagnoses"`
	ActivePrescriptions int `json:"active_prescriptions"`
	CompletedTests     int `json:"completed_tests"`
}

type DoctorWorkloadResponse struct {
	DoctorID            int       `json:"doctor_id"`
	DoctorName          string    `json:"doctor_name"`
	TotalPatients       int       `json:"total_patients"`
	TotalAppointments   int       `json:"total_appointments"`
	AppointmentsThisWeek int      `json:"appointments_this_week"`
	AverageDuration     float64   `json:"average_duration"`
	TotalDiagnoses      int       `json:"total_diagnoses"`
	TotalPrescriptions  int       `json:"total_prescriptions"`
	ActiveTreatmentPlans int      `json:"active_treatment_plans"`
	LastAppointment     *time.Time `json:"last_appointment"`
}

type PatientMedicalHistoryResponse struct {
	PatientID       int                      `json:"patient_id"`
	PatientName     string                   `json:"patient_name"`
	DateOfBirth     time.Time                `json:"date_of_birth"`
	Diagnoses       []DiagnosisResponse      `json:"diagnoses"`
	Appointments    []AppointmentResponse    `json:"appointments"`
	Prescriptions   []PrescriptionResponse   `json:"prescriptions"`
	MedicalTests    []MedicalTestResponse    `json:"medical_tests"`
	TreatmentPlans  []TreatmentPlanResponse  `json:"treatment_plans"`
}

// Batch Operation Response DTOs
type BatchCreateResponse struct {
	SuccessCount int           `json:"success_count"`
	FailureCount int           `json:"failure_count"`
	Errors       []BatchError  `json:"errors,omitempty"`
}

type BatchError struct {
	Index   int    `json:"index"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Validation Error Response
type ValidationErrorResponse struct {
	Error  string                   `json:"error"`
	Code   int                      `json:"code"`
	Fields map[string]FieldError   `json:"fields"`
}

type FieldError struct {
	Message string `json:"message"`
	Value   string `json:"value,omitempty"`
}
