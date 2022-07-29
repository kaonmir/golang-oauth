package db

import (
	"github.com/go-redis/redis"
	"github.com/kaonmir/OAuth/config"
	_ "github.com/mattn/go-sqlite3"
)

var client *redis.Client

func Init() {
	env := config.Env()

	// New Redis Client with port 6379

	_client := redis.NewClient(&redis.Options{
		Addr:     env.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client = _client
}

func GetDB() *redis.Client {
	return client
}
