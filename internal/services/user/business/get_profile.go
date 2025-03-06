package business

import (
	"context"
	core "social_todo/internal/common"
	common "social_todo/internal/common/error"
	"social_todo/internal/services/user/entity"
)

func (userBiz userBusiness) GetProfile(ctx context.Context) (*entity.User, error) {
	requester := core.GetRequester(ctx)
	requesterID := core.GetRequesterID(requester)

	user, err := userBiz.repo.GetUserByID(ctx, requesterID)
	if err != nil {
		return nil, common.ErrUnauthorized.WithError(entity.ErrCannotGetUser.Error()).WithDebug(err.Error())
	}

	return user, nil
}
