package mysql

import (
	"context"
	"social_todo/internal/services/user/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) CreateUser(ctx context.Context, data *entity.UserCreate) error {
	db := repo.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return errors.WithStack(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return errors.WithStack(err)
	}

	return nil
}
