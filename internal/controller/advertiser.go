package controller

import (
	"web3-practice/internal/domain/dto"
	"web3-practice/internal/middleware/response"
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
		response.Exception(response.BAD_REQUEST, err)
	}
	advertisers, err := ac.asvc.FindAdvertiserByEmail(request.Email)
	if err != nil {
		response.Exception(response.INTERNAL_SERVER_ERROR, err)
	}
	if len(advertisers) > 0 {
		response.Exception(response.CONFLICT, err)
	}
	tx, _ := ctx.Keys["tx"].(*gorm.DB)
	if err := ac.asvc.CreateAdvertiser(request, tx); err != nil {
		response.Exception(response.INTERNAL_SERVER_ERROR, err)
	}
	response.Response(ctx, response.OK, "")
}
