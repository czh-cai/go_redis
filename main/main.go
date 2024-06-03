package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	// 创建Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.231:6379", // Redis服务器地址
		Password: "123456",             // 密码，没有则留空
		DB:       0,                    // 使用默认DB
	})

	ctx := context.Background()

	// 设置键值
	err := rdb.Set(ctx, "test_key", "test_value", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取键值
	val, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key", val) // 输出: key value

	// 关闭客户端连接
	err = rdb.Close()
	if err != nil {
		panic(err)
	}
}
