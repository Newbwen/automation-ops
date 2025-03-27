package controllers

import (
	"github.com/Newbwen/automation-ops/backend/database"
	"github.com/Newbwen/automation-ops/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHosts(c *gin.Context) {
	var hosts []models.Host
	database.DB.Find(&hosts)
	c.JSON(http.StatusOK, hosts)
}
