package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := client.Context()

	val, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}
