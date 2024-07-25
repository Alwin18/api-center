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

type ProjectHandler struct {
	DB *gorm.DB
}

func (h *ProjectHandler) GetListProjects(c *gin.Context) {
	var request models.ListProjectsRequest
	teamIDStr := c.Query("team_id")
	teamID, err := strconv.Atoi(teamIDStr)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", "team_id is required")
		return
	}

	request.TeamID = teamID

	if request.TeamID == 0 {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", "team_id is required")
		return
	}

	projects, err := services.GetListProjects(h.DB, teamID)
	if err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "internal server error", err.Error())
		return
	}

	common.SuccessResponse(c, http.StatusOK, "projects fetched successfully", projects)
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var request models.CreateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if request.TeamID == 0 || request.ProjectName == "" || request.Icon == "" || request.ProjectType == "" {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", "team_id, project_name, icon, project_type is required")
		return
	}

	if err := services.CreateProject(h.DB, &request); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, "project created successfully", nil)
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	var request models.UpdateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if request.ID == 0 || request.ProjectName == "" || request.Icon == "" {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", "id, project_name, icon is required")
		return
	}

	if err := services.UpdateProject(h.DB, &request); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, "project updated successfully", nil)
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	var request models.DeleteProjectsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if request.ID == 0 {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", "id is required")
		return
	}

	if err := services.DeleteProject(h.DB, request); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, "project deleted successfully", nil)
}

func (h *ProjectHandler) AddProjectFavorite(c *gin.Context) {
	var request models.AddFavoriteProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if request.UserID == 0 || request.ProjectID == 0 {
		common.ErrorResponse(c, http.StatusBadRequest, "invalid request", "user_id, project_id is required")
		return
	}

	if err := services.AddProjectFavorite(h.DB, &request); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, "project added to favorite successfully", nil)
}
