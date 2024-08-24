package biz

import (
	"context"
	"errors"
	"main.go/module/order_by_user/model"
)

func (biz *OrderUserBiz) UpdateStatusSell(ctx context.Context, id, sellId int, order model.StatusOrder) (*model.Order, error) {
	item, err := biz.store.GetOrder(ctx, "id = ? and sell_id = ?", id, sellId, "OwnerUser", "OwnerItem")
	if err != nil {
		return nil, err
	}
	if order == model.StatusOrderCancel {
		return nil, errors.New("The seller does not have the right to cancel")
	}
	if item.Status == model.StatusOrderCancel || item.Status == model.StatusOrderDone {
		return nil, errors.New("this item has been cancelled or completed")
	}
	if err := biz.store.UpdateStatusOrder(ctx, id, order); err != nil {
		return nil, err
	}
	return item, nil
}
