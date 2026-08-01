package config

import "github.com/redis/go-redis/v9"

func NewCached(cnf *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cnf.Redis.Host,
		Password: cnf.Redis.Password,
		DB:       cnf.Redis.DB,
	})
}
