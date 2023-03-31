package db

import (
	"context"
	"fmt"
	//"os"

	"github.com/redis/go-redis/v9"
)

func Cache() (cache *redis.Client){
	ctx := context.Background()

	cache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password:  "",
		DB:       0,
	})

	redis_error := cache.Set(ctx, "teste", "teste", 0).Err()
	
	if redis_error != nil {
		fmt.Print()
	}

	err := cache.Del(ctx, "teste")

	if err != nil{
		fmt.Print(err)
	}

	return cache
}
