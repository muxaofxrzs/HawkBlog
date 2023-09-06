package redis

import (
	"context"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 2cb677a71339b35127a74f83095f18be7858e549
	"github.com/go-redis/redis/v8"
)

var ClientRe *redis.Client

<<<<<<< HEAD
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
=======
func CreateRedisClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "1.94.27.198:6379", // Redis 服务器地址和端口号
		Password: "hawk123",               // Redis 访问密码，如果没有设置密码，可以留空
		DB:       0,                // Redis 数据库索引，默认为 0
	})
	//通过 redis.NewClient() 函数创建了一个 Redis 客户端连接。需要提供 Redis 服务器的地址和端口号、密码（如果有设置的话）以及数据库索引。
	//Ping() 方法用于测试与 Redis 的连接是否正常
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		// 处理 Redis 连接错误
		panic(err)
	}

	ClientRe = client
}
>>>>>>> 2cb677a71339b35127a74f83095f18be7858e549
