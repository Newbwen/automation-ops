package controllers

import (
	"net/http"

	"github.com/Newbwen/automation-ops/backend/database"
	"github.com/Newbwen/automation-ops/backend/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Pagination struct {
	CurrentPage int           `json:"current_page"`
	PageSize    int           `json:"page_size"`
	Total       int64         `json:"total"`
	Data        []models.Host `json:"data"`
}

type HostForm struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port int    `json:"port"`
	Os   string `json:"os"`
}

// 编辑主机
func EditHost(c *gin.Context) {
	hostID := c.Param("id")
	var host models.Host
	// 根据ID查询主机
	var hostForm HostForm
	database.DB.First(&host, "id = ?", hostID)
	if err := c.ShouldBindJSON(&hostForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	host.Name = hostForm.Name
	host.IP = hostForm.IP
	host.Port = hostForm.Port
	host.Os = hostForm.Os

	if err := database.DB.Save(&host).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新主机失败"})
		return
	}
	c.JSON(http.StatusOK, host)
}

func AddHost(c *gin.Context) {
	var host models.Host
	var hostForm HostForm
	if err := c.ShouldBindJSON(&hostForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	host.Name = hostForm.Name
	host.IP = hostForm.IP
	host.Port = hostForm.Port
	host.Os = hostForm.Os
	if err := database.DB.Create(&host).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create host",
		})
		return
	}
	c.JSON(http.StatusCreated, host)
}

// 分页获取主机列表
func GetHosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	searchQuery := c.Query("search")

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}
	//构建查询
	query := database.DB.Model(&models.Host{})
	if searchQuery != "" {
		query = query.Where("name LIKE ? OR ip LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}
	var total int64
	query.Count(&total)
	var hosts []models.Host
	offset := (page - 1) * pageSize
	query.Limit(pageSize).Offset(offset).Find(&hosts)
	c.JSON(http.StatusOK, Pagination{
		CurrentPage: page,
		PageSize:    pageSize,
		Total:       total,
		Data:        hosts,
	})
}

// 删除主机
func DeleteHost(c *gin.Context) {
	hostID := c.Param("id")
	var host models.Host
	database.DB.First(&host, "id = ?", hostID)
	if err := database.DB.Delete(&host).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete host",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Host deleted successfully",
	})
}
