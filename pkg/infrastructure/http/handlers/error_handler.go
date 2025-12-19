package handlers

import (
	"errors"
	"net/http"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

func handleServiceError(c interface{ JSON(int, interface{}) }, err error) {
	var customErr *entities.CustomError
	if errors.As(err, &customErr) {
		switch customErr.Code {
		case "INVALID_INPUT":
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   customErr.Code,
				Message: customErr.Message,
				Code:    http.StatusBadRequest,
			})
		case "NOT_FOUND":
			c.JSON(http.StatusNotFound, ErrorResponse{
				Error:   customErr.Code,
				Message: customErr.Message,
				Code:    http.StatusNotFound,
			})
		case "CONFLICT":
			c.JSON(http.StatusConflict, ErrorResponse{
				Error:   customErr.Code,
				Message: customErr.Message,
				Code:    http.StatusConflict,
			})
		case "UNAUTHORIZED":
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Error:   customErr.Code,
				Message: customErr.Message,
				Code:    http.StatusUnauthorized,
			})
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error:   "INTERNAL_ERROR",
				Message: customErr.Message,
				Code:    http.StatusInternalServerError,
			})
		}
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Error:   "INTERNAL_ERROR",
		Message: err.Error(),
		Code:    http.StatusInternalServerError,
	})
}
