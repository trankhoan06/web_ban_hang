package biz

import (
	"context"
	"main.go/module/item/model"
)

type CategoryItemStorage interface {
	ListItemCategory(ctx context.Context, category string) (*[]model.TodoList, error)
}
type CategoryItemBiz struct {
	store CategoryItemStorage
}

func NewCategoryItemBiz(store CategoryItemStorage) *CategoryItemBiz {
	return &CategoryItemBiz{store: store}
}
func (biz *CategoryItemBiz) NewListItemCategory(ctx context.Context, category string) (*[]model.TodoList, error) {
	result, err := biz.store.ListItemCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return result, nil
}
