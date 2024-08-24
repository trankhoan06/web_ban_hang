package biz

import (
	"context"
)

func (biz *FollowBiz) NewGetAmountUserFollow(ctx context.Context, userId int) (int, error) {
	amount, err := biz.store.GetAmountFollow(ctx, "by_user_id", userId)
	if err != nil {
		return 0, err
	}
	return amount, nil

}
