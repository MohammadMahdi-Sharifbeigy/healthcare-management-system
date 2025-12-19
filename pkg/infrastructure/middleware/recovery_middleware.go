package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware recovers from panic with detailed logging
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Get request ID
				requestID := ""
				if id, ok := c.Get("RequestID"); ok {
					requestID = id.(string)
				}

				// Get stack trace
				var buf strings.Builder
				buf.WriteString(fmt.Sprintf("Panic recovered!\n"))
				buf.WriteString(fmt.Sprintf("RequestID: %s\n", requestID))
				buf.WriteString(fmt.Sprintf("Error: %v\n", err))
				buf.WriteString(fmt.Sprintf("Time: %s\n", time.Now().Format(time.RFC3339)))
				buf.WriteString("Stack:\n")

				// Print stack frames
				pc := make([]uintptr, 32)
				numFrames := runtime.Callers(3, pc)
				frames := runtime.CallersFrames(pc[:numFrames])

				for {
					frame, more := frames.Next()
					buf.WriteString(fmt.Sprintf("  at %s:%d in %s\n",
						frame.File,
						frame.Line,
						frame.Function,
					))
					if !more {
						break
					}
				}

				log.Print(buf.String())

				// Return error response
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error:     "INTERNAL_SERVER_ERROR",
					Message:   "The server encountered an unexpected condition",
					Code:      http.StatusInternalServerError,
					RequestID: requestID,
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}
