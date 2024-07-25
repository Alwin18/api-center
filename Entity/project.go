package entity

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	TeamID      int            `gorm:"not null" json:"team_id"`
	ProjectName string         `gorm:"size:50;not null" json:"project_name"`
	Icon        string         `gorm:"size:50;default:'fa-solid fa-folder'" json:"icon"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	ProjectType string         `gorm:"size:50;not null" json:"project_type"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Project) TableName() string {
	return "projects"
}
