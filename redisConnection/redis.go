package redisConnection

import (
	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func init() {
	RedisPool = NewPool()
}

func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   10000,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, nil
		},
	}
}

func GetRedisConn() redis.Conn {
	return RedisPool.Get()
}
