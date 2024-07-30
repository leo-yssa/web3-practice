package response

import (
	"fmt"
	"net/http"
	"strings"
	"web3-practice/internal/domain/dto"

	"github.com/gin-gonic/gin"
)

type ResponseStatus int

const (
	OK ResponseStatus = iota + 1
	BAD_REQUEST
	CONFLICT
	UNAUTHORIZED
	INTERNAL_SERVER_ERROR
)

func (r ResponseStatus) GetStatus() string {
	return [...]string{
		"OK",
		"BAD_REQUEST",
		"CONFLICT",
		"UNAUTHORIZED",
		"INTERNAL_SERVER_ERROR",
	}[r-1]
}

func (e ResponseStatus) GetMessage() string {
	return [...]string{
		"OK",
		"BAD_REQUEST",
		"CONFLICT",
		"UNAUTHORIZED",
		"INTERNAL_SERVER_ERROR",
	}[e-1]
}

func ExceptionHandler(ctx *gin.Context, err interface{}) {
	str := fmt.Sprint(err)
	strArr := strings.Split(str, ":")
	key := strArr[0]
	msg := strings.Trim(strArr[1], " ")
	switch key {
	case BAD_REQUEST.GetStatus():
		ctx.JSON(http.StatusBadRequest, BuildReponse(BAD_REQUEST, &dto.Error{
			Message: msg,
		}))
	case CONFLICT.GetStatus():
		ctx.JSON(http.StatusConflict, BuildReponse(CONFLICT, &dto.Error{
			Message: msg,
		}))
	case UNAUTHORIZED.GetStatus():
		ctx.JSON(http.StatusConflict, BuildReponse(UNAUTHORIZED, &dto.Error{
			Message: msg,
		}))
	case INTERNAL_SERVER_ERROR.GetStatus():
		ctx.JSON(http.StatusInternalServerError, BuildReponse(INTERNAL_SERVER_ERROR, &dto.Error{
			Message: msg,
		}))
	}
	ctx.Abort()

}

func Response[T any](ctx *gin.Context, status ResponseStatus, data T) {
	switch status {
	case OK:
		ctx.JSON(http.StatusOK, BuildReponse(OK, data))
	}
}

func Exception(status ResponseStatus, err error) {
	err = fmt.Errorf("%s: %v", status.GetStatus(), err)
	if err != nil {
		panic(err)
	}
}

func BuildReponse[T any](status ResponseStatus, data T) dto.Response[T] {
	return dto.Response[T]{
		Status:  status.GetStatus(),
		Message: status.GetMessage(),
		Data:    data,
	}
}
