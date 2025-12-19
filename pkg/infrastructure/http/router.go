package http

import (
	"github.com/gin-gonic/gin"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/infrastructure/http/handlers"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/infrastructure/middleware"
)

type Handlers struct {
	Patient       *handlers.PatientHandler
	Doctor        *handlers.DoctorHandler
	Appointment   *handlers.AppointmentHandler
	Diagnosis     *handlers.DiagnosisHandler
	Prescription  *handlers.PrescriptionHandler
	MedicalTest   *handlers.MedicalTestHandler
	TreatmentPlan *handlers.TreatmentPlanHandler
	Disease       *handlers.DiseaseHandler
	Medication    *handlers.MedicationHandler
}

// NewHandlers creates new handlers with all dependencies
func NewHandlers(
	patientService *services.PatientService,
	doctorService *services.DoctorService,
	appointmentService *services.AppointmentService,
	diagnosisService *services.DiagnosisService,
	prescriptionService *services.PrescriptionService,
	medicalTestService *services.MedicalTestService,
	treatmentPlanService *services.TreatmentPlanService,
	diseaseService *services.DiseaseService,
	medicationService *services.MedicationService,
) *Handlers {
	return &Handlers{
		Patient:       handlers.NewPatientHandler(patientService),
		Doctor:        handlers.NewDoctorHandler(doctorService),
		Appointment:   handlers.NewAppointmentHandler(appointmentService),
		Diagnosis:     handlers.NewDiagnosisHandler(diagnosisService),
		Prescription:  handlers.NewPrescriptionHandler(prescriptionService),
		MedicalTest:   handlers.NewMedicalTestHandler(medicalTestService),
		TreatmentPlan: handlers.NewTreatmentPlanHandler(treatmentPlanService),
		Disease:       handlers.NewDiseaseHandler(diseaseService),
		Medication:    handlers.NewMedicationHandler(medicationService),
	}
}

// SetupRoutes configures all routes with middleware
func SetupRoutes(router *gin.Engine, h *Handlers, config *RouterConfig) {
	// Global middleware
	router.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestIDMiddleware(),
		middleware.LoggerMiddleware(),
		middleware.CORSMiddleware(config.AllowedOrigins),
		middleware.ContentTypeMiddleware(),
		middleware.MaxRequestSizeMiddleware(config.MaxRequestSize),
		middleware.ErrorHandlerMiddleware(),
	)

	// Rate limiting (optional)
	if config.RateLimitEnabled {
		router.Use(middleware.RateLimitMiddleware(
			config.RateLimitRequests,
			config.RateLimitWindow,
		))
	}

	// Health check endpoints (no auth required)
	middleware.InitHealth()
	router.GET("/health", middleware.HealthCheckHandler())
	router.GET("/health/live", middleware.LivenessProbeHandler())
	router.GET("/health/ready", middleware.ReadinessProbeHandler())

	// API v1 routes
	v1 := router.Group("/api/v1")

	// Public endpoints (no auth required)
	setupPublicRoutes(v1, h)

	// Protected endpoints (auth required)
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())
	setupProtectedRoutes(protected, h)

	// Admin endpoints (auth + admin role required)
	admin := v1.Group("")
	admin.Use(
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("admin"),
	)
	setupAdminRoutes(admin, h)

	// 404 handler
	router.NoRoute(middleware.NotFoundHandler())
}

// setupPublicRoutes sets up routes that don't require authentication
func setupPublicRoutes(group *gin.RouterGroup, h *Handlers) {
	// Disease catalog (read-only)
	group.GET("/diseases", h.Disease.GetDiseases)
	group.GET("/diseases/:id", h.Disease.GetDisease)
	group.GET("/diseases/search", h.Disease.SearchDiseases)
	group.GET("/diseases/icd/:code", h.Disease.GetByICDCode)
	group.GET("/diseases/category/:cat", h.Disease.GetByCategory)

	// Medication catalog (read-only)
	group.GET("/medications", h.Medication.GetMedications)
	group.GET("/medications/:id", h.Medication.GetMedication)
	group.GET("/medications/search", h.Medication.SearchMedications)
	group.GET("/medications/generic/:name", h.Medication.GetByGenericName)
	group.GET("/medications/form/:form", h.Medication.GetByForm)

	// Test catalog (read-only)
	group.GET("/test-catalog", h.MedicalTest.GetTestCatalog)

	// Doctor lookup (read-only)
	group.GET("/doctors", h.Doctor.GetDoctors)
	group.GET("/doctors/:id", h.Doctor.GetDoctor)
	group.GET("/doctors/specialization/:spec", h.Doctor.GetDoctorsBySpecialization)
	group.GET("/doctors/department/:dept", h.Doctor.GetDoctorsByDepartment)
}

// setupProtectedRoutes sets up routes that require authentication
func setupProtectedRoutes(group *gin.RouterGroup, h *Handlers) {
	// Patient routes
	group.POST("/patients", h.Patient.CreatePatient)
	group.GET("/patients", h.Patient.GetPatients)
	group.GET("/patients/search", h.Patient.SearchPatients)
	group.GET("/patients/:id", h.Patient.GetPatient)
	group.PUT("/patients/:id", h.Patient.UpdatePatient)
	group.DELETE("/patients/:id", h.Patient.DeletePatient)
	group.GET("/patients/:id/appointments", h.Patient.GetPatientAppointments)

	// Appointment routes
	group.POST("/appointments", h.Appointment.CreateAppointment)
	group.GET("/appointments", h.Appointment.GetAppointments)
	group.GET("/appointments/:id", h.Appointment.GetAppointment)
	group.PUT("/appointments/:id", h.Appointment.UpdateAppointment)
	group.PATCH("/appointments/:id/reschedule", h.Appointment.RescheduleAppointment)
	group.DELETE("/appointments/:id", h.Appointment.CancelAppointment)
	group.GET("/appointments/patient/:patient_id", h.Appointment.GetPatientAppointments)
	group.GET("/appointments/doctor/:doctor_id", h.Appointment.GetDoctorAppointments)
	group.GET("/appointments/upcoming", h.Appointment.GetUpcomingAppointments)

	// Diagnosis routes
	group.POST("/diagnoses", h.Diagnosis.CreateDiagnosis)
	group.GET("/diagnoses/:id", h.Diagnosis.GetDiagnosis)
	group.PUT("/diagnoses/:id", h.Diagnosis.UpdateDiagnosis)
	group.DELETE("/diagnoses/:id", h.Diagnosis.DeleteDiagnosis)
	group.GET("/diagnoses/patient/:patient_id", h.Diagnosis.GetPatientDiagnoses)
	group.GET("/diagnoses/patient/:patient_id/severe", h.Diagnosis.GetSevereDiagnoses)
	group.GET("/diagnoses/doctor/:doctor_id", h.Diagnosis.GetDoctorDiagnoses)

	// Prescription routes
	group.POST("/prescriptions", h.Prescription.CreatePrescription)
	group.GET("/prescriptions/:id", h.Prescription.GetPrescription)
	group.POST("/prescriptions/:id/medications", h.Prescription.AddMedication)
	group.DELETE("/prescriptions/:id/medications/:medication_id", h.Prescription.RemoveMedication)
	group.GET("/prescriptions/patient/:patient_id/active", h.Prescription.GetActivePrescriptions)
	group.GET("/prescriptions/patient/:patient_id/expiring", h.Prescription.GetExpiringPrescriptions)

	// Medical test routes
	group.POST("/medical-tests", h.MedicalTest.CreateTest)
	group.GET("/medical-tests/:id", h.MedicalTest.GetTest)
	group.PUT("/medical-tests/:id", h.MedicalTest.UpdateTest)
	group.PATCH("/medical-tests/:id/interpret", h.MedicalTest.InterpretTest)
	group.GET("/medical-tests/patient/:patient_id", h.MedicalTest.GetPatientTests)
	group.GET("/medical-tests/abnormal", h.MedicalTest.GetAbnormalResults)
	group.GET("/medical-tests/critical", h.MedicalTest.GetCriticalResults)

	// Treatment plan routes
	group.POST("/treatment-plans", h.TreatmentPlan.CreatePlan)
	group.GET("/treatment-plans/:id", h.TreatmentPlan.GetPlan)
	group.PUT("/treatment-plans/:id", h.TreatmentPlan.UpdatePlan)
	group.PATCH("/treatment-plans/:id/status", h.TreatmentPlan.UpdateStatus)
	group.GET("/treatment-plans/patient/:patient_id", h.TreatmentPlan.GetPatientPlans)
	group.GET("/treatment-plans/patient/:patient_id/active", h.TreatmentPlan.GetActivePlans)
	group.GET("/treatment-plans/:id/progress", h.TreatmentPlan.GetProgress)
}

// setupAdminRoutes sets up routes that require admin role
func setupAdminRoutes(group *gin.RouterGroup, h *Handlers) {
	// Doctor management (admin only)
	group.POST("/doctors", h.Doctor.CreateDoctor)
	group.PUT("/doctors/:id", h.Doctor.UpdateDoctor)

	// Disease management (admin only)
	group.POST("/diseases", h.Disease.CreateDisease)

	// Medication management (admin only)
	group.POST("/medications", h.Medication.CreateMedication)
}
