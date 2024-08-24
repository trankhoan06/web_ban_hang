package biz

import (
	"context"
	"main.go/module/item/model"
)

type ListOwnItemStorage interface {
	ListOwnItem(ctx context.Context, userId int, moreKey ...string) (*[]model.TodoList, error)
}
type ListOwnItemBiz struct {
	store ListOwnItemStorage
}

func NewListOwnItemBiz(store ListOwnItemStorage) *ListOwnItemBiz {
	return &ListOwnItemBiz{store: store}
}
func (biz *ListOwnItemBiz) ListNewItem(ctx context.Context, userId int) (*[]model.TodoList, error) {
	result, err := biz.store.ListOwnItem(ctx, userId, "Owner")
	if err != nil {
		return nil, err
	}
	return result, nil
}
