package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthToken struct {
	UserID    int
	Email     string
	Role      string // admin, doctor, nurse, patient
	ExpiresAt int64
}

// AuthMiddleware validates JWT tokens (placeholder for real JWT implementation)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "MISSING_AUTH",
				"message": "Authorization header required",
			})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "INVALID_AUTH_HEADER",
				"message": "Authorization header format: Bearer <token>",
			})
			c.Abort()
			return
		}

		token := parts[1]

		// TODO: Implement actual JWT validation
		// For now, accept any non-empty token
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "INVALID_TOKEN",
				"message": "Token is empty",
			})
			c.Abort()
			return
		}

		// In production, validate JWT token and extract claims
		// For this example, we'll set a placeholder user
		c.Set("user_id", 1)
		c.Set("user_email", "user@hospital.com")
		c.Set("user_role", "doctor")

		c.Next()
	}
}

// RoleMiddleware checks if user has required role
func RoleMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED",
				"message": "User not authenticated",
			})
			c.Abort()
			return
		}

		role := userRole.(string)
		hasRole := false
		for _, required := range requiredRoles {
			if required == role || required == "*" {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "FORBIDDEN",
				"message": "Insufficient permissions for this operation",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// OptionalAuthMiddleware allows but doesn't require authentication
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				token := parts[1]
				if token != "" {
					// Token is valid, set user info
					c.Set("user_id", 1)
					c.Set("user_email", "user@hospital.com")
					c.Set("user_role", "doctor")
					c.Set("authenticated", true)
				}
			}
		}
		
		// Allow request even if no auth provided
		c.Set("authenticated", false)
		c.Next()
	}
}
