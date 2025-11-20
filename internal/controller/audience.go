package controller

import (
	"net/http"
	"web3-practice/internal/domain/dto"
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
	if err := ac.cache.Set(uuid, state, 0).Err(); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, &dto.Response{
		Message: "OK",
	})
}
