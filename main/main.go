package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	"time"
)

func main() {
	clusterNodes := []string{
		"192.168.0.231:6380",
		"192.168.0.231:6381",
		"192.168.0.232:6382",
		"192.168.0.232:6383",
		"192.168.0.233:6384",
		"192.168.0.233:6385",
	}
	// 创建一个 Redis Cluster 客户端
	opts := &redis.ClusterOptions{
		Addrs:    clusterNodes,
		Password: "123123", // 如果你的 Redis 需要密码，请在这里设置
	}

	// 创建Redis客户端
	client := redis.NewClusterClient(opts)

	ctx := context.Background()

	// 设置键值
	// specified duration is 1µs, but minimal supported value is 1ms - truncating to 1ms
	//应该是duration 类型。这个时间单位不是s。
	//Go语言文档中指出最小的时间间隔是1纳秒（1ns）,默认纳秒单位
	err := client.Set(ctx, "test_key", "test_value", 5*60*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// 获取键值
	val, err := client.Get(ctx, "test_key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key", val) // 输出: key value

	// 关闭客户端连接
	err = client.Close()
	if err != nil {
		panic(err)
	}
}
