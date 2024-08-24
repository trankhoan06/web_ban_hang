package biz

import (
	"context"
	"errors"
	"main.go/module/order_by_user/model"
)

func (biz *OrderUserBiz) NewGetOrder(ctx context.Context, id int, userId int) (*model.Order, error) {
	order, err := biz.store.GetOrder(ctx, "id = ? and user_id=?", id, userId, "OwnerUser", "OwnerItem")
	if err != nil {
		return nil, err
	}
	if order.Status == model.StatusOrderCancel {
		return nil, errors.New("this order has been cancel")
	}
	if order.Status == model.StatusOrderDone {
		return nil, errors.New("this order has been complete")
	}
	return order, nil
}
