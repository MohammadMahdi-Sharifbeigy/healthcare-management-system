package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type MedicalTestHandler struct {
	service *services.MedicalTestService
}

func NewMedicalTestHandler(service *services.MedicalTestService) *MedicalTestHandler {
	return &MedicalTestHandler{service: service}
}

// POST /medical-tests
func (h *MedicalTestHandler) CreateTest(c *gin.Context) {
	var req MedicalTestCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	testDate, err := time.Parse("2006-01-02", req.TestDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_DATE",
			Message: "test_date must be YYYY-MM-DD format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var interpretedBy *int
	if req.InterpretedByID > 0 {
		interpretedBy = &req.InterpretedByID
	}

	test := &entities.MedicalTest{
		PatientID:       req.PatientID,
		TestCatalogID:   req.TestCatalogID,
		TestDate:        testDate,
		ResultValue:     req.ResultValue,
		NormalRange:     req.NormalRange,
		Findings:        req.Findings,
		ResultStatus:    req.ResultStatus,
		ImageURL:        req.ImageURL,
		InterpretedBy:   interpretedBy,
	}

	err = h.service.CreateTest(c.Request.Context(), test)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Data:      test,
		Message:   "medical test created successfully",
		Timestamp: time.Now(),
	})
}

// GET /medical-tests/:id
func (h *MedicalTestHandler) GetTest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "test ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	test, err := h.service.GetTest(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      test,
		Message:   "medical test retrieved successfully",
		Timestamp: time.Now(),
	})
}

// PUT /medical-tests/:id
func (h *MedicalTestHandler) UpdateTest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "test ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req MedicalTestUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	test, err := h.service.GetTest(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	if req.ResultValue != "" {
		test.ResultValue = req.ResultValue
	}
	if req.NormalRange != "" {
		test.NormalRange = req.NormalRange
	}
	if req.Findings != "" {
		test.Findings = req.Findings
	}
	if req.ResultStatus != "" {
		test.ResultStatus = req.ResultStatus
	}

	err = h.service.UpdateTestResults(c.Request.Context(), test)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      test,
		Message:   "medical test updated successfully",
		Timestamp: time.Now(),
	})
}

// PATCH /medical-tests/:id/interpret
func (h *MedicalTestHandler) InterpretTest(c *gin.Context) {
	testID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "test ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req struct {
		DoctorID int `json:"doctor_id" binding:"required,gt=0"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	err = h.service.InterpretTest(c.Request.Context(), testID, req.DoctorID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      nil,
		Message:   "test interpretation assigned",
		Timestamp: time.Now(),
	})
}

// GET /medical-tests/patient/:patient_id
func (h *MedicalTestHandler) GetPatientTests(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	tests, err := h.service.GetPatientTests(c.Request.Context(), patientID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      tests,
		Message:   "patient tests retrieved",
		Timestamp: time.Now(),
	})
}

// GET /medical-tests/abnormal
func (h *MedicalTestHandler) GetAbnormalResults(c *gin.Context) {
	tests, err := h.service.GetAbnormalResults(c.Request.Context())
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      tests,
		Message:   "abnormal test results retrieved",
		Timestamp: time.Now(),
	})
}

// GET /medical-tests/critical
func (h *MedicalTestHandler) GetCriticalResults(c *gin.Context) {
	tests, err := h.service.GetCriticalResults(c.Request.Context())
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      tests,
		Message:   "critical test results retrieved",
		Timestamp: time.Now(),
	})
}
