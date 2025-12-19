package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type AppointmentHandler struct {
	service *services.AppointmentService
}

func NewAppointmentHandler(service *services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: service}
}

// POST /appointments
func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var req AppointmentCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	appointmentDate, err := time.Parse(time.RFC3339, req.AppointmentDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_DATE",
			Message: "appointment_date must be RFC3339 format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	appointment := &entities.Appointment{
		PatientID:       req.PatientID,
		DoctorID:        req.DoctorID,
		AppointmentDate: appointmentDate,
		Reason:          req.Reason,
		Status:          req.Status,
	}

	err = h.service.CreateAppointment(c.Request.Context(), appointment)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Data:      appointment,
		Message:   "appointment created successfully",
		Timestamp: time.Now(),
	})
}

// GET /appointments/:id
func (h *AppointmentHandler) GetAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "appointment ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	appointment, err := h.service.GetAppointmentDetails(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      appointment,
		Message:   "appointment retrieved successfully",
		Timestamp: time.Now(),
	})
}

// PUT /appointments/:id
func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "appointment ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req AppointmentUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	appointment, err := h.service.GetAppointmentDetails(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	if req.Reason != "" {
		appointment.Reason = req.Reason
	}
	if req.Status != "" {
		appointment.Status = req.Status
	}
	if req.BloodPressure != "" {
		appointment.BloodPressure = req.BloodPressure
	}
	if req.HeartRate != "" {
		appointment.HeartRate = req.HeartRate
	}
	if req.DurationMinutes > 0 {
		appointment.DurationMinutes = &req.DurationMinutes
	}
	if req.Notes != "" {
		appointment.Notes = req.Notes
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      appointment,
		Message:   "appointment updated successfully",
		Timestamp: time.Now(),
	})
}

// PATCH /appointments/:id/reschedule
func (h *AppointmentHandler) RescheduleAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "appointment ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req AppointmentRescheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	newDate, err := time.Parse(time.RFC3339, req.NewDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_DATE",
			Message: "new_date must be RFC3339 format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	err = h.service.RescheduleAppointment(c.Request.Context(), id, newDate)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      nil,
		Message:   "appointment rescheduled successfully",
		Timestamp: time.Now(),
	})
}

// DELETE /appointments/:id
func (h *AppointmentHandler) CancelAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "appointment ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	err = h.service.CancelAppointment(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      nil,
		Message:   "appointment cancelled successfully",
		Timestamp: time.Now(),
	})
}

// GET /appointments/patient/:patient_id
func (h *AppointmentHandler) GetPatientAppointments(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	appointments, err := h.service.GetPatientAppointments(c.Request.Context(), patientID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      appointments,
		Message:   "patient appointments retrieved",
		Timestamp: time.Now(),
	})
}

// GET /appointments/doctor/:doctor_id
func (h *AppointmentHandler) GetDoctorAppointments(c *gin.Context) {
	doctorID, err := strconv.Atoi(c.Param("doctor_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "doctor ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	appointments, err := h.service.GetDoctorAppointments(c.Request.Context(), doctorID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      appointments,
		Message:   "doctor appointments retrieved",
		Timestamp: time.Now(),
	})
}

// GET /appointments/upcoming
func (h *AppointmentHandler) GetUpcomingAppointments(c *gin.Context) {
	appointments, err := h.service.GetUpcomingAppointments(c.Request.Context())
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      appointments,
		Message:   "upcoming appointments retrieved",
		Timestamp: time.Now(),
	})
}
