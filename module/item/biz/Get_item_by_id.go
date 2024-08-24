package biz

import (
	"context"
	"main.go/common"
	"main.go/module/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoList, error)
}
type GetItemBiz struct {
	store   GetItemStorage
	request common.Requester
}

func NewGetItemBiz(store GetItemStorage, requester common.Requester) *GetItemBiz {
	return &GetItemBiz{store: store, request: requester}
}
func (biz *GetItemBiz) GetNewItem(ctx context.Context, id int) (*model.TodoList, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
