package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProjectMember struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    int            `gorm:"not null" json:"user_id"`
	RoleID    int            `gorm:"not null" json:"role_id"`
	ProjectID int            `gorm:"not null" json:"project_id"`
	CreatedBy string         `gorm:"not null" json:"created_by"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

func (ProjectMember) TableName() string {
	return "project_member"
}
