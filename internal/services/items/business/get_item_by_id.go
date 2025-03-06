package business

import (
	"context"
	"social_todo/internal/services/items/entity"
)

func (biz *itemBusiness) GetItemById(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) {
	return biz.itemRepo.GetItem(ctx, cond)
}
