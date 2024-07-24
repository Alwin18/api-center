package middleware

import (
	"log"

	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a middleware function that logs the details of each request.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Stop timer
		duration := time.Since(startTime)

		// Get status code
		statusCode := c.Writer.Status()

		// Log details
		log.Printf(
			"%s %s %d %s",
			c.Request.Method,
			c.Request.RequestURI,
			statusCode,
			duration,
		)
	}
}
