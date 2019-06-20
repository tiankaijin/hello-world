package main

import (
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func initRedis() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	return
}
func GetConn() redis.Conn {
	return pool.Get()
}
func PutConn(conn redis.Conn) {
	conn.Close()
}
