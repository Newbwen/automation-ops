package controllers

import (
	"github.com/Newbwen/automation-ops/backend/database"
	"github.com/Newbwen/automation-ops/backend/models"
	"github.com/Newbwen/automation-ops/backend/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type RegisterInput struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context) {
	regInput := RegisterInput{}
	var result models.Users
	var role_name string
	var role_id string
	if err := c.ShouldBindJSON(&regInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误！" + err.Error()})
		return
	}
	if regInput.Password != regInput.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "两次密码输入不一致！"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败！" + err.Error()})
		return
	}
	if _, err := services.GetEmail(database.DB, regInput.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "邮箱已注册！"})
		return
	}
	role_name = "user"
	role_id, _ = services.GetDefaultId(database.DB, role_name)
	user := models.Users{
		Username:  "普通用户",
		Password:  string(hashedPassword),
		Email:     regInput.Email,
		RoleID:    role_id,
		LastLogin: time.Now(),
	}
	if result, err = services.CreateUser(database.DB, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败！" + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)

}
func LoginUser(c *gin.Context) {
	req := models.Users{}
	if err := c.ShouldBindJSON(&req); err != nil {
		req := models.Users{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求！"})
			return
		}
		var user models.Users
		if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在！"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误！"})
			return
		}
		user.LastLogin = time.Now()
		database.DB.Save(&user)
		c.JSON(http.StatusOK, user)
	}
}
