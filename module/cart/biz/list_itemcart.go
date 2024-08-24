package biz

import (
	"context"
	"main.go/common"
	"main.go/module/cart/model"
)

type ListItemCartStorage interface {
	ListItemCart(ctx context.Context, userId int, moreKey ...string) (*[]model.CartUser, error)
}
type ListItemCartBiz struct {
	store ListItemCartStorage
}

func NewListItemCartBiz(store ListItemCartStorage) *ListItemCartBiz {
	return &ListItemCartBiz{store: store}
}
func (biz *ListItemCartBiz) NewCartItemCart(ctx context.Context, requester common.Requester) (*[]model.CartUser, error) {
	data, err := biz.store.ListItemCart(ctx, requester.GetUserId(), model.OwnerItem)
	if err != nil {
		return nil, err
	}
	return data, nil
}
