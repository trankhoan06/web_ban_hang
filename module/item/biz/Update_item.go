package biz

import (
	"context"
	"main.go/common"
	"main.go/module/item/model"
)

type UpdateItemStorage interface {
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoUpdateItem) error
	GetOwnItem(ctx context.Context, cond map[string]interface{}, userid map[string]interface{}) (*model.TodoList, error)
}
type UpdateItemBiz struct {
	store   UpdateItemStorage
	request common.Requester
}

func NewUpdateItemBiz(store UpdateItemStorage, requester common.Requester) *UpdateItemBiz {
	return &UpdateItemBiz{store: store, request: requester}
}
func (biz *UpdateItemBiz) UpdateNewItem(ctx context.Context, id int, data *model.TodoUpdateItem) error {
	item, err := biz.store.GetOwnItem(ctx, map[string]interface{}{"id": id}, map[string]interface{}{"user_id": biz.request.GetUserId()})
	if err != nil {
		return err
	}
	if item.UserId != data.UserId {
		return common.ErrNoPermiss
	}
	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}
	return nil
}
