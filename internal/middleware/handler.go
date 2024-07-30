package middleware

import (
	"web3-practice/internal/config"
	"web3-practice/internal/controller"
	"web3-practice/internal/middleware/response"
	"web3-practice/internal/service"
	"web3-practice/pkg/swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewGinHandler(rdb *gorm.DB, ctrl controller.Controller, cfg *config.Config) *gin.Engine {
	r := gin.New()
	swagger.Initialize(r)
	ts := service.NewTokenService(cfg)
	api := r.Group("", defaultHandler(rdb))
	{
		advertiser := api.Group("/advertiser")
		advertiser.POST("/sign-up", ctrl.SignUp)
		advertiser.POST("/sign-in", ctrl.SignIn, issueToken(ts))
		advertiser.POST("/refresh", verifyRefreshToken(ts), issueToken(ts))
		audience := api.Group("/audience")
		audience.GET("/google", ctrl.GoogleAuthCodeURL)
		audience.POST("/google", ctrl.GoogleLogin, issueToken(ts))
		audience.POST("/refresh", verifyRefreshToken(ts), issueToken(ts))
	}
	return r
}

func verifyRefreshToken(ts service.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		aud := ts.VerifyRefreshToken(token)
		if "" == aud {
			response.Exception(response.UNAUTHORIZED, nil)
		}
		ctx.Set("aud", aud)
		ctx.Next()
	}
}

func verifyAccessToken(ts service.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		aud := ts.VerifyAccessToken(token)
		if "" == aud {
			response.Exception(response.UNAUTHORIZED, nil)
		}
		ctx.Set("aud", aud)
		ctx.Next()
	}
}

func issueToken(ts service.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		aud, _ := ctx.Keys["aud"].(string)
		token, err := ts.Issue(aud)
		if err != nil {
			response.Exception(response.INTERNAL_SERVER_ERROR, err)
		}
		response.Response(ctx, response.OK, token)
	}
}

func defaultHandler(rdb *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tx := rdb.Begin()
		defer func() {
			if err := recover(); err != nil {
				tx.Rollback()
				response.ExceptionHandler(ctx, err)
			} else {
				tx.Commit()
			}
		}()
		ctx.Set("tx", tx)
		ctx.Next()
	}
}
