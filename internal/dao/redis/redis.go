package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ClientRe *redis.Client

func NewRe() {
	client := redis.NewClient(&redis.Options{
		Addr:     "1.94.27.198:6379",
		Password: "hawk123",
		DB:       0,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		fmt.Println("redis连接失败")
	}
	ClientRe = client
}
func CloseRe() {
	ClientRe.Del(context.Background(), "library")
	_ = ClientRe.Close()
}
