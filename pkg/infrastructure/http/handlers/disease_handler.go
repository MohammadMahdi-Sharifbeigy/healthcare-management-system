package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DiseaseHandler struct {
	service *services.DiseaseService
}

func NewDiseaseHandler(service *services.DiseaseService) *DiseaseHandler {
	return &DiseaseHandler{service: service}
}

// POST /diseases (admin only)
func (h *DiseaseHandler) CreateDisease(c *gin.Context) {
	var req DiseaseCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	disease := &entities.Disease{
		DiseaseName: req.DiseaseName,
		Category:    req.Category,
		ICDCode:     req.ICDCode,
		Description: req.Description,
	}

	err := h.service.CreateDisease(c.Request.Context(), disease)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Data:      disease,
		Message:   "disease created successfully",
		Timestamp: time.Now(),
	})
}

// GET /diseases/:id
func (h *DiseaseHandler) GetDisease(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "disease ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	disease, err := h.service.GetDisease(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      disease,
		Message:   "disease retrieved successfully",
		Timestamp: time.Now(),
	})
}

// GET /diseases
func (h *DiseaseHandler) GetDiseases(c *gin.Context) {
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

	diseases, err := h.service.GetAllDiseases(c.Request.Context(), limit, offset)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, PaginationResponse{
		Data:   diseases,
		Limit:  limit,
		Offset: offset,
	})
}

// GET /diseases/search?query=name
func (h *DiseaseHandler) SearchDiseases(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "MISSING_QUERY",
			Message: "query parameter required",
			Code:    http.StatusBadRequest,
		})
		return
	}

	diseases, err := h.service.SearchByName(c.Request.Context(), query)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      diseases,
		Message:   "disease search completed",
		Timestamp: time.Now(),
	})
}

// GET /diseases/icd/:code
func (h *DiseaseHandler) GetByICDCode(c *gin.Context) {
	code := c.Param("code")
	disease, err := h.service.GetByICDCode(c.Request.Context(), code)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      disease,
		Message:   "disease retrieved by ICD code",
		Timestamp: time.Now(),
	})
}

// GET /diseases/category/:cat
func (h *DiseaseHandler) GetByCategory(c *gin.Context) {
	category := c.Param("cat")
	diseases, err := h.service.GetByCategory(c.Request.Context(), category)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      diseases,
		Message:   "diseases retrieved by category",
		Timestamp: time.Now(),
	})
}
