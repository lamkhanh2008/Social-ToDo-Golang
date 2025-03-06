package business

import (
	"context"
	"social_todo/internal/services/items/entity"
)

type ItemRepository interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error)
	CreateItem(ctx context.Context, data *entity.TodoItemCreation) error
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type itemBusiness struct {
	itemRepo ItemRepository
}

func NewItemBusiness(itemRepo ItemRepository) *itemBusiness {
	return &itemBusiness{
		itemRepo: itemRepo,
	}
}
