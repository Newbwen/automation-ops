package services

import (
	"github.com/Newbwen/automation-ops/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetUser(DB *gorm.DB, id int) (models.Users, error)
	CreateUser(DB *gorm.DB, user models.Users) (models.Users, error)
	UpdateUser(DB *gorm.DB, user models.Users) (models.Users, error)
	DeleteUser(DB *gorm.DB, id int) error
}

func GetUser(DB *gorm.DB, id int) (models.Users, error) {
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

func DeleteUser(DB *gorm.DB, id int) error {
	err := DB.Delete(&models.Users{}, id).Error
	return err
}
