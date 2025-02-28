package controllers

import (
	"api-center/common"
	"api-center/models"
	"api-center/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeamHandler struct {
	DB *gorm.DB
}

func (h *TeamHandler) GetListTeams(c *gin.Context) {
	var request models.ListTeamRequest

	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", "user_id must be an integer")
		return
	}

	request.UserID = userID

	if request.UserID == 0 {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", "user_id is required")
		return
	}

	teams, err := services.GetListTeams(h.DB, request.UserID)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, "teams fetched successfully", teams)
}

func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var request models.CreateTeamsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if request.UserID == 0 || request.TeamName == "" {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", "user_id, team_name is required")
		return
	}

	if err := services.CreateTeam(h.DB, request); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, "team created successfully", nil)
}
