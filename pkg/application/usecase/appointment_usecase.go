package usecase

import (
	"context"
	"time"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/dto"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type AppointmentUseCase struct {
	service *services.AppointmentService
}

func NewAppointmentUseCase(service *services.AppointmentService) *AppointmentUseCase {
	return &AppointmentUseCase{service: service}
}

// ScheduleAppointment creates new appointment with validation
func (uc *AppointmentUseCase) ScheduleAppointment(ctx context.Context, req *dto.CreateAppointmentRequest) (*dto.AppointmentResponse, error) {
	appointmentDate, _ := time.Parse(time.RFC3339, req.AppointmentDate)

	appointment := &entities.Appointment{
		PatientID:       req.PatientID,
		DoctorID:        req.DoctorID,
		AppointmentDate: appointmentDate,
		Reason:          req.Reason,
		Status:          req.Status,
	}

	if err := uc.service.CreateAppointment(ctx, appointment); err != nil {
		return nil, err
	}

	return &dto.AppointmentResponse{
		AppointmentID:   appointment.AppointmentID,
		PatientID:       appointment.PatientID,
		DoctorID:        appointment.DoctorID,
		AppointmentDate: appointment.AppointmentDate,
		Reason:          appointment.Reason,
		Status:          appointment.Status,
		CreatedAt:       appointment.CreatedAt,
		UpdatedAt:       appointment.UpdatedAt,
	}, nil
}

// GetAppointmentDetails retrieves appointment information
func (uc *AppointmentUseCase) GetAppointmentDetails(ctx context.Context, appointmentID int) (*dto.AppointmentResponse, error) {
	appointment, err := uc.service.GetAppointmentDetails(ctx, appointmentID)
	if err != nil {
		return nil, err
	}

	return &dto.AppointmentResponse{
		AppointmentID:   appointment.AppointmentID,
		PatientID:       appointment.PatientID,
		DoctorID:        appointment.DoctorID,
		AppointmentDate: appointment.AppointmentDate,
		Reason:          appointment.Reason,
		Status:          appointment.Status,
		BloodPressure:   appointment.BloodPressure,
		HeartRate:       appointment.HeartRate,
		DurationMinutes: appointment.DurationMinutes,
		Notes:           appointment.Notes,
		CreatedAt:       appointment.CreatedAt,
		UpdatedAt:       appointment.UpdatedAt,
	}, nil
}

// RescheduleAppointment changes appointment date
func (uc *AppointmentUseCase) RescheduleAppointment(ctx context.Context, appointmentID int, req *dto.AppointmentRescheduleRequest) error {
	newDate, _ := time.Parse(time.RFC3339, req.NewDate)
	return uc.service.RescheduleAppointment(ctx, appointmentID, newDate)
}

// CancelAppointment cancels an appointment
func (uc *AppointmentUseCase) CancelAppointment(ctx context.Context, appointmentID int) error {
	return uc.service.CancelAppointment(ctx, appointmentID)
}

// GetPatientAppointments retrieves all appointments for a patient
func (uc *AppointmentUseCase) GetPatientAppointments(ctx context.Context, patientID int) ([]*dto.AppointmentListResponse, error) {
	appointments, err := uc.service.GetPatientAppointments(ctx, patientID)
	if err != nil {
		return nil, err
	}

	var responses []*dto.AppointmentListResponse
	for _, a := range appointments {
		responses = append(responses, &dto.AppointmentListResponse{
			AppointmentID:   a.AppointmentID,
			AppointmentDate: a.AppointmentDate,
			Reason:          a.Reason,
			Status:          a.Status,
		})
	}

	return responses, nil
}

// GetUpcomingAppointments retrieves upcoming appointments
func (uc *AppointmentUseCase) GetUpcomingAppointments(ctx context.Context) ([]*dto.AppointmentListResponse, error) {
	appointments, err := uc.service.GetUpcomingAppointments(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*dto.AppointmentListResponse
	for _, a := range appointments {
		responses = append(responses, &dto.AppointmentListResponse{
			AppointmentID:   a.AppointmentID,
			AppointmentDate: a.AppointmentDate,
			Reason:          a.Reason,
			Status:          a.Status,
		})
	}

	return responses, nil
}
