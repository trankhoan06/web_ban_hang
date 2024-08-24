package biz

import (
	"context"
	"main.go/module/order_by_user/model"
)

func (biz *OrderUserBiz) NewListOrderSell(ctx context.Context, sellId int) (*[]model.Order, error) {
	orders, err := biz.store.ListOrder(ctx, sellId, "sell_id", []string{"OwnerUser"}, "OwnerItem")
	if err != nil {
		return nil, err
	}
	return orders, nil
}
