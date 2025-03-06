package api

import (
	"fmt"
	"net/http"
	core "social_todo/internal/common"
	common "social_todo/internal/common/error"

	"github.com/gin-gonic/gin"
)

func (service *itemService) DeleteItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := core.UIDFromString(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": fmt.Sprintf("UserId error %+v", err),
			})
		}
		if err := service.itemBiz.DeleteItemByID(ctx, int(id.GetLocalID())); err != nil {
			common.ErrorResponse(ctx, err)
			return
		}
		common.SuccessResponse(ctx, common.NewDataResponse(true))
	}
}
