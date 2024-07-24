package controllers

import (
	"api-center/common"
	"api-center/models"
	"api-center/services"
	"api-center/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if user.Email == nil || user.UserName == nil || user.Password == nil {
		missingField := ""
		switch {
		case user.Email == nil:
			missingField = "email"
		case user.UserName == nil:
			missingField = "user_name"
		case user.Password == nil:
			missingField = "password"
		}
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", missingField+" is required")
		return
	}

	hashPassword, err := utils.HashPassword(*user.Password)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	userName := strings.ToLower(*user.UserName)

	user = models.CreateUserRequest{
		UserName: &userName,
		Email:    user.Email,
		Password: &hashPassword,
	}

	err = services.CreateUser(h.DB, &user)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusCreated, "user created successfully", nil)
}
