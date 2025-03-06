package api

import (
	"fmt"
	"net/http"
	core "social_todo/internal/common"
	common "social_todo/internal/common/error"
	"social_todo/internal/services/items/entity"

	"github.com/gin-gonic/gin"
)

func (service *itemService) CreateItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var itemData entity.TodoItemCreation

		if err := ctx.ShouldBind(&itemData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": fmt.Sprintf("Bind itemdata failed %+v", err),
			}) //xem lai cach tra loi
			return
		}

		requester := core.GetRequester(ctx)
		itemData.UserID = core.GetRequesterID(requester)

		if err := service.itemBiz.CreateNewItem(ctx, &itemData); err != nil {
			common.ErrorResponse(ctx, err)
			return
		}

		common.SuccessResponse(ctx, common.NewDataResponse(itemData.ID))
	}
}
