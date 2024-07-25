package services

import (
	entity "api-center/Entity"
	"api-center/models"
	"api-center/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.CreateUserRequest) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// check if email or user_name is exist
	var existEmail string
	if err := tx.Model(&entity.User{}).Select("email").Where("email = ?", *user.Email).Scan(&existEmail).Error; err != nil {
		tx.Rollback()
		return err
	}

	if existEmail != "" {
		tx.Rollback()
		return errors.New("email already exist")
	}

	var existUser *string
	if err := tx.Model(&entity.User{}).Select("user name").Where("user_name = ?", *user.UserName).Scan(&existUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	if existUser != nil {
		tx.Rollback()
		return errors.New("user_name already exist")
	}

	saveUser := &entity.User{
		UserName:  *user.UserName,
		Password:  *user.Password,
		Email:     *user.Email,
		CreatedAt: time.Now(),
	}

	if err := db.Model(&entity.User{}).Create(&saveUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	// create first team
	saveTeam := &entity.Team{
		UserID:    int(saveUser.ID),
		TeamName:  "Personal Teams",
		Icon:      "fa-solid fa-people-group",
		CreatedAt: time.Now(),
	}

	if err := tx.Model(&entity.Team{}).Create(&saveTeam).Error; err != nil {
		tx.Rollback()
		return err
	}

	// create first project
	saveProject := &entity.Project{
		TeamID:      int(saveTeam.ID),
		ProjectName: "My Projects",
		Icon:        "fa-solid fa-folder",
		CreatedAt:   time.Now(),
	}
	if err := tx.Model(&entity.Project{}).Create(&saveProject).Error; err != nil {
		tx.Rollback()
		return err
	}

	var role entity.Role

	// get role id owner
	roleID, err := role.GetId(tx, "owner")
	if err != nil {
		tx.Rollback()
		return err
	}

	// cerate first project member
	if err := tx.Model(&entity.ProjectMember{}).Create(&entity.ProjectMember{
		ID:        roleID,
		UserID:    int(saveUser.ID),
		RoleID:    int(roleID),
		ProjectID: int(saveProject.ID),
		CreatedBy: saveUser.UserName,
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

func Login(db *gorm.DB, user *models.LoginRequest) (*models.LoginResponse, error) {
	var existUser *entity.User
	if err := db.Model(&entity.User{}).Select("id", "user_name", "email", "password").Where("user_name = ?", *user.UserName).Scan(&existUser).Error; err != nil {
		return nil, err
	}

	if existUser == nil {
		return nil, errors.New("user not found")
	}

	// check password
	if !utils.CheckPasswordHash(*user.Password, existUser.Password) {
		return nil, errors.New("password not match")
	}

	// generate token
	token, err := utils.GenerateToken(existUser.UserName)
	if err != nil {
		return nil, err
	}

	data := &models.LoginResponse{
		Token:   &token,
		UseName: &existUser.UserName,
		Email:   &existUser.Email,
		ID:      &existUser.ID,
	}

	return data, nil
}
