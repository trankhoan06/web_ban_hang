package biz

import (
	"context"
	"errors"
	"main.go/module/order_by_user/model"
)

func (biz *OrderUserBiz) NewCancelOrder(ctx context.Context, id, userId int) error {
	item, err := biz.store.GetOrder(ctx, "id = ? and user_id=?", id, userId, "OwnerUser", "OwnerItem")
	if err != nil {
		return err
	}
	if item.Status != model.StatusOrderPrepare {
		return errors.New("you can't cancel this order")
	}
	if err := biz.store.UpdateStatusOrder(ctx, id, model.StatusOrderCancel); err != nil {
		return err
	}
	return nil
}
