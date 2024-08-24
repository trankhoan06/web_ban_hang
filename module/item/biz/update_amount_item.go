package biz

import (
	"context"
	"main.go/common"
	"main.go/module/item/model"
)

type UpdateAmountItemStorage interface {
	UpdateAmountItem(ctx context.Context, itemId, userId, amount int) error
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoList, error)
}
type UpdateAmountItemBiz struct {
	store UpdateAmountItemStorage
}

func NewUpdateAmountItemBiz(store UpdateAmountItemStorage) *UpdateAmountItemBiz {
	return &UpdateAmountItemBiz{store: store}
}

func (biz *UpdateAmountItemBiz) NewUpdateAmountItem(ctx context.Context, itemID, userId, amount int) error {
	item, err := biz.store.GetItem(ctx, map[string]interface{}{"id": itemID})
	if err != nil {
		return err
	}
	if item.UserId != userId {
		return common.ErrNoPermiss
	}
	if err := biz.store.UpdateAmountItem(ctx, itemID, userId, amount); err != nil {
		return err
	}
	return nil
}
