package mysql

import (
	"context"
	"fmt"
	"social_todo/internal/services/items/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) GetItem(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) {
	var data entity.TodoItem

	if err := repo.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Record not found")
		}

		return nil, errors.WithStack(err)
	}

	return &data, nil
}
