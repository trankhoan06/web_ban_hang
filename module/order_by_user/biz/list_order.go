package biz

import (
	"context"
	"main.go/module/order_by_user/model"
)

func (biz *OrderUserBiz) NewListOrder(ctx context.Context, userId int) (*[]model.Order, error) {
	order, err := biz.store.ListOrder(ctx, userId, "user_id", []string{"OwnerUser"}, "OwnerItem")
	if err != nil {
		return nil, err
	}
	return order, nil
}
