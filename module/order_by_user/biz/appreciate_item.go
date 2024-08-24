package biz

import (
	"context"
	"errors"
	"main.go/module/order_by_user/model"
)

func (biz *OrderUserBiz) NewAppreciateOrder(ctx context.Context, appreciate *model.UpdateOrder) (int, error) {
	order, err := biz.store.GetOrder(ctx, "id=? and user_id=?", appreciate.Id, appreciate.UserId, "OwnerUser", "OwnerItem")
	if err != nil {
		return 0, err
	}
	if order.Status != model.StatusOrderDone {
		return 0, errors.New("you can't appreciate this item")
	}
	if err := biz.store.UpdateOrder(ctx, appreciate); err != nil {
		return 0, err
	}
	return order.ItemId, nil
}
