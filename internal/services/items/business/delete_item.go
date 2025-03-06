package business

import "context"

func (biz *itemBusiness) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	return biz.itemRepo.DeleteItem(ctx, cond)
}
