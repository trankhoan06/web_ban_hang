package biz

import (
	"context"
)

func (biz *FollowBiz) NewGetAmountFollow(ctx context.Context, userId int) (int, error) {
	amount, err := biz.store.GetAmountFollow(ctx, "user_id", userId)
	if err != nil {
		return 0, err
	}
	return amount, nil

}
