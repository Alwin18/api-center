package common

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	Status  bool         `json:"status"`
	Code    int32        `json:"code"`
	Message string       `json:"message"`
	Data    *interface{} `json:"data,omitempty"`
}

// SuccessResponse generates a success response
func SuccessResponse(c *gin.Context, code int32, message string, data *interface{}) {
	c.JSON(int(code), BaseResponse{
		Status:  true,
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse generates an error response
func ErrorResponse(c *gin.Context, code int32, message string, err interface{}) {
	c.JSON(int(code), gin.H{
		"status":  false,
		"code":    code,
		"message": message,
		"error":   err,
	})
}
