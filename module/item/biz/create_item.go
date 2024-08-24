package biz

import (
	"context"
	"main.go/module/item/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoCreateItem) error
}
type CreateItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *CreateItemBiz {
	return &CreateItemBiz{store: store}
}
func (biz *CreateItemBiz) CreateNewItem(ctx context.Context, data *model.TodoCreateItem) error {
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
