package mysql

import (
	"context"
	"social_todo/internal/services/items/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) CreateItem(ctx context.Context, data *entity.TodoItemCreation) error {
	if err := repo.db.Table(entity.TodoItem{}.TableName()).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
