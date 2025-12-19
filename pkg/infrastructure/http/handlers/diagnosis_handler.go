package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DiagnosisHandler struct {
	service *services.DiagnosisService
}

func NewDiagnosisHandler(service *services.DiagnosisService) *DiagnosisHandler {
	return &DiagnosisHandler{service: service}
}

// POST /diagnoses
func (h *DiagnosisHandler) CreateDiagnosis(c *gin.Context) {
	var req DiagnosisCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	diagnosisDate, err := time.Parse("2006-01-02", req.DiagnosisDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_DATE",
			Message: "diagnosis_date must be YYYY-MM-DD format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	diagnosis := &entities.Diagnosis{
		PatientID:     req.PatientID,
		DoctorID:      req.DoctorID,
		DiseaseID:     req.DiseaseID,
		DiagnosisDate: diagnosisDate,
		Severity:      req.Severity,
		Notes:         req.Notes,
	}

	err = h.service.CreateDiagnosis(c.Request.Context(), diagnosis)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Data:      diagnosis,
		Message:   "diagnosis created successfully",
		Timestamp: time.Now(),
	})
}

// GET /diagnoses/:id
func (h *DiagnosisHandler) GetDiagnosis(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "diagnosis ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	diagnosis, err := h.service.GetDiagnosis(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      diagnosis,
		Message:   "diagnosis retrieved successfully",
		Timestamp: time.Now(),
	})
}

// PUT /diagnoses/:id
func (h *DiagnosisHandler) UpdateDiagnosis(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "diagnosis ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req DiagnosisCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	diagnosisDate, err := time.Parse("2006-01-02", req.DiagnosisDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_DATE",
			Message: "diagnosis_date must be YYYY-MM-DD format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	diagnosis := &entities.Diagnosis{
		DiagnosisID:  id,
		PatientID:    req.PatientID,
		DoctorID:     req.DoctorID,
		DiseaseID:    req.DiseaseID,
		DiagnosisDate: diagnosisDate,
		Severity:     req.Severity,
		Notes:        req.Notes,
	}

	err = h.service.UpdateDiagnosis(c.Request.Context(), diagnosis)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      diagnosis,
		Message:   "diagnosis updated successfully",
		Timestamp: time.Now(),
	})
}

// DELETE /diagnoses/:id
func (h *DiagnosisHandler) DeleteDiagnosis(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "diagnosis ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	err = h.service.DeleteDiagnosis(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      nil,
		Message:   "diagnosis deleted successfully",
		Timestamp: time.Now(),
	})
}

// GET /diagnoses/patient/:patient_id
func (h *DiagnosisHandler) GetPatientDiagnoses(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	diagnoses, err := h.service.GetPatientDiagnoses(c.Request.Context(), patientID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      diagnoses,
		Message:   "patient diagnoses retrieved",
		Timestamp: time.Now(),
	})
}

// GET /diagnoses/patient/:patient_id/severe
func (h *DiagnosisHandler) GetSevereDiagnoses(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	diagnoses, err := h.service.GetSevereDiagnoses(c.Request.Context(), patientID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      diagnoses,
		Message:   "severe diagnoses retrieved",
		Timestamp: time.Now(),
	})
}

// GET /diagnoses/doctor/:doctor_id
func (h *DiagnosisHandler) GetDoctorDiagnoses(c *gin.Context) {
	doctorID, err := strconv.Atoi(c.Param("doctor_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "doctor ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	diagnoses, err := h.service.GetDoctorDiagnoses(c.Request.Context(), doctorID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      diagnoses,
		Message:   "doctor diagnoses retrieved",
		Timestamp: time.Now(),
	})
}
