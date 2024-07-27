package middleware

import (
	"log"
	"net/http"

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
