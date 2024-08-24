package biz

import (
	"context"
	"main.go/common"
	"main.go/module/item/model"
)

type DeletedItemStorage interface {
	DeletedItem(ctx context.Context, cond map[string]interface{}) error
	GetOwnItem(ctx context.Context, cond map[string]interface{}, user map[string]interface{}) (*model.TodoList, error)
}
type DeletedItemBiz struct {
	store   DeletedItemStorage
	request common.Requester
}

func NewDeletedItemBiz(store DeletedItemStorage, requester common.Requester) *DeletedItemBiz {
	return &DeletedItemBiz{store: store, request: requester}
}
func (biz *DeletedItemBiz) DeleteItem(ctx context.Context, id int) error {
	item, err := biz.store.GetOwnItem(ctx, map[string]interface{}{"id": id}, map[string]interface{}{"user_id": biz.request.GetUserId()})
	if err != nil {
		return err
	}
	if biz.request.GetRole() != "admin" && biz.request.GetUserId() != item.UserId {
		return common.ErrNoPermiss
	}
	if err1 := biz.store.DeletedItem(ctx, map[string]interface{}{"id": id}); err1 != nil {
		return err1
	}
	return nil
}
