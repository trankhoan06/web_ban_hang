package biz

import (
	"context"
	modelOrder "main.go/module/order_by_user/model"
)

func (biz *OrderUserBiz) GetOrderSell(ctx context.Context, sellId, id int) (*modelOrder.Order, error) {
	order, err := biz.store.GetOrder(ctx, "id = ? and sell_id=?", id, sellId, "OwnerUser", "OwnerItem")
	if err != nil {
		return nil, err
	}
	return order, nil
}
