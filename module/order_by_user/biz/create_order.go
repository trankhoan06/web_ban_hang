package biz

import (
	"context"
	"errors"
	"main.go/module/order_by_user/model"
)

func (biz *OrderBiz) CreateOrder(ctx context.Context, data *model.CreateOrder) error {
	item, err := biz.store1.GetItem(ctx, map[string]interface{}{"id": data.ItemId})
	if err != nil {
		return err
	}
	if item.AmountItem == 0 {
		return errors.New("item has been sold out")
	}
	data.SellId = item.UserId
	if err := biz.store.CreateOrder(ctx, data); err != nil {
		return err
	}
	return nil
}
