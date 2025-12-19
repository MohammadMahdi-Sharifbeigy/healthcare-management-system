package utils

import "time"

// Application Constants
const (
	AppName        = "Healthcare Management System"
	AppVersion     = "1.0.0"
	DefaultTimeout = 30 * time.Second
)

// API Constants
const (
	APIVersion = "v1"
	APIBaseURL = "/api/v1"
)

// HTTP Headers
const (
	HeaderContentType      = "Content-Type"
	HeaderAuthorization   = "Authorization"
	HeaderRequestID       = "X-Request-ID"
	HeaderCorrelationID   = "X-Correlation-ID"
	ContentTypeJSON       = "application/json"
)

// Database Constants
const (
	DefaultPageSize = 10
	MaxPageSize     = 100
	DefaultOffset   = 0
)

// Pagination Limits
const (
	MinPageSize = 1
	MaxPageSize = 100
)

// Entity Field Constraints
const (
	FirstNameMinLength  = 2
	FirstNameMaxLength  = 100
	LastNameMinLength   = 2
	LastNameMaxLength   = 100
	EmailMaxLength      = 150
	PhoneMaxLength      = 20
	AddressMaxLength    = 255
	LicenseMaxLength    = 50
	SpecialtyMaxLength  = 100
	DepartmentMaxLength = 100
	DiseaseNameLength   = 150
	ICDCodeLength       = 20
	DosageMaxLength     = 10
	FrequencyMaxLength  = 50
)

// Medication Constants
const (
	FormTablet     = "tablet"
	FormCapsule    = "capsule"
	FormInjection  = "injection"
	FormLiquid     = "liquid"
)

// Severity Levels
const (
	SeverityMild     = "mild"
	SeverityModerate = "moderate"
	SeveritySevere   = "severe"
)

// Appointment Status
const (
	StatusConfirmed = "confirmed"
	StatusCompleted = "completed"
	StatusCancelled = "cancelled"
)

// Test Result Status
const (
	ResultNormal   = "normal"
	ResultAbnormal = "abnormal"
	ResultCritical = "critical"
)

// Treatment Plan Status
const (
	PlanActive    = "active"
	PlanCompleted = "completed"
	PlanCancelled = "cancelled"
)

// Gender Types
const (
	GenderMale   = "Male"
	GenderFemale = "Female"
	GenderOther  = "Other"
)

// Blood Types
const (
	BloodTypeOPositive  = "O+"
	BloodTypeONegative  = "O-"
	BloodTypeAPositive  = "A+"
	BloodTypeANegative  = "A-"
	BloodTypeBPositive  = "B+"
	BloodTypeBNegative  = "B-"
	BloodTypeABPositive = "AB+"
	BloodTypeABNegative = "AB-"
)

// Treatment Types
const (
	TreatmentPharmacotherapy = "pharmacotherapy"
	TreatmentPhysiotherapy   = "physiotherapy"
	TreatmentPsychotherapy   = "psychotherapy"
	TreatmentOccupational    = "occupational"
	TreatmentSpeech          = "speech"
	TreatmentRehabilitation  = "rehabilitation"
)

// Connection Pool Defaults
const (
	DefaultMaxOpenConns    = 25
	DefaultMaxIdleConns    = 5
	DefaultConnMaxLifetime = 5 * time.Minute
)

// Pagination Defaults
const (
	DefaultLimit  = 10
	DefaultOffset = 0
)

// Time Format Constants
const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = time.RFC3339
)

// Rate Limiting
const (
	DefaultRateLimit     = 100
	DefaultRateLimitWindow = time.Minute
)

// Cache Durations
const (
	CacheDuration5Min   = 5 * time.Minute
	CacheDuration10Min  = 10 * time.Minute
	CacheDuration1Hour  = 1 * time.Hour
	CacheDuration1Day   = 24 * time.Hour
)

// Age Constraints
const (
	MinPatientAge = 0
	MaxPatientAge = 150
)

// Session Duration (in minutes)
const (
	MinSessionDuration = 15
	MaxSessionDuration = 480 // 8 hours
)

// Valid Frequencies
var ValidFrequencies = []string{
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

// Valid Medications Forms
var ValidMedicationForms = []string{
	FormTablet,
	FormCapsule,
	FormInjection,
	FormLiquid,
}

// Valid Severities
var ValidSeverities = []string{
	SeverityMild,
	SeverityModerate,
	SeveritySevere,
}

// Valid Blood Types
var ValidBloodTypes = []string{
	BloodTypeOPositive,
	BloodTypeONegative,
	BloodTypeAPositive,
	BloodTypeANegative,
	BloodTypeBPositive,
	BloodTypeBNegative,
	BloodTypeABPositive,
	BloodTypeABNegative,
}

// Valid Genders
var ValidGenders = []string{
	GenderMale,
	GenderFemale,
	GenderOther,
}
