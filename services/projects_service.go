package services

import (
	entity "api-center/Entity"
	"api-center/models"
	"time"

	"gorm.io/gorm"
)

func GetListProjects(db *gorm.DB, teamId int) (response []*models.ListProjectsResponse, err error) {
	var projects []*entity.Project

	if err := db.Model(&entity.Project{}).Select("id", "project_name", "icon").Where("team_id = ?", teamId).Order("updated_at desc").Find(&projects).Error; err != nil {
		return nil, err
	}

	for _, project := range projects {
		response = append(response, &models.ListProjectsResponse{
			ID:          project.ID,
			ProjectName: project.ProjectName,
			Icon:        project.Icon,
		})
	}

	return
}

func CreateProject(db *gorm.DB, project *models.CreateProjectRequest) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&entity.Project{}).Create(&entity.Project{
		TeamID:      project.TeamID,
		ProjectName: project.ProjectName,
		Icon:        project.Icon,
		CreatedAt:   time.Now(),
		ProjectType: project.ProjectType,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func DeleteProject(db *gorm.DB, req models.DeleteProjectsRequest) error {
	if err := db.Unscoped().Model(&entity.Project{}).Where("id = ?", req.ID).Delete(&entity.Project{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProject(db *gorm.DB, project *models.UpdateProjectRequest) error {
	if err := db.Model(&entity.Project{}).Where("id = ?", project.ID).Updates(&entity.Project{
		ProjectName: project.ProjectName,
		Icon:        project.Icon,
		UpdatedAt:   time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func AddProjectFavorite(db *gorm.DB, req *models.AddFavoriteProjectRequest) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := db.Model(&entity.ProjectFavorite{}).Create(&entity.ProjectFavorite{
		ProjectID: req.ProjectID,
		UserID:    req.UserID,
		CreatedAt: time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
