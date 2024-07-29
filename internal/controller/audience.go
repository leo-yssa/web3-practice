package controller

import (
	"net/http"
	"web3-practice/internal/domain/dto"
	"web3-practice/internal/middleware/exception"
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

func (ac *audienceController) GoogleAuthCodeURL(ctx *gin.Context) {
	uuid := util.GenerateULID("AUD")
	state := util.GetState()
	if err := ac.cache.Set(uuid, state, 0).Err(); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, &dto.GoogleAuthCodeURL{
		Uuid:  uuid,
		State: state,
	})
}

func (ac *audienceController) GoogleLogin(ctx *gin.Context) {
	var request *dto.GoogleLogin
	if err := ctx.ShouldBindJSON(&request); err != nil {
		panic(exception.BAD_REQUEST)
	}
	value, err := ac.cache.Get(request.Uuid).Result()
	if err != nil {
		panic(err)
	}
	if value != request.State {
		panic(exception.BAD_REQUEST)
	}
	ctx.JSON(http.StatusOK, "")
}
