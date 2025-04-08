package services

import (
	"github.com/Newbwen/automation-ops/backend/models"
	"gorm.io/gorm"
)

type HostService struct {
}

type HostServiceInterface interface {
	GetHost(DB *gorm.DB, id string) (models.Host, error)
	GetHosts(DB *gorm.DB) ([]models.Host, error)
	CreateHost(DB *gorm.DB, host models.Host) (models.Host, error)
	UpdateHost(DB *gorm.DB, host models.Host) (models.Host, error)
	DeleteHost(DB *gorm.DB, id string) error
}

func (s *HostService) GetHost(DB *gorm.DB, id string) (models.Host, error) {
	var host models.Host
	err := DB.First(&host, id).Error
	return host, err
}

// 分页获取主机列表
func (s *HostService) GetHosts(DB *gorm.DB, page, size int) ([]models.Host, error) {
	var hosts []models.Host
	err := DB.Limit(size).Offset((page - 1) * size).Find(&hosts).Error
	return hosts, err
}

func (s *HostService) CreateHost(DB *gorm.DB, host models.Host) (models.Host, error) {
	err := DB.Create(&host).Error
	return host, err
}

func (s *HostService) UpdateHost(DB *gorm.DB, host models.Host) (models.Host, error) {
	err := DB.Save(&host).Error
	return host, err
}

func (s *HostService) DeleteHost(DB *gorm.DB, id string) error {
	err := DB.Delete(&models.Host{}, id).Error
	return err
}
