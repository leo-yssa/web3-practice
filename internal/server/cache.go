package server

import (
	"errors"
	"strings"
	"web3-practice/internal/config"

	"github.com/go-redis/redis"
)

func newCache(cfg *config.Config) (*redis.Client, error) {
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
