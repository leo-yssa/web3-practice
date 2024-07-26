package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Controller interface {
	SignUp(ctx *gin.Context)
}

func NewController(rdb *gorm.DB, cache *redis.Client) Controller {
	return &controller{
		advertiserController: newAdvertiserController(rdb),
	}
}

type controller struct {
	*advertiserController
}
