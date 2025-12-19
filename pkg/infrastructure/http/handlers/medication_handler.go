package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type MedicationHandler struct {
	service *services.MedicationService
}

func NewMedicationHandler(service *services.MedicationService) *MedicationHandler {
	return &MedicationHandler{service: service}
}

// POST /medications (admin only)
func (h *MedicationHandler) CreateMedication(c *gin.Context) {
	var req MedicationCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	medication := &entities.Medication{
		MedicationName: req.MedicationName,
		GenericName:    req.GenericName,
		Form:           req.Form,
		Strength:       req.Strength,
		Manufacturer:   req.Manufacturer,
		SideEffects:    req.SideEffects,
	}

	err := h.service.CreateMedication(c.Request.Context(), medication)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Data:      medication,
		Message:   "medication created successfully",
		Timestamp: time.Now(),
	})
}

// GET /medications/:id
func (h *MedicationHandler) GetMedication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "medication ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	medication, err := h.service.GetMedication(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      medication,
		Message:   "medication retrieved successfully",
		Timestamp: time.Now(),
	})
}

// GET /medications
func (h *MedicationHandler) GetMedications(c *gin.Context) {
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

	medications, err := h.service.GetAllMedications(c.Request.Context(), limit, offset)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, PaginationResponse{
		Data:   medications,
		Limit:  limit,
		Offset: offset,
	})
}

// GET /medications/search?query=name
func (h *MedicationHandler) SearchMedications(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "MISSING_QUERY",
			Message: "query parameter required",
			Code:    http.StatusBadRequest,
		})
		return
	}

	medications, err := h.service.SearchByName(c.Request.Context(), query)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      medications,
		Message:   "medication search completed",
		Timestamp: time.Now(),
	})
}

// GET /medications/generic/:name
func (h *MedicationHandler) GetByGenericName(c *gin.Context) {
	name := c.Param("name")
	medications, err := h.service.GetByGenericName(c.Request.Context(), name)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      medications,
		Message:   "medications retrieved by generic name",
		Timestamp: time.Now(),
	})
}

// GET /medications/form/:form
func (h *MedicationHandler) GetByForm(c *gin.Context) {
	form := c.Param("form")
	medications, err := h.service.GetByForm(c.Request.Context(), form)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      medications,
		Message:   "medications retrieved by form",
		Timestamp: time.Now(),
	})
}
