package entity

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    int            `gorm:"not null" json:"user_id"`
	TeamName  string         `gorm:"size:50;not null" json:"team_name"`
	Icon      string         `gorm:"default:'fa-solid fa-people-group'" json:"icon"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Team) TableName() string {
	return "teams"
}
