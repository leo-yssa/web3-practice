package server

import (
	"errors"
	"strings"

	"github.com/go-redis/redis"
)

func newCache(cfg *Config) (*redis.Client, error) {
	if strings.ToLower(cfg.Cache.Kind) == "redis" {
		c := redis.NewClient(&redis.Options{
			Addr:     cfg.Cache.Host,
			Password: cfg.Cache.Secret,
			DB:       0,
		})
		_, err := c.Ping().Result()
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	return nil, errors.New("not supported cache")
}
