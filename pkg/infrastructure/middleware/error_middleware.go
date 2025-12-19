package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	Code       int    `json:"code"`
	RequestID  string `json:"request_id,omitempty"`
	Timestamp  string `json:"timestamp,omitempty"`
}

// ErrorHandlerMiddleware recovers from panics and handles errors
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Get request ID if available
				requestID := ""
				if id, ok := c.Get("RequestID"); ok {
					requestID = id.(string)
				}

				// Log panic
				fmt.Printf("[PANIC] RequestID: %s\n", requestID)
				fmt.Printf("[PANIC] Error: %v\n", err)
				fmt.Printf("[PANIC] Stack:\n%s\n", debug.Stack())

				// Return error response
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error:     "INTERNAL_SERVER_ERROR",
					Message:   "An unexpected error occurred",
					Code:      http.StatusInternalServerError,
					RequestID: requestID,
				})
				c.Abort()
			}
		}()

		c.Next()

		// Check for aborted requests
		if len(c.Errors) > 0 {
			requestID := ""
			if id, ok := c.Get("RequestID"); ok {
				requestID = id.(string)
			}

			// Log error
			for _, err := range c.Errors {
				fmt.Printf("[ERROR] RequestID: %s - %v\n", requestID, err.Error())
			}
		}
	}
}

// NotFoundHandler handles 404 errors
func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "NOT_FOUND",
			Message: fmt.Sprintf("Endpoint %s %s not found", c.Request.Method, c.Request.URL.Path),
			Code:    http.StatusNotFound,
		})
	}
}

// MethodNotAllowedHandler handles 405 errors
func MethodNotAllowedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, ErrorResponse{
			Error:   "METHOD_NOT_ALLOWED",
			Message: fmt.Sprintf("Method %s not allowed for %s", c.Request.Method, c.Request.URL.Path),
			Code:    http.StatusMethodNotAllowed,
		})
	}
}
