package biz

import (
	"context"
	"main.go/common"
	"main.go/module/item/model"
)

type GetOwnItemStorage interface {
	GetOwnItem(ctx context.Context, cond map[string]interface{}, userid map[string]interface{}) (*model.TodoList, error)
}
type GetOwnItemBiz struct {
	store   GetOwnItemStorage
	request common.Requester
}

func NewGetOwnItemBiz(store GetOwnItemStorage, requester common.Requester) *GetOwnItemBiz {
	return &GetOwnItemBiz{store: store, request: requester}
}
func (biz *GetOwnItemBiz) GetNewItem(ctx context.Context, id int) (*model.TodoList, error) {
	data, err := biz.store.GetOwnItem(ctx, map[string]interface{}{"id": id}, map[string]interface{}{"user_id": biz.request.GetUserId()})
	if err != nil {
		return nil, err
	}
	return data, nil
}
