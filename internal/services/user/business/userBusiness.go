package business

import (
	"context"
	"social_todo/internal/services/user/entity"
)

type UserRepository interface {
	FindUser(ctx context.Context, conds map[string]interface{}, moreInfo ...string) (*entity.User, error)
	CreateUser(ctx context.Context, data *entity.UserCreate) error
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
}

type userBusiness struct {
	repo UserRepository
}

func NewUserBusiness(userRepo UserRepository) *userBusiness {
	return &userBusiness{repo: userRepo}
}
