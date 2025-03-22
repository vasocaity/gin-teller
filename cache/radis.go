package cache

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func Init() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_HOST"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	// ใช้ Context กับ Redis
	err := rdb.Set(ctx, "name", "Golang", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Value:", val)
}
