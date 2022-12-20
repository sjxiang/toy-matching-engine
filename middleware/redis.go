package middleware

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"

	"toy-matching-engine/pkg/log"
)


var RedisClient *redis.Client


func Init()  {
	addr := viper.GetString("redis.addr")
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: "",  
		DB: 0,        
	})

	if _, err := RedisClient.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	log.Printf("Connected to redis: %s", addr)
}