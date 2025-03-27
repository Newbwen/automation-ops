package controllers

import (
	"net/http"
	"time"

	"github.com/Newbwen/automation-ops/backend/database"
	"github.com/Newbwen/automation-ops/backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(e *gin.Engine, db *gorm.DB) {
	e.POST("/register", func(c *gin.Context) {
		req := models.Users{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求！"})
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败！"})
			return
		}
		user := models.Users{
			Username:  req.Username,
			Password:  string(hashedPassword),
			Email:     req.Email,
			RoleID:    req.RoleID,
			LastLogin: time.Now(),
		}
		err = db.Create(&user).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败！"})
			return
		}
		c.JSON(http.StatusCreated, user)
	})
}
func LoginUser(e *gin.Engine, db *gorm.DB) {
	e.POST("/login", func(c *gin.Context) {
		req := models.Users{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求！"})
			return
		}
		db := database.DB
		var user models.Users
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在！"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误！"})
			return
		}
		user.LastLogin = time.Now()
		db.Save(&user)
		c.JSON(http.StatusOK, user)
	})
}
