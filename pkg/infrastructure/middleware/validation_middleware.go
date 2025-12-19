package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ContentTypeMiddleware validates Content-Type header for POST/PUT/PATCH requests
func ContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			contentType := c.GetHeader("Content-Type")
			if contentType == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "MISSING_CONTENT_TYPE",
					"message": "Content-Type header required for POST/PUT/PATCH requests",
				})
				c.Abort()
				return
			}

			if contentType != "application/json" {
				c.JSON(http.StatusUnsupportedMediaType, gin.H{
					"error": "UNSUPPORTED_MEDIA_TYPE",
					"message": "Content-Type must be application/json",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// MaxRequestSizeMiddleware limits request body size
func MaxRequestSizeMiddleware(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		
		if c.Request.ContentLength > maxBytes {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": "REQUEST_TOO_LARGE",
				"message": "Request body exceeds maximum allowed size",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
