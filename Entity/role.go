package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:20;not null" json:"name"`
	Slug      string         `gorm:"size:20;not null" json:"slug"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Role) TableName() string {
	return "roles"
}

func (*Role) GetId(db *gorm.DB, slug string) (uint, error) {
	var id uint
	if err := db.Model(&Role{}).Select("id").Where("slug = ?", slug).First(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}
