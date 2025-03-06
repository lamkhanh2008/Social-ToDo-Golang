package mysql

import (
	"context"
	"social_todo/internal/services/items/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := "Deleted"

	if err := repo.db.Table(entity.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": deletedStatus,
	}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
