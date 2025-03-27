package main

import (
	"fmt"
	"github.com/Newbwen/automation-ops/backend/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int    `json:"age"`
	Tel     string `json:"tel"`
	Address string `json:"address"`
}

func main() {
	db := *database.DB
	if db.Error != nil {
		fmt.Println(db.Error)
	}
	var user User
	if err := db.Where("id = ?", 7).First(&user).Error; err != nil {
		fmt.Println(err)
		return
	}
	result := map[string]interface{}{
		"name":    user.Name,
		"sex":     user.Sex,
		"age":     user.Age,
		"tel":     user.Tel,
		"address": user.Address,
	}
	fmt.Printf("name: %s, sex: %s, age: %d, tel: %s, address: %s", result["name"], result["sex"], result["age"], result["tel"], result["address"])
}
