package repositories

import (
	"context"
	"time"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type AppointmentRepository interface {
	// CRUD operations
	Create(ctx context.Context, appointment *entities.Appointment) error
	GetByID(ctx context.Context, id int) (*entities.Appointment, error)
	Update(ctx context.Context, appointment *entities.Appointment) error
	Delete(ctx context.Context, id int) error

	// Search operations
	GetAll(ctx context.Context, limit, offset int) ([]entities.Appointment, error)
	GetByPatientID(ctx context.Context, patientID int) ([]entities.Appointment, error)
	GetByDoctorID(ctx context.Context, doctorID int) ([]entities.Appointment, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]entities.Appointment, error)
	GetByStatus(ctx context.Context, status string) ([]entities.Appointment, error)
	CheckAvailability(ctx context.Context, doctorID int, appointmentDate time.Time) (bool, error)
}
