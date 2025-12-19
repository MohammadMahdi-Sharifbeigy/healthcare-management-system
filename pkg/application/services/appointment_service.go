package services

import (
	"context"
	"time"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/repositories"
)

type AppointmentService struct {
	appointmentRepo repositories.AppointmentRepository
	patientRepo     repositories.PatientRepository
	doctorRepo      repositories.DoctorRepository
}

func NewAppointmentService(
	appointmentRepo repositories.AppointmentRepository,
	patientRepo repositories.PatientRepository,
	doctorRepo repositories.DoctorRepository,
) *AppointmentService {
	return &AppointmentService{
		appointmentRepo: appointmentRepo,
		patientRepo:     patientRepo,
		doctorRepo:      doctorRepo,
	}
}

func (s *AppointmentService) CreateAppointment(ctx context.Context, appointment *entities.Appointment) error {
	if err := appointment.Validate(); err != nil {
		return err
	}

	// Verify patient exists
	if _, err := s.patientRepo.GetByID(ctx, appointment.PatientID); err != nil {
		return entities.ErrNotFound("patient not found")
	}

	// Verify doctor exists
	if _, err := s.doctorRepo.GetByID(ctx, appointment.DoctorID); err != nil {
		return entities.ErrNotFound("doctor not found")
	}

	// Check doctor availability
	available, err := s.appointmentRepo.CheckAvailability(ctx, appointment.DoctorID, appointment.AppointmentDate)
	if err != nil {
		return err
	}
	if !available {
		return entities.ErrConflict("doctor is not available at this time")
	}

	return s.appointmentRepo.Create(ctx, appointment)
}

func (s *AppointmentService) GetAppointmentDetails(ctx context.Context, appointmentID int) (*entities.Appointment, error) {
	if appointmentID == 0 {
		return nil, entities.ErrInvalidInput("appointment ID required")
	}

	return s.appointmentRepo.GetByID(ctx, appointmentID)
}

func (s *AppointmentService) RescheduleAppointment(ctx context.Context, appointmentID int, newDate time.Time) error {
	if appointmentID == 0 {
		return entities.ErrInvalidInput("appointment ID required")
	}

	appointment, err := s.appointmentRepo.GetByID(ctx, appointmentID)
	if err != nil {
		return err
	}

	if appointment.IsCancelled() {
		return entities.ErrConflict("cannot reschedule cancelled appointment")
	}

	// Check new availability
	available, err := s.appointmentRepo.CheckAvailability(ctx, appointment.DoctorID, newDate)
	if err != nil {
		return err
	}
	if !available {
		return entities.ErrConflict("doctor is not available at new time")
	}

	appointment.AppointmentDate = newDate
	return s.appointmentRepo.Update(ctx, appointment)
}

func (s *AppointmentService) CancelAppointment(ctx context.Context, appointmentID int) error {
	if appointmentID == 0 {
		return entities.ErrInvalidInput("appointment ID required")
	}

	appointment, err := s.appointmentRepo.GetByID(ctx, appointmentID)
	if err != nil {
		return err
	}

	if appointment.IsCancelled() {
		return entities.ErrConflict("appointment already cancelled")
	}

	appointment.Status = "cancelled"
	return s.appointmentRepo.Update(ctx, appointment)
}

func (s *AppointmentService) GetPatientAppointments(ctx context.Context, patientID int) ([]entities.Appointment, error) {
	if patientID == 0 {
		return nil, entities.ErrInvalidInput("patient ID required")
	}

	return s.appointmentRepo.GetByPatientID(ctx, patientID)
}

func (s *AppointmentService) GetDoctorAppointments(ctx context.Context, doctorID int) ([]entities.Appointment, error) {
	if doctorID == 0 {
		return nil, entities.ErrInvalidInput("doctor ID required")
	}

	return s.appointmentRepo.GetByDoctorID(ctx, doctorID)
}

func (s *AppointmentService) GetUpcomingAppointments(ctx context.Context) ([]entities.Appointment, error) {
	return s.appointmentRepo.GetByStatus(ctx, "confirmed")
}
