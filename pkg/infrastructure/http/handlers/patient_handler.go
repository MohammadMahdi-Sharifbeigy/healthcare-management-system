package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type PatientHandler struct {
	service *services.PatientService
}

func NewPatientHandler(service *services.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}

// POST /patients
func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var req PatientCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_DATE",
			Message: "date_of_birth must be YYYY-MM-DD format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	patient := &entities.Patient{
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		DateOfBirth:      dob,
		Gender:           req.Gender,
		Email:            req.Email,
		Phone:            req.Phone,
		Address:          req.Address,
		EmergencyContact: req.EmergencyContact,
		BloodType:        req.BloodType,
		MedicalHistory:   req.MedicalHistory,
	}

	err = h.service.RegisterPatient(c.Request.Context(), patient)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Data:      patient,
		Message:   "patient registered successfully",
		Timestamp: time.Now(),
	})
}

// GET /patients/:id
func (h *PatientHandler) GetPatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	patient, err := h.service.GetPatientProfile(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      patient,
		Message:   "patient retrieved successfully",
		Timestamp: time.Now(),
	})
}

// PUT /patients/:id
func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req PatientUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	patient, err := h.service.GetPatientProfile(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	if req.FirstName != "" {
		patient.FirstName = req.FirstName
	}
	if req.LastName != "" {
		patient.LastName = req.LastName
	}
	if req.Phone != "" {
		patient.Phone = req.Phone
	}
	if req.Address != "" {
		patient.Address = req.Address
	}
	if req.EmergencyContact != "" {
		patient.EmergencyContact = req.EmergencyContact
	}
	if req.MedicalHistory != "" {
		patient.MedicalHistory = req.MedicalHistory
	}

	err = h.service.UpdatePatientInfo(c.Request.Context(), patient)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      patient,
		Message:   "patient updated successfully",
		Timestamp: time.Now(),
	})
}

// DELETE /patients/:id
func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	err = h.service.DeletePatient(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      nil,
		Message:   "patient deleted successfully",
		Timestamp: time.Now(),
	})
}

// GET /patients
func (h *PatientHandler) GetPatients(c *gin.Context) {
	limit := 20
	offset := 0

	if l := c.Query("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		}
	}

	if o := c.Query("offset"); o != "" {
		if val, err := strconv.Atoi(o); err == nil && val >= 0 {
			offset = val
		}
	}

	patients, err := h.service.GetPatientsByPage(c.Request.Context(), limit, offset)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, PaginationResponse{
		Data:   patients,
		Limit:  limit,
		Offset: offset,
	})
}

// GET /patients/search?query=name
func (h *PatientHandler) SearchPatients(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "MISSING_QUERY",
			Message: "query parameter required",
			Code:    http.StatusBadRequest,
		})
		return
	}

	patients, err := h.service.SearchPatients(c.Request.Context(), query)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      patients,
		Message:   "search completed",
		Timestamp: time.Now(),
	})
}

// GET /patients/:id/appointments
func (h *PatientHandler) GetPatientAppointments(c *gin.Context) {
	// Will be implemented in appointment handler with full context
	c.JSON(http.StatusOK, SuccessResponse{
		Data:      []interface{}{},
		Message:   "implementation in appointment handler",
		Timestamp: time.Now(),
	})
}
