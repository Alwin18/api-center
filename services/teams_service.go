package services

import (
	entity "api-center/Entity"
	"api-center/models"
	"time"

	"gorm.io/gorm"
)

func GetListTeams(db *gorm.DB, userId int) (response []*models.ListTeamResponse, err error) {
	var teams []*entity.Team
	if err := db.Model(&entity.Team{}).Select("id", "team_name", "icon").Where("user_id = ?", userId).Order("updated_at desc").Find(&teams).Error; err != nil {
		return nil, err
	}

	for _, team := range teams {
		response = append(response, &models.ListTeamResponse{
			ID:       team.ID,
			TeamName: team.TeamName,
			Icon:     team.Icon,
		})
	}

	return
}

func CreateTeam(db *gorm.DB, req models.CreateTeamsRequest) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	team := &entity.Team{
		TeamName:  req.TeamName,
		Icon:      req.Icon,
		UserID:    req.UserID,
		CreatedAt: time.Now(),
	}

	if err := tx.Model(&entity.Team{}).Create(&team).Error; err != nil {
		tx.Rollback()
		return err
	}

	// create default project
	project := &entity.Project{
		TeamID:      int(team.ID),
		ProjectName: "My Projects",
		Icon:        "fa-solid fa-folder",
		CreatedAt:   time.Now(),
		ProjectType: "HTTP",
	}
	if err := tx.Model(&entity.Project{}).Create(&project).Error; err != nil {
		tx.Rollback()
		return err
	}

	var role entity.Role
	var user entity.User

	// get role id owner
	roleID, err := role.GetId(tx, "owner")
	if err != nil {
		tx.Rollback()
		return err
	}

	// cerate first project member
	if err := tx.Model(&entity.ProjectMember{}).Create(&entity.ProjectMember{
		UserID:    int(req.UserID),
		RoleID:    int(roleID),
		ProjectID: int(project.ID),
		CreatedBy: user.GetUsername(tx, req.UserID),
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
