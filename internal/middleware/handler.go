package middleware

import (
	"web3-practice/internal/controller"
	"web3-practice/internal/middleware/exception"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewGinHandler(rdb *gorm.DB, ctrl controller.Controller) *gin.Engine {
	r := gin.New()
	api := r.Group("", defaultHandler(rdb))
	{
		api.POST("/advertiser", ctrl.SignUp)
		api.GET("/audience/google", ctrl.GoogleAuthCodeURL)
		api.POST("/audience/google", ctrl.GoogleLogin)
	}
	return r
}

func defaultHandler(rdb *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tx := rdb.Begin()
		defer func() {
			if err := recover(); err != nil {
				tx.Rollback()
				exception.ExceptionHandler(ctx, err)
			} else {
				tx.Commit()
			}
		}()
		ctx.Set("tx", tx)
		ctx.Next()
	}
}
