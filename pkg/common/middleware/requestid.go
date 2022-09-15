package middleware

import (
	"github.com/Mr-LvGJ/gobase/pkg/common/constant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID is a middleware that injects a 'X-Request-ID' into the context and request/response header of each request.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get(constant.XRequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Expose it for use in the application
		c.Set(constant.XRequestIDKey, requestID)

		// Set X-Request-ID header
		c.Writer.Header().Set(constant.XRequestIDKey, requestID)
		c.Next()
	}
}
