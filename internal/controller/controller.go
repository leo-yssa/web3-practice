package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Controller interface {
	SignUp(ctx *gin.Context)
	AuthCodeURL(ctx *gin.Context)
}

func NewController(rdb *gorm.DB, cache *redis.Client) Controller {
	return &controller{
		advertiserController: newAdvertiserController(rdb),
		audienceController:   newAudienceController(rdb, cache),
	}
}

type controller struct {
	*advertiserController
	*audienceController
}
