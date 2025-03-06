package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, err error) {
	// if cErr, ok := err.(StatusC)
	ctx.JSON(
		http.StatusInternalServerError,
		ErrInternalServerError.WithError(err.Error()),
	)

}

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Extra  interface{} `json:"extra,omitempty"`
}

func NewResponse(data, paging, extra interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging, Extra: extra}
}

func NewDataResponse(data interface{}) *successResponse {
	return NewResponse(data, nil, nil)
}

func SuccessResponse(ctx *gin.Context, data *successResponse) {
	ctx.JSON(http.StatusOK, data)
}
