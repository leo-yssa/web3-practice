package controller

import (
	"net/http"
	"web3-practice/internal/domain/dto"
	"web3-practice/internal/middleware/exception"
	"web3-practice/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func newAdvertiserController(rdb *gorm.DB) *advertiserController {
	return &advertiserController{
		asvc: service.NewAdvertiserService(rdb),
	}
}

type advertiserController struct {
	asvc service.AdvertiserService
}

func (ac *advertiserController) SignUp(ctx *gin.Context) {
	var request *dto.AdvertiserCreation
	if err := ctx.ShouldBindJSON(&request); err != nil {
		panic(exception.BAD_REQUEST)
	}
	advertisers, err := ac.asvc.FindAdvertiserByEmail(request.Email)
	if err != nil {
		panic(err)
	}
	if len(advertisers) > 0 {
		panic(exception.CONFLICT)
	}
	tx, _ := ctx.Keys["tx"].(*gorm.DB)
	if err := ac.asvc.CreateAdvertiser(request, tx); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, &dto.Response{
		Message: "OK",
	})
}
