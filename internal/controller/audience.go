package controller

import (
	"web3-practice/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func newAudienceController(rdb *gorm.DB, cache *redis.Client) *audienceController {
	return &audienceController{
		cache: cache,
	}
}

type audienceController struct {
	cache *redis.Client
}

func (ac *audienceController) AuthCodeURL(ctx *gin.Context) {
	uuid := util.GenerateUUID()
	state := util.GetState()
	txp := ac.cache.TxPipeline()
	txp.Save()

}
