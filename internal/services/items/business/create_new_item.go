package business

import (
	"context"
	"social_todo/internal/services/items/entity"
)

func (biz *itemBusiness) CreateNewItem(ctx context.Context, data *entity.TodoItemCreation) error {
	return biz.itemRepo.CreateItem(ctx, data)
}
