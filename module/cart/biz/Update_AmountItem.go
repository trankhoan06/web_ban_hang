package biz

import (
	"context"
	"main.go/module/cart/model"
)

type UpdateAmountItemStorage interface {
	UpdateItemCart(ctx context.Context, update *model.CartUpdateUser) error
	FindItem(ctx context.Context, itemId, userId int) (*model.CartUser, error)
}

type CartUpdateAmountItemBiz struct {
	store UpdateAmountItemStorage
}

func NewCartUpdateAmountItemBiz(store UpdateAmountItemStorage) *CartUpdateAmountItemBiz {
	return &CartUpdateAmountItemBiz{
		store: store,
	}
}
func (biz *CartUpdateAmountItemBiz) NewCartUpdateAmountItem(ctx context.Context, update *model.CartUpdateUser) error {
	if _, err := biz.store.FindItem(ctx, update.ItemId, update.UserId); err != nil {
		return err
	}
	if err := biz.store.UpdateItemCart(ctx, update); err != nil {
		return err
	}
	return nil
}
