package biz

import (
	"context"
	"main.go/module/item/model"
)

type ListItemStorage interface {
	ListItem(ctx context.Context, paging *model.Paging, filter *model.Filter, moreKey ...string) (*[]model.TodoList, error)
}
type ListItemBiz struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *ListItemBiz {
	return &ListItemBiz{store: store}
}
func (biz *ListItemBiz) ListNewItem(ctx context.Context, paging *model.Paging, filter *model.Filter) (*[]model.TodoList, error) {
	result, err := biz.store.ListItem(ctx, paging, filter, "Owner")
	paging.Process()
	if err != nil {
		return nil, err
	}
	return result, nil
}
