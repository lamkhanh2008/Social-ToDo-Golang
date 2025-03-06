package api

import (
	"context"
	"social_todo/internal/services/items/entity"
)

type ItemBusiness interface {
	CreateNewItem(ctx context.Context, data *entity.TodoItemCreation) error
	DeleteItemByID(ctx context.Context, id int) error
	GetItemById(ctx context.Context, id int) (*entity.TodoItem, error)
}

type itemService struct {
	itemBiz ItemBusiness
}

func NewItemService(itemBiz ItemBusiness) *itemService {
	return &itemService{
		itemBiz: itemBiz,
	}
}
