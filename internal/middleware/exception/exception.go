package exception

import (
	"net/http"
	"web3-practice/internal/domain/dto"

	"github.com/gin-gonic/gin"
)

type ExceptionStatus int

const (
	BAD_REQUEST ExceptionStatus = iota + 1
	CONFLICT
	UNAUTHORIZED
)

func (e ExceptionStatus) GetMessage() *dto.Response {
	return [...]*dto.Response{
		{
			Message: "BAD_REQUEST",
		},
		{
			Message: "CONFLICT",
		},
		{
			Message: "UNAUTHORIZED",
		},
	}[e-1]
}

func ExceptionHandler(ctx *gin.Context, err interface{}) {
	switch err {
	case BAD_REQUEST:
		ctx.JSON(http.StatusBadRequest, BAD_REQUEST.GetMessage())
	case CONFLICT:
		ctx.JSON(http.StatusConflict, CONFLICT.GetMessage())
	case UNAUTHORIZED:
		ctx.JSON(http.StatusConflict, CONFLICT.GetMessage())
	default:
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.Abort()
}
