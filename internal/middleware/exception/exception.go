package exception

import (
	"net/http"
	"web3-practice/internal/domain/dto"

	"github.com/gin-gonic/gin"
)

type ExceptionStatus int

const (
	INTERNAL_SERVER_ERROR ExceptionStatus = iota + 1
	BAD_REQUEST
	CONFLICT
)

func (e ExceptionStatus) GetMessage() *dto.Response {
	return [...]*dto.Response{
		{
			Message: "INTERNAL_SERVER_ERROR",
		},
		{
			Message: "BAD_REQUEST",
		},
		{
			Message: "CONFLICT",
		},
	}[e-1]
}

func ExceptionHandler(ctx *gin.Context, err interface{}) {
	switch err {
	case BAD_REQUEST:
		ctx.JSON(http.StatusBadRequest, BAD_REQUEST.GetMessage())
	case CONFLICT:
		ctx.JSON(http.StatusConflict, CONFLICT.GetMessage())
	default:
		ctx.JSON(http.StatusInternalServerError, INTERNAL_SERVER_ERROR.GetMessage())
	}
	ctx.Abort()
}
