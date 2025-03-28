package services

import (
	"github.com/Newbwen/automation-ops/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetUser(DB *gorm.DB, id string) (models.Users, error)
	CreateUser(DB *gorm.DB, user models.Users) (models.Users, error)
	UpdateUser(DB *gorm.DB, user models.Users) (models.Users, error)
	DeleteUser(DB *gorm.DB, id string) error
	GetEmail(DB *gorm.DB, email string) (string, error)
	GetUserByUsername(DB *gorm.DB, username string) error
}

func GetUser(DB *gorm.DB, id string) (models.Users, error) {
	var user models.Users
	err := DB.First(&user, id).Error
	return user, err
}

func CreateUser(DB *gorm.DB, user models.Users) (models.Users, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)
	err = DB.Create(&user).Error
	return user, err
}

func UpdateUser(DB *gorm.DB, user models.Users) (models.Users, error) {
	err := DB.Save(&user).Error
	return user, err
}

func DeleteUser(DB *gorm.DB, id string) error {
	err := DB.Delete(&models.Users{}, id).Error
	return err
}
func GetEmail(DB *gorm.DB, email string) (string, error) {
	user := models.Users{}
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
func GetUserByUsername(DB *gorm.DB, username string) error {
	user := models.Users{}
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
