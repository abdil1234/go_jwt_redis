package db

import "github.com/go-redis/redis"

var (
	// ReCon is the connection handle
	// for the database redis
	ReCon *redis.Client
)
