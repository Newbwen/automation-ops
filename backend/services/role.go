package services

import (
	"github.com/Newbwen/automation-ops/backend/models"
	"gorm.io/gorm"
)

type RoleService interface {
	GetRole(DB *gorm.DB, id string) (models.Role, error)
	GetRoleByName(DB *gorm.DB, name string) error
	CreateRole(DB *gorm.DB, role models.Role) (models.Role, error)
	UpdateRole(DB *gorm.DB, role models.Role) (models.Role, error)
	DeleteRole(DB *gorm.DB, id string) error
	GetDefaultId(DB *gorm.DB, name string) (string, error)
}

func GetRole(DB *gorm.DB, id string) (models.Role, error) {
	var role models.Role
	err := DB.First(&role, id).Error
	return role, err
}
func GetRoleByName(DB *gorm.DB, name string) error {
	var role models.Role
	err := DB.Where("name = ?", name).First(&role).Error
	return err
}

func CreateRole(DB *gorm.DB, role models.Role) (models.Role, error) {
	err := DB.Create(&role).Error
	return role, err
}

func UpdateRole(DB *gorm.DB, role models.Role) (models.Role, error) {
	err := DB.Save(&role).Error
	return role, err
}

func DeleteRole(DB *gorm.DB, id string) error {
	err := DB.Delete(&models.Role{}, id).Error
	return err
}

func GetDefaultId(DB *gorm.DB, name string) (string, error) {
	var role models.Role
	err := DB.Where("name = ?", name).First(&role).Error
	return role.ID, err
}
