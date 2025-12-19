package usecase

import (
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
)

type UseCaseFactory struct {
	Patient       *PatientUseCase
	Doctor        *DoctorUseCase
	Appointment   *AppointmentUseCase
	Diagnosis     *DiagnosisUseCase
	Prescription  *PrescriptionUseCase
	MedicalTest   *MedicalTestUseCase
	TreatmentPlan *TreatmentPlanUseCase
	Disease       *DiseaseUseCase
	Medication    *MedicationUseCase
}

// NewUseCaseFactory creates all usecase instances
func NewUseCaseFactory(
	patientService *services.PatientService,
	doctorService *services.DoctorService,
	appointmentService *services.AppointmentService,
	diagnosisService *services.DiagnosisService,
	prescriptionService *services.PrescriptionService,
	medicalTestService *services.MedicalTestService,
	treatmentPlanService *services.TreatmentPlanService,
	diseaseService *services.DiseaseService,
	medicationService *services.MedicationService,
) *UseCaseFactory {
	return &UseCaseFactory{
		Patient:       NewPatientUseCase(patientService),
		Doctor:        NewDoctorUseCase(doctorService),
		Appointment:   NewAppointmentUseCase(appointmentService),
		Diagnosis:     NewDiagnosisUseCase(diagnosisService),
		Prescription:  NewPrescriptionUseCase(prescriptionService),
		MedicalTest:   NewMedicalTestUseCase(medicalTestService),
		TreatmentPlan: NewTreatmentPlanUseCase(treatmentPlanService),
		Disease:       NewDiseaseUseCase(diseaseService),
		Medication:    NewMedicationUseCase(medicationService),
	}
}
