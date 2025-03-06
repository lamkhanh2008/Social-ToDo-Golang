package business

import (
	"context"
	core "social_todo/internal/common"
	common "social_todo/internal/common/error"
	"social_todo/internal/common/utils"
	"social_todo/internal/services/user/entity"
)

func (biz *userBusiness) Login(ctx context.Context, data *entity.UserLogin) error {
	user, _ := biz.repo.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return common.ErrBadRequest.
			WithError(entity.ErrEmailOrPasswordInvalid.Error())
	}

	err := utils.CheckPassword(user.Password, core.HashPasswordFormat, user.Salt, data.Password)
	if err != nil {
		return common.ErrBadRequest.WithError(entity.ErrEmailOrPasswordInvalid.Error()).WithDebug(err.Error())
	}

	// salt, err := util.RandomString(50)
	// if err != nil {
	// 	return core.ErrInternalServerError.
	// 		WithError(entity.ErrCannotCreateUser.Error()).
	// 		WithDebug(err.Error())
	// }

	// data.Password, err = util.HashPassword(common.HashPasswordFormat, salt, data.Password)
	// if err != nil {
	// 	return core.ErrInternalServerError.
	// 		WithError(entity.ErrCannotCreateUser.Error()).
	// 		WithDebug(err.Error())
	// }

	// data.Salt = salt
	// data.Role = entity.RoleUser.String()

	// if err := biz.repo.CreateUser(ctx, data); err != nil {
	// 	return core.ErrInternalServerError.
	// 		WithError(entity.ErrCannotCreateUser.Error()).
	// 		WithDebug(err.Error())
	// }

	// return nil
}
