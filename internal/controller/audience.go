package controller

import (
	"fmt"
	"web3-practice/internal/domain/dto"
	"web3-practice/internal/middleware/response"
	"web3-practice/internal/repository"
	"web3-practice/internal/service"
	"web3-practice/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func newAudienceController(
	repo repository.Repository,
	cache *redis.Client,
	google service.GoogleService,
) *audienceController {
	return &audienceController{
		cache: cache,
		gsrvc: google,
	}
}

type audienceController struct {
	cache *redis.Client
	gsrvc service.GoogleService
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
	url := ac.gsrvc.AuthCodeURL(state)
	response.Response(ctx, response.OK, &dto.GoogleAuthCodeURL{
		GoogleAuthState: &dto.GoogleAuthState{
			Uuid:  uuid,
			State: state,
		},
		Url: url,
	})
}

// Audience godoc
// @Tags Audience
// @Summary Google Login
// @Description Google Login
// @Produce json
// @Router /audience/google [post]
// @Param login body dto.GoogleLogin true "login"
// @Success 200 {object} dto.Response{data=dto.Jwt}
// @Failure 400 {object} dto.Response{data=dto.Error}
// @Failure 401 {object} dto.Response{data=dto.Error}
// @Failure 500 {object} dto.Response{data=dto.Error}
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
	token, err := ac.gsrvc.Exchange(ctx, request.Code)
	if err != nil {
		response.Exception(response.UNAUTHORIZED, err)
	}
	audience, err := ac.gsrvc.UserInfo(token.AccessToken)
	if err != nil {
		response.Exception(response.INTERNAL_SERVER_ERROR, err)
	}
	fmt.Println(audience)
	// ctx.Set("aud", advertisers[0].Id)
	ctx.Next()
}
