package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type PrescriptionHandler struct {
	service *services.PrescriptionService
}

func NewPrescriptionHandler(service *services.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{service: service}
}

// POST /prescriptions
func (h *PrescriptionHandler) CreatePrescription(c *gin.Context) {
	var req PrescriptionCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	prescribedDate, err := time.Parse("2006-01-02", req.PrescribedDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_DATE",
			Message: "prescribed_date must be YYYY-MM-DD format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	prescription := &entities.Prescription{
		PatientID:      req.PatientID,
		DoctorID:       req.DoctorID,
		PrescribedDate: prescribedDate,
	}

	err = h.service.CreatePrescription(c.Request.Context(), prescription)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Data:      prescription,
		Message:   "prescription created successfully",
		Timestamp: time.Now(),
	})
}

// GET /prescriptions/:id
func (h *PrescriptionHandler) GetPrescription(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "prescription ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	prescription, err := h.service.GetPrescription(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      prescription,
		Message:   "prescription retrieved successfully",
		Timestamp: time.Now(),
	})
}

// POST /prescriptions/:id/medications
func (h *PrescriptionHandler) AddMedication(c *gin.Context) {
	prescriptionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "prescription ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req PrescriptionMedicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_DATE",
			Message: "end_date must be YYYY-MM-DD format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	pm := &entities.PrescriptionMedication{
		PrescriptionID: prescriptionID,
		MedicationID:   req.MedicationID,
		Dosage:         req.Dosage,
		Frequency:      req.Frequency,
		Instructions:   req.Instructions,
		EndDate:        endDate,
	}

	err = h.service.AddMedicationToPrescription(c.Request.Context(), pm)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Data:      pm,
		Message:   "medication added to prescription",
		Timestamp: time.Now(),
	})
}

// DELETE /prescriptions/:id/medications/:medication_id
func (h *PrescriptionHandler) RemoveMedication(c *gin.Context) {
	prescriptionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "prescription ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	medicationID, err := strconv.Atoi(c.Param("medication_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "medication ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	err = h.service.RemoveMedicationFromPrescription(c.Request.Context(), prescriptionID, medicationID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      nil,
		Message:   "medication removed from prescription",
		Timestamp: time.Now(),
	})
}

// GET /prescriptions/patient/:patient_id/active
func (h *PrescriptionHandler) GetActivePrescriptions(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	prescriptions, err := h.service.GetActivePrescriptions(c.Request.Context(), patientID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      prescriptions,
		Message:   "active prescriptions retrieved",
		Timestamp: time.Now(),
	})
}

// GET /prescriptions/patient/:patient_id/expiring
func (h *PrescriptionHandler) GetExpiringPrescriptions(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	daysAhead := 7
	if d := c.Query("days"); d != "" {
		if val, err := strconv.Atoi(d); err == nil && val > 0 {
			daysAhead = val
		}
	}

	prescriptions, err := h.service.GetExpiringPrescriptions(c.Request.Context(), patientID, daysAhead)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      prescriptions,
		Message:   "expiring prescriptions retrieved",
		Timestamp: time.Now(),
	})
}
