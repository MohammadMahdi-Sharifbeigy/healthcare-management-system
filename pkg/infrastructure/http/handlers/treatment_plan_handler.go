package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type TreatmentPlanHandler struct {
	service *services.TreatmentPlanService
}

func NewTreatmentPlanHandler(service *services.TreatmentPlanService) *TreatmentPlanHandler {
	return &TreatmentPlanHandler{service: service}
}

// POST /treatment-plans
func (h *TreatmentPlanHandler) CreatePlan(c *gin.Context) {
	var req TreatmentPlanCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_DATE",
			Message: "start_date must be YYYY-MM-DD format",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var endDate *time.Time
	if req.EndDate != "" {
		ed, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "INVALID_DATE",
				Message: "end_date must be YYYY-MM-DD format",
				Code:    http.StatusBadRequest,
			})
			return
		}
		endDate = &ed
	}

	plan := &entities.TreatmentPlan{
		PatientID:       req.PatientID,
		DoctorID:        req.DoctorID,
		Diagnosis:       req.Diagnosis,
		TreatmentType:   req.TreatmentType,
		StartDate:       startDate,
		EndDate:         endDate,
		SessionDuration: req.SessionDuration,
		Status:          "active",
		Goals:           req.Goals,
	}

	err = h.service.CreateTreatmentPlan(c.Request.Context(), plan)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Data:      plan,
		Message:   "treatment plan created successfully",
		Timestamp: time.Now(),
	})
}

// GET /treatment-plans/:id
func (h *TreatmentPlanHandler) GetPlan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "plan ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	plan, err := h.service.GetTreatmentPlan(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      plan,
		Message:   "treatment plan retrieved successfully",
		Timestamp: time.Now(),
	})
}

// PUT /treatment-plans/:id
func (h *TreatmentPlanHandler) UpdatePlan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "plan ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req TreatmentPlanUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	plan, err := h.service.GetTreatmentPlan(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	if req.Diagnosis != "" {
		plan.Diagnosis = req.Diagnosis
	}
	if req.TreatmentType != "" {
		plan.TreatmentType = req.TreatmentType
	}
	if req.EndDate != "" {
		ed, _ := time.Parse("2006-01-02", req.EndDate)
		plan.EndDate = &ed
	}
	if req.SessionDuration > 0 {
		plan.SessionDuration = req.SessionDuration
	}
	if req.Goals != "" {
		plan.Goals = req.Goals
	}
	if req.ProgressNotes != "" {
		plan.ProgressNotes = req.ProgressNotes
	}

	err = h.service.UpdateTreatmentPlan(c.Request.Context(), plan)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      plan,
		Message:   "treatment plan updated successfully",
		Timestamp: time.Now(),
	})
}

// PATCH /treatment-plans/:id/status
func (h *TreatmentPlanHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "plan ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req StatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	err = h.service.UpdatePlanStatus(c.Request.Context(), id, req.Status)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      nil,
		Message:   "treatment plan status updated",
		Timestamp: time.Now(),
	})
}

// GET /treatment-plans/patient/:patient_id
func (h *TreatmentPlanHandler) GetPatientPlans(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	plans, err := h.service.GetPatientPlans(c.Request.Context(), patientID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      plans,
		Message:   "patient treatment plans retrieved",
		Timestamp: time.Now(),
	})
}

// GET /treatment-plans/patient/:patient_id/active
func (h *TreatmentPlanHandler) GetActivePlans(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "patient ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	plans, err := h.service.GetActivePlans(c.Request.Context(), patientID)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      plans,
		Message:   "active treatment plans retrieved",
		Timestamp: time.Now(),
	})
}

// GET /treatment-plans/:id/progress
func (h *TreatmentPlanHandler) GetProgress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "plan ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	completedSessions := 0
	if cs := c.Query("completed_sessions"); cs != "" {
		if val, err := strconv.Atoi(cs); err == nil && val >= 0 {
			completedSessions = val
		}
	}

	progress, err := h.service.CalculateProgress(c.Request.Context(), id, completedSessions)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data: map[string]interface{}{
			"plan_id":             id,
			"completed_sessions":  completedSessions,
			"progress_percentage": progress,
		},
		Message:   "treatment progress calculated",
		Timestamp: time.Now(),
	})
}
