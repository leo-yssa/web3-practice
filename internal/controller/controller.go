package controller

import (
	"web3-practice/internal/config"
	"web3-practice/internal/repository"
	"web3-practice/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Controller interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	GoogleAuthCodeURL(ctx *gin.Context)
	GoogleLogin(ctx *gin.Context)
}

func NewController(
	repo repository.Repository,
	cache *redis.Client,
	cfg *config.Config,
) Controller {
	return &controller{
		advertiserController: newAdvertiserController(repo),
		audienceController: newAudienceController(
			repo,
			cache,
			service.NewGoogleService(cfg, service.LOGIN),
		),
	}
}

type controller struct {
	*advertiserController
	*audienceController
}
