package services

import (
	entity "api-center/Entity"
	"api-center/models"

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
