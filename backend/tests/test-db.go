package main

import (
	"context"
	"fmt"
	"github.com/Newbwen/automation-ops/backend/database"
	"github.com/go-redis/redis/v8"
)

func AddForRedis(rdb *redis.Client, key string, value string) error {
	err := rdb.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	rdb := database.RedisClient
	if rdb == nil {
		fmt.Println("Redis is not initialized")
		return
	}
	err := AddForRedis(rdb, "test", "test")
	if err != nil {
		fmt.Println(err)
	}
}
