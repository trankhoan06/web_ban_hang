package biz

import (
	"context"
	"errors"
	"main.go/module/order_by_user/model"
)

func (biz *OrderUserBiz) NewUpdateOrder(ctx context.Context, data *model.UpdateOrder) error {
	order, err := biz.store.GetOrder(ctx, "id = ? and user_id=?", data.Id, data.UserId, "OwnerUser", "OwnerItem")
	if err != nil {
		return err
	}
	if order.Status != model.StatusOrderPrepare {
		return errors.New("you can't update infomation")
	}
	if err := biz.store.UpdateOrder(ctx, data); err != nil {
		return err
	}
	return nil
}
