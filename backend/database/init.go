package database

import (
	"github.com/Newbwen/automation-ops/backend/models"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"strconv"
)

var DB *gorm.DB

func init() {
	configFile, err := ioutil.ReadFile("config/mysql.yaml")
	if err != nil {
		panic(err)
	}
	config := models.Config{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}
	dsn := config.Database.User + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + strconv.Itoa(int(config.Database.Port)) + ")/" + config.Database.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
