package db

import (
	"context"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

func NewRedis() *redis.Client {
	// Initialize a Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Default DB
	})
	return rdb
}

//func ()Set() {
//	// Example: Setting a value
//	err := rdb.Set(ctx, "key", "value", 0).Err()
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func Get() {
//	// Example: Getting a value
//	val, err := rdb.Get(ctx, "key").Result()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("key:", val)
//}
//
//func Close() {
//	// Don't forget to close the client when you're done with it
//	rdb.Close()
//}
