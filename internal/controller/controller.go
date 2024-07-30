package controller

import (
	"web3-practice/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Controller interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	GoogleAuthCodeURL(ctx *gin.Context)
	GoogleLogin(ctx *gin.Context)
}

func NewController(repo repository.Repository, cache *redis.Client) Controller {
	return &controller{
		advertiserController: newAdvertiserController(repo),
		audienceController:   newAudienceController(repo, cache),
	}
}

type controller struct {
	*advertiserController
	*audienceController
}
