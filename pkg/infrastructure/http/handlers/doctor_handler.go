package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type DoctorHandler struct {
	service *services.DoctorService
}

func NewDoctorHandler(service *services.DoctorService) *DoctorHandler {
	return &DoctorHandler{service: service}
}

// POST /doctors
func (h *DoctorHandler) CreateDoctor(c *gin.Context) {
	var req DoctorCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	doctor := &entities.Doctor{
		DoctorID:       req.DoctorID,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		LicenseNumber:  req.LicenseNumber,
		Specialization: req.Specialization,
		Department:     req.Department,
	}

	err := h.service.RegisterDoctor(c.Request.Context(), doctor)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Data:      doctor,
		Message:   "doctor registered successfully",
		Timestamp: time.Now(),
	})
}

// GET /doctors/:id
func (h *DoctorHandler) GetDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "doctor ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	doctor, err := h.service.GetDoctorProfile(c.Request.Context(), id)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      doctor,
		Message:   "doctor retrieved successfully",
		Timestamp: time.Now(),
	})
}

// PUT /doctors/:id
func (h *DoctorHandler) UpdateDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "INVALID_ID",
			Message: "doctor ID must be a number",
			Code:    http.StatusBadRequest,
		})
		return
	}

	var req DoctorCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	doctor := &entities.Doctor{
		DoctorID:       id,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		LicenseNumber:  req.LicenseNumber,
		Specialization: req.Specialization,
		Department:     req.Department,
	}

	err = h.service.UpdateDoctorInfo(c.Request.Context(), doctor)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      doctor,
		Message:   "doctor updated successfully",
		Timestamp: time.Now(),
	})
}

// GET /doctors
func (h *DoctorHandler) GetDoctors(c *gin.Context) {
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

	doctors, err := h.service.GetAllDoctors(c.Request.Context(), limit, offset)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, PaginationResponse{
		Data:   doctors,
		Limit:  limit,
		Offset: offset,
	})
}

// GET /doctors/specialization/:spec
func (h *DoctorHandler) GetDoctorsBySpecialization(c *gin.Context) {
	specialization := c.Param("spec")
	doctors, err := h.service.GetDoctorsBySpecialization(c.Request.Context(), specialization)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      doctors,
		Message:   "doctors retrieved by specialization",
		Timestamp: time.Now(),
	})
}

// GET /doctors/department/:dept
func (h *DoctorHandler) GetDoctorsByDepartment(c *gin.Context) {
	department := c.Param("dept")
	doctors, err := h.service.GetDoctorsByDepartment(c.Request.Context(), department)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Data:      doctors,
		Message:   "doctors retrieved by department",
		Timestamp: time.Now(),
	})
}
