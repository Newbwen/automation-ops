package controllers

import (
	"github.com/Newbwen/automation-ops/backend/config"
	"github.com/Newbwen/automation-ops/backend/database"
	"github.com/Newbwen/automation-ops/backend/models"
	"github.com/Newbwen/automation-ops/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type RegisterInput struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	//CaptchaID       string `json:"captcha_id" binding:"required"`    // 验证码ID
	//CaptchaValue    string `json:"captcha_value" binding:"required"` // 验证码值
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 验证密码是否一致
	if input.Password != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码不一致"})
		return
	}

	// 查询用户是否存在
	var user models.Users
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "邮箱已存在"})
		return
	}

	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}
	roleID, _ := services.GetDefaultId(database.DB, "user")
	user = models.Users{
		Username:  "普通用户",
		Password:  string(hashedPassword),
		Email:     input.Email,
		RoleID:    roleID,
		LastLogin: time.Now(),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}
func LoginUser(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 查询用户是否存在
	var user models.Users
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		// 模糊提示避免信息泄露（不明确说明是邮箱不存在）
		c.JSON(http.StatusUnauthorized, gin.H{"error": "邮箱或密码错误"})
		return
	}

	// 验证密码哈希
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "邮箱或密码错误"})
		return
	}
	//生成JWTtoken，有效期72小时
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(), // 有效期72小时
	})

	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成token失败"})
		return
	}

	user.LastLogin = time.Now()
	database.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	}) // 返回给客户端存储在 localStorage 中
}

func GetUser(c *gin.Context) {
	// 从请求头中获取token
	userID, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未认证的用户"})
		return
	}

	//查询用户信息，排查敏感字段
	var user models.Users
	if err := database.DB.Select("id", "username", "email").Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})

}

// 修改密码
func ChangePassword(c *gin.Context) {
	// 从请求头中获取token
	userID, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未认证的用户"})
		return
	}

	// 接收参数
	var input struct {
		OldPassword     string `json:"old_password" binding:"required"`
		NewPassword     string `json:"new_password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 查询用户信息
	var user models.Users
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	// 验证旧密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "旧密码错误"})
		return
	}
	// 验证新密码是否一致
	if input.NewPassword != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "新密码不一致"})
		return
	}

	// 更新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新密码失败"})
		return
	}
	user.Password = string(hashedPassword)
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新密码失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "密码修改成功"})

}

// 用户信息
func UserInfo(c *gin.Context) {
	// 从请求头中获取token
	userID, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未认证的用户"})
		return
	}

	// 查询用户信息
	var user models.Users
	if err := database.DB.Select("id", "username", "email", "role_id", "created_at", "last_login").Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	// 查询角色信息
	var role models.Role
	if err := database.DB.Select("name").Where("id = ?", user.RoleID).First(&role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "角色不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username":   user.Username,
		"email":      user.Email,
		"role":       role.Name,
		"created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
		"last_login": user.LastLogin.Format("2006-01-02 15:04:05"),
	})
}
