package controller

import (
	"web3-practice/internal/domain/dto"
	"web3-practice/internal/middleware/response"
	"web3-practice/internal/repository"
	"web3-practice/internal/service"
	"web3-practice/pkg/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func newAdvertiserController(repo repository.Repository) *advertiserController {
	return &advertiserController{
		asvc: service.NewAdvertiserService(repo),
	}
}

type advertiserController struct {
	asvc service.AdvertiserService
}

// 광고 godoc
// @Tags Advertiser
// @Summary 광고주 등록
// @Description 광고주 등록
// @Produce json
// @Router /advertiser/sign-up [post]
// @Param advertiser body dto.AdvertiserCreation true "광고주 정보"
// @Success 201 {object} dto.Response
// @Failure 409 {object} dto.Response{data=dto.Error}
// @Failure 500 {object} dto.Response{data=dto.Error}
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
	response.Response(ctx, response.CREATED, "")
}

// 광고 godoc
// @Tags Advertiser
// @Summary 광고주 로그인
// @Description 광고주 로그인
// @Produce json
// @Router /advertiser/sign-in [post]
// @Param advertiser body dto.Advertiser true "광고주 정보"
// @Success 200 {object} response.Response{data=dto.Jwt}
// @Failure 400 {object} response.Response{data=dto.Error}
// @Failure 401 {object} response.Response{data=dto.Error}
// @Failure 406 {object} response.Response{data=dto.Error}
// @Failure 500 {object} response.Response{data=dto.Error}
func (ac *advertiserController) SignIn(ctx *gin.Context) {
	var request *dto.Advertiser
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Exception(response.BAD_REQUEST, err)
	}
	advertisers, err := ac.asvc.FindAdvertiserByEmail(request.Email)
	if err != nil {
		response.Exception(response.INTERNAL_SERVER_ERROR, err)
	}
	if len(advertisers) < 1 {
		response.Exception(response.NOT_ACCEPTABLE, nil)
	}
	if err := util.ComparePassword(advertisers[0].Secret, request.Secret); err != nil {
		response.Exception(response.UNAUTHORIZED, err)
	}
	/* Additional
	 * Gateway Login
	 *  -> Update Advertiser's Gateway Token & Wallet address
	 */
	ctx.Set("aud", advertisers[0].Id)
	ctx.Next()
}
