package services

import (
	"github.com/Newbwen/automation-ops/backend/models"
	"gorm.io/gorm"
)

type RoleService interface {
	GetRole(DB *gorm.DB, id int) (models.Role, error)
	CreateRole(DB *gorm.DB, role models.Role) (models.Role, error)
	UpdateRole(DB *gorm.DB, role models.Role) (models.Role, error)
	DeleteRole(DB *gorm.DB, id int) error
}

func GetRole(DB *gorm.DB, id int) (models.Role, error) {
	var role models.Role
	err := DB.First(&role, id).Error
	return role, err
}

func CreateRole(DB *gorm.DB, role models.Role) (models.Role, error) {
	err := DB.Create(&role).Error
	return role, err
}

func UpdateRole(DB *gorm.DB, role models.Role) (models.Role, error) {
	err := DB.Save(&role).Error
	return role, err
}

func DeleteRole(DB *gorm.DB, id int) error {
	err := DB.Delete(&models.Role{}, id).Error
	return err
}
