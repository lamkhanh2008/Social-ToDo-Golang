package mysql

import (
	"context"
	common "social_todo/internal/common/error"
	"social_todo/internal/services/user/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) FindUser(ctx context.Context, conds map[string]interface{}, moreInfo ...string) (*entity.User, error) {
	db := repo.db.Table(entity.User{}.TableName())
	for _, value := range moreInfo {
		db = db.Preload(value)
	}

	var user entity.User
	if err := db.Where(conds).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &user, nil
}

func (repo *mysqlRepo) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	db := repo.db.Table(entity.User{}.TableName())

	var user entity.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &user, nil
}
