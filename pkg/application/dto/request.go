package dto

// Patient Request DTOs
type CreatePatientRequest struct {
	FirstName        string `json:"first_name" binding:"required,min=2,max=100"`
	LastName         string `json:"last_name" binding:"required,min=2,max=100"`
	DateOfBirth      string `json:"date_of_birth" binding:"required"`
	Gender           string `json:"gender" binding:"required,oneof=Male Female Other"`
	Email            string `json:"email" binding:"required,email"`
	Phone            string `json:"phone" binding:"required,min=10"`
	Address          string `json:"address" binding:"required"`
	EmergencyContact string `json:"emergency_contact" binding:"required"`
	BloodType        string `json:"blood_type" binding:"required,oneof=O+ O- A+ A- B+ B- AB+ AB-"`
	MedicalHistory   string `json:"medical_history"`
}

type UpdatePatientRequest struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	EmergencyContact string `json:"emergency_contact"`
	MedicalHistory   string `json:"medical_history"`
}

// Doctor Request DTOs
type CreateDoctorRequest struct {
	DoctorID       int    `json:"doctor_id" binding:"required,min=10000000,max=99999999"`
	FirstName      string `json:"first_name" binding:"required,min=2,max=100"`
	LastName       string `json:"last_name" binding:"required,min=2,max=100"`
	Email          string `json:"email" binding:"required,email"`
	Phone          string `json:"phone" binding:"required,min=10"`
	LicenseNumber  string `json:"license_number" binding:"required"`
	Specialization string `json:"specialization" binding:"required"`
	Department     string `json:"department" binding:"required"`
}

type UpdateDoctorRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Specialization string `json:"specialization"`
	Department     string `json:"department"`
}

// Appointment Request DTOs
type CreateAppointmentRequest struct {
	PatientID       int    `json:"patient_id" binding:"required,gt=0"`
	DoctorID        int    `json:"doctor_id" binding:"required,gt=0"`
	AppointmentDate string `json:"appointment_date" binding:"required"`
	Reason          string `json:"reason" binding:"required,min=5"`
	Status          string `json:"status" binding:"required,oneof=confirmed completed cancelled"`
}

type UpdateAppointmentRequest struct {
	Reason          string `json:"reason"`
	Status          string `json:"status"`
	BloodPressure   string `json:"blood_pressure"`
	HeartRate       string `json:"heart_rate"`
	DurationMinutes int    `json:"duration_minutes"`
	Notes           string `json:"notes"`
}

type RescheduleAppointmentRequest struct {
	NewDate string `json:"new_date" binding:"required"`
}

// Diagnosis Request DTOs
type CreateDiagnosisRequest struct {
	PatientID     int    `json:"patient_id" binding:"required,gt=0"`
	DoctorID      int    `json:"doctor_id" binding:"required,gt=0"`
	DiseaseID     int    `json:"disease_id" binding:"required,gt=0"`
	DiagnosisDate string `json:"diagnosis_date" binding:"required"`
	Severity      string `json:"severity" binding:"required,oneof=mild moderate severe"`
	Notes         string `json:"notes"`
}

type UpdateDiagnosisRequest struct {
	Severity string `json:"severity"`
	Notes    string `json:"notes"`
}

// Prescription Request DTOs
type CreatePrescriptionRequest struct {
	PatientID      int    `json:"patient_id" binding:"required,gt=0"`
	DoctorID       int    `json:"doctor_id" binding:"required,gt=0"`
	PrescribedDate string `json:"prescribed_date" binding:"required"`
}

type AddMedicationRequest struct {
	MedicationID int    `json:"medication_id" binding:"required,gt=0"`
	Dosage       string `json:"dosage" binding:"required"`
	Frequency    string `json:"frequency" binding:"required"`
	Instructions string `json:"instructions" binding:"required"`
	EndDate      string `json:"end_date" binding:"required"`
}

type UpdatePrescriptionRequest struct {
	PatientID      int    `json:"patient_id"`
	DoctorID       int    `json:"doctor_id"`
	PrescribedDate string `json:"prescribed_date"`
}

// Medical Test Request DTOs
type CreateMedicalTestRequest struct {
	PatientID       int    `json:"patient_id" binding:"required,gt=0"`
	TestCatalogID   int    `json:"test_catalog_id" binding:"required,gt=0"`
	TestDate        string `json:"test_date" binding:"required"`
	ResultValue     string `json:"result_value"`
	NormalRange     string `json:"normal_range"`
	Findings        string `json:"findings"`
	ResultStatus    string `json:"result_status" binding:"required,oneof=normal abnormal critical"`
	ImageURL        string `json:"image_url"`
	InterpretedByID int    `json:"interpreted_by_id"`
}

type UpdateMedicalTestRequest struct {
	ResultValue  string `json:"result_value"`
	NormalRange  string `json:"normal_range"`
	Findings     string `json:"findings"`
	ResultStatus string `json:"result_status" binding:"oneof=normal abnormal critical"`
}

type InterpretTestRequest struct {
	DoctorID int `json:"doctor_id" binding:"required,gt=0"`
}

// Treatment Plan Request DTOs
type CreateTreatmentPlanRequest struct {
	PatientID       int    `json:"patient_id" binding:"required,gt=0"`
	DoctorID        int    `json:"doctor_id" binding:"required,gt=0"`
	Diagnosis       string `json:"diagnosis" binding:"required"`
	TreatmentType   string `json:"treatment_type" binding:"required"`
	StartDate       string `json:"start_date" binding:"required"`
	EndDate         string `json:"end_date"`
	SessionDuration int    `json:"session_duration" binding:"required,gt=0"`
	Goals           string `json:"goals" binding:"required"`
}

type UpdateTreatmentPlanRequest struct {
	Diagnosis       string `json:"diagnosis"`
	TreatmentType   string `json:"treatment_type"`
	EndDate         string `json:"end_date"`
	SessionDuration int    `json:"session_duration"`
	Goals           string `json:"goals"`
	ProgressNotes   string `json:"progress_notes"`
}

type UpdateTreatmentPlanStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=active completed cancelled"`
}

// Disease Request DTOs
type CreateDiseaseRequest struct {
	DiseaseName string `json:"disease_name" binding:"required"`
	Category    string `json:"category" binding:"required"`
	ICDCode     string `json:"icd_code" binding:"required"`
	Description string `json:"description"`
}

type UpdateDiseaseRequest struct {
	DiseaseName string `json:"disease_name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

// Medication Request DTOs
type CreateMedicationRequest struct {
	MedicationName string `json:"medication_name" binding:"required"`
	GenericName    string `json:"generic_name" binding:"required"`
	Form           string `json:"form" binding:"required,oneof=tablet capsule injection liquid"`
	Strength       string `json:"strength" binding:"required"`
	Manufacturer   string `json:"manufacturer" binding:"required"`
	SideEffects    string `json:"side_effects"`
}

type UpdateMedicationRequest struct {
	MedicationName string `json:"medication_name"`
	GenericName    string `json:"generic_name"`
	Form           string `json:"form"`
	Strength       string `json:"strength"`
	Manufacturer   string `json:"manufacturer"`
	SideEffects    string `json:"side_effects"`
}

// Test Catalog Request DTOs
type CreateTestCatalogRequest struct {
	TestName     string `json:"test_name" binding:"required"`
	TestCategory string `json:"test_category" binding:"required,oneof=imaging laboratory cognitive"`
	Description  string `json:"description"`
	NormalRange  string `json:"normal_range"`
}

type UpdateTestCatalogRequest struct {
	TestName     string `json:"test_name"`
	TestCategory string `json:"test_category"`
	Description  string `json:"description"`
	NormalRange  string `json:"normal_range"`
}

// Search/Filter Request DTOs
type SearchRequest struct {
	Query  string `json:"query" binding:"required"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type PaginationRequest struct {
	Limit  int `json:"limit" binding:"required,gt=0,max=100"`
	Offset int `json:"offset" binding:"required,gte=0"`
}

type DateRangeRequest struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

// Status Update DTOs
type StatusUpdateRequest struct {
	Status string `json:"status" binding:"required"`
}

type ProgressUpdateRequest struct {
	CompletedSessions int `json:"completed_sessions" binding:"required,gte=0"`
}

// Query Parameter DTOs
type GetPatientAppointmentsQuery struct {
	Status string `form:"status"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type GetActivePrescriptionsQuery struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type GetExpiringPrescriptionsQuery struct {
	Days   int `form:"days"`
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type GetPatientTestsQuery struct {
	Status string `form:"status"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type GetPatientPlansQuery struct {
	Status string `form:"status"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type DoctorFilterQuery struct {
	Specialization string `form:"specialization"`
	Department     string `form:"department"`
	Limit          int    `form:"limit"`
	Offset         int    `form:"offset"`
}

type DiseaseFilterQuery struct {
	Category string `form:"category"`
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
}

type MedicationFilterQuery struct {
	Form   string `form:"form"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}
