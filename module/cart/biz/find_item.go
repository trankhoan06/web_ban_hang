package biz

import (
	"context"
	"main.go/common"
	modelcart "main.go/module/cart/model"
)

type FindItemStorage interface {
	FindItem(ctx context.Context, itemId, userId int) (*modelcart.CartUser, error)
}
type FindItemBiz struct {
	store FindItemStorage
}

func NewFindItemBiz(store FindItemStorage) *FindItemBiz {
	return &FindItemBiz{
		store: store,
	}
}
func (biz *FindItemBiz) NewFindItem(ctx context.Context, itemId int, requester common.Requester) (*modelcart.CartUser, error) {
	//nếu item đã có trong giỏ thì thêm số lượng
	cart, err := biz.store.FindItem(ctx, itemId, requester.GetUserId())
	if err != nil {
		return nil, err
	}
	return cart, nil
}
