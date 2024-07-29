package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserName  string         `gorm:"size:12;not null" json:"user_name"`
	Password  string         `gorm:"type:text;not null" json:"password"`
	Email     string         `gorm:"size:50;not null;unique" json:"email"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (*User) GetUsername(db *gorm.DB, id int) string {
	var username string
	if err := db.Model(&User{}).Select("user_name").Where("id = ?", id).First(&username).Error; err != nil {
		return ""
	}
	return username
}
