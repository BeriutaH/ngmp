package config

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	// DBDefault mysql默认连接
	DBDefault *gorm.DB
	// RedisDefault redis默认连接
	RedisDefault *redis.Client
	TimeString   = "2006-01-02 15:04:05"
)
