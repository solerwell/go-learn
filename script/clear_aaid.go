package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

var ctx = context.Background()

const AAID_KEY_PREFIX = "AAID:"

func main1() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "some password",
		DB:       0,
	})
	defer redisClient.Close()

	keyPrefix := AAID_KEY_PREFIX
	if len(os.Args) > 1 {
		keyPrefix = os.Args[1]
	}

	var foundedRecordCount int = 0
	iter := redisClient.Scan(ctx, 0, keyPrefix, 0).Iterator()
	fmt.Printf("YOUR SEARCH PATTERN= %s\n", keyPrefix)
	for iter.Next(ctx) {
		fmt.Printf("Deleted= %s\n", iter.Val())
		redisClient.Del(ctx, iter.Val())
		foundedRecordCount++
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("Deleted Count %d\n", foundedRecordCount)

}
