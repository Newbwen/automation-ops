package database

import (
	"fmt"
	"github.com/Newbwen/automation-ops/backend/models"
	"github.com/Newbwen/automation-ops/backend/services"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type InitProject struct {
}

func (i *InitProject) Init() {
	DB.AutoMigrate(&models.Role{}, &models.Users{}, &models.Host{})
	createRole()
	addUser()
}

// 创建默认角色,如果存在则不创建
func createRole() {
	roles := []models.Role{
		{Name: "admin", Description: "管理员"},
		{Name: "user", Description: "普通用户"},
	}
	if err := services.GetRoleByName(DB, "admin"); err == nil {
		fmt.Println("admin role already exists")
		return
	}
	if err := services.GetRoleByName(DB, "user"); err == nil {
		fmt.Println("user role already exists")
		return
	}
	DB.Create(&roles)
}

// 创建默认用户
func addUser() {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("Admin@123"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	if err := services.GetUserByUsername(DB, "Admin"); err == nil {
		fmt.Println("Admin user already exists")
		return
	}
	user := models.Users{
		Username:  "Admin",
		Password:  string(hashedPassword),
		Email:     "admin@example.com",
		RoleID:    queryRoleByName(),
		LastLogin: time.Now(),
	}
	if err := DB.Create(&user).Error; err != nil {
		fmt.Println(err)
	}
}

// 查询角色ID
func queryRoleByName() string {
	var role models.Role
	if err := DB.Where("name = ?", "admin").First(&role).Error; err != nil {
		return ""
	}
	return role.ID
}
