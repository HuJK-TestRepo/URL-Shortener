package cache_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/steveyiyo/url-shortener/internal/tools"
)

var Redis *redis.Client

// Init Client
func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

// Add data to Redis
func AddData(key string, value string, second time.Duration) bool {
	ctx := context.Background()

	err := Redis.Set(ctx, key, value, second*time.Second).Err()
	return tools.ErrCheck(err)
}

// Query data from Redis
func QueryData(key string) (bool, string) {
	ctx := context.Background()

	var return_status bool
	var return_value string

	value, err := Redis.Get(ctx, key).Result()
	tools.ErrCheck(err)

	if err == redis.Nil {
		return_status = false
		return_value = ""
	} else if !tools.ErrCheck(err) {
		log.Println(err)
	} else {
		return_status = true
		return_value = value
	}
	return return_status, return_value
}

// It's a test function.
func TestMain(t *testing.T) {
	InitRedis()
	AddData("hi", "pong", 5)
	status, data := QueryData("hi")
	if status {
		fmt.Println(data)
	} else {
		fmt.Println("QaQ")
	}
	time.Sleep(6 * time.Second)
	_, data = QueryData("hi")
	fmt.Println(data)
}
