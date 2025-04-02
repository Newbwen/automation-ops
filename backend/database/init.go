package database

import (
	"context"
	"github.com/Newbwen/automation-ops/backend/models"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"strconv"
	"time"
)

var DB *gorm.DB
var RedisClient *redis.Client

func init() {
	configFile, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		panic(err)
	}
	// begin mysql init
	mysqlConfig := models.MysqlConfig{}
	err = yaml.Unmarshal(configFile, &mysqlConfig)
	if err != nil {
		panic(err)
	}
	dsn := mysqlConfig.Database.User + ":" + mysqlConfig.Database.Password + "@tcp(" + mysqlConfig.Database.Host + ":" + strconv.Itoa(int(mysqlConfig.Database.Port)) + ")/" + mysqlConfig.Database.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// end mysql init
	// begin redis init
	redisConfig := models.RedisConfig{}
	err = yaml.Unmarshal(configFile, &redisConfig)
	if err != nil {
		panic(err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Redis.Host + ":" + strconv.Itoa(redisConfig.Redis.Port),
		Password: redisConfig.Redis.Password,
		DB:       redisConfig.Redis.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		panic("Redis connection failed:" + err.Error())
	}
	// end redis init
}
