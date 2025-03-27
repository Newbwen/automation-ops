package main

import (
	"fmt"
	"github.com/Newbwen/automation-ops/backend/controllers"
	"github.com/Newbwen/automation-ops/backend/database"
	"github.com/Newbwen/automation-ops/backend/models"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	database.DB.AutoMigrate(&models.Users{}, &models.Role{})
	r := gin.Default()
	controllers.RegisterUser(r, database.DB)
	controllers.LoginUser(r, database.DB)
	r.Run(":8080")
}

func createRole() {
	db := database.DB
	role := models.Role{
		Name:        "admin",
		Description: "administrator",
	}
	db.Create(&role)
}

func queryRoleByName() string {
	db := database.DB
	var role models.Role
	if err := db.Where("name = ?", "admin").First(&role).Error; err != nil {
		return ""
	}
	return role.ID
}
func addUser() {
	db := database.DB
	user := models.Users{
		Username:  "zhangsan",
		Password:  "Admin@123",
		Email:     "admin@admin.com",
		RoleID:    queryRoleByName(),
		LastLogin: time.Now(),
	}
	if err := db.Create(&user).Error; err != nil {
		fmt.Println(err)
	}
}
