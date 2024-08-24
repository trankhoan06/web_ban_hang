package biz

import (
	"context"
	"main.go/module/item/model"
)

type DeletedAllItemStorage interface {
	DeletedItem(ctx context.Context, cond map[string]interface{}) error
	ListOwnItem(ctx context.Context, userId int, moreKey ...string) (*[]model.TodoList, error)
}
type DeletedAllItemBiz struct {
	store DeletedAllItemStorage
}

func NewDeletedAllItemBiz(store DeletedAllItemStorage) *DeletedAllItemBiz {
	return &DeletedAllItemBiz{store: store}
}
func (biz *DeletedAllItemBiz) DeleteAllItem(ctx context.Context, UserId int) (*[]model.TodoList, error) {
	item, err := biz.store.ListOwnItem(ctx, UserId, "Owner")
	if err != nil {
		return nil, err
	}
	if err1 := biz.store.DeletedItem(ctx, map[string]interface{}{"user_id": UserId}); err1 != nil {
		return nil, err1
	}
	return item, nil
}
