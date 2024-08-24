package biz

import (
	"context"
	"main.go/module/cart/model"
)

type DeletedAmountItemStorage interface {
	DeletedItemCart(ctx context.Context, itemId, userId int) error
	FindItem(ctx context.Context, itemId, userId int) (*model.CartUser, error)
}

type CartDeletedAmountItemBiz struct {
	store DeletedAmountItemStorage
}

func NewCartDeletedAmountItemBiz(store DeletedAmountItemStorage) *CartDeletedAmountItemBiz {
	return &CartDeletedAmountItemBiz{
		store: store,
	}
}
func (biz *CartDeletedAmountItemBiz) NewCartDeletedAmountItem(ctx context.Context, itemId, userId int) error {
	if _, err := biz.store.FindItem(ctx, itemId, userId); err != nil {
		return err
	}
	if err := biz.store.DeletedItemCart(ctx, itemId, userId); err != nil {
		return err
	}
	return nil
}
