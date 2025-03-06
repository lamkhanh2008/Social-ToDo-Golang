package api

import (
	"net/http"
	core "social_todo/internal/common"
	common "social_todo/internal/common/error"

	"github.com/gin-gonic/gin"
)

func (service *itemService) GetItemByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := core.UIDFromString(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": err,
			})
		}

		data, err := service.itemBiz.GetItemById(ctx, int(id.GetLocalID()))
		if err != nil {
			common.ErrorResponse(ctx, err)
			return
		}

		common.SuccessResponse(ctx, common.NewDataResponse(data))
	}
}
