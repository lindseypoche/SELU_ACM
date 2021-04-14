package redis

import (
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var once sync.Once
var rdb *redis.Client

const (
	defaultHost = "0.0.0.0:6379"
)

func GetClient() *redis.Client {

	uri := os.Getenv("REDIS_HOST")
	if uri == "" {
		uri = defaultHost
	}
	pass := os.Getenv("REDIS_INIT_PASSWORD") // default is empty ""

	once.Do(func() {

		rdb = redis.NewClient(&redis.Options{
			Addr:     uri,
			Password: pass,
			DB:       0,
		})
	})

	return rdb
}

// old

// var rdb *redis.Client

// func GetClient() *redis.Client {

// 	if rdb != nil {
// 		return rdb
// 	}

// 	if rdb.Options().PoolSize < 1 {

// 	}

// 	rdb = redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	return rdb
// }
