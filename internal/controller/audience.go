package controller

import (
	"web3-practice/internal/domain/dto"
	"web3-practice/internal/middleware/response"
	"web3-practice/internal/repository"
	"web3-practice/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func newAudienceController(repo repository.Repository, cache *redis.Client) *audienceController {
	return &audienceController{
		cache: cache,
	}
}

type audienceController struct {
	cache *redis.Client
}

// Audience godoc
// @Tags Audience
// @Summary Google OAuth URL 전송
// @Description Google OAuth URL 전송
// @Produce json
// @Router /audience/google [get]
// @Success 200 {object} dto.Response{data=dto.GoogleAuthCodeURL}
// @Failure 500 {object} dto.Response{data=dto.Error}
func (ac *audienceController) GoogleAuthCodeURL(ctx *gin.Context) {
	uuid := util.GenerateULID("AUD")
	state := util.GetState()
	if err := ac.cache.Set(uuid, state, 0).Err(); err != nil {
		response.Exception(response.INTERNAL_SERVER_ERROR, err)
	}
	response.Response(ctx, response.OK, &dto.GoogleAuthCodeURL{
		Uuid:  uuid,
		State: state,
	})
}

// Audience godoc
// @Tags Audience
// @Summary Google Login
// @Description Google Login
// @Produce json
// @Router /audience/google [post]
// @Param login body dto.GoogleLogin true "login"
// @Success 200 {object} response.Response{data=dto.Jwt}
// @Failure 400 {object} rresponse.Response{data=dto.Error}
// @Failure 401 {object} response.Response{data=dto.Error}
// @Failure 500 {object} response.Response{data=dto.Error}
func (ac *audienceController) GoogleLogin(ctx *gin.Context) {
	var request *dto.GoogleLogin
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Exception(response.BAD_REQUEST, err)
	}
	value, err := ac.cache.Get(request.Uuid).Result()
	if err != nil {
		response.Exception(response.INTERNAL_SERVER_ERROR, err)
	}
	if value != request.State {
		response.Exception(response.BAD_REQUEST, err)
	}
	// ctx.Set("aud", advertisers[0].Id)
	ctx.Next()
}
