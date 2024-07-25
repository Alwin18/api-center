package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProjectFavorite struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    int            `gorm:"not null" json:"user_id"`
	ProjectID int            `gorm:"default:null" json:"project_id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

func (ProjectFavorite) TableName() string {
	return "project_favorite"
}
