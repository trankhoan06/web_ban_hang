package biz

import (
	"context"
	"main.go/common"
	modelcart "main.go/module/cart/model"
	modelitem "main.go/module/item/model"
)

type AddItemStorage interface {
	FindItem(ctx context.Context, itemId, userId int) (*modelcart.CartUser, error)
	CoutItem(ctx context.Context, itemId, userId, amount int) error
	AddItem(ctx context.Context, item *modelcart.CartCreateUser) error
	UpdateStatusItem(ctx context.Context, item *modelcart.CartUser) error
}
type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*modelitem.TodoList, error)
}
type AddItemBiz struct {
	store  AddItemStorage
	store1 GetItemStorage
}

func NewAddItemBiz(store AddItemStorage, store1 GetItemStorage) *AddItemBiz {
	return &AddItemBiz{
		store:  store,
		store1: store1,
	}
}
func (biz *AddItemBiz) NewAddItem(ctx context.Context, item *modelcart.CartCreateUser, requester common.Requester) error {
	_, err := biz.store1.GetItem(ctx, map[string]interface{}{"id": item.ItemId})
	if err != nil {
		return common.ErrItem(err)
	}
	//nếu item đã có trong giỏ thì thêm số lượng
	item1, err := biz.store.FindItem(ctx, item.ItemId, item.UserId)
	//nếu sản phẩm đã bị xóa thì update lại status và amount
	if item1 != nil && *item1.Status == modelcart.StatusRemove {
		sta := modelcart.StatusActive
		item2 := &modelcart.CartUser{
			ItemId: item.ItemId,
			UserId: item.UserId,
			Amount: item.Amount,
			Status: &sta,
		}
		if err := biz.store.UpdateStatusItem(ctx, item2); err != nil {
			return err
		}
		return nil
		//	nếu sản phẩm có rồi thì thêm số lượng
	} else if err == nil {
		if err := biz.store.CoutItem(ctx, item.ItemId, item.UserId, item.Amount); err != nil {
			return err
		}
		return nil
	}
	//thêm item vào cart
	if err = biz.store.AddItem(ctx, item); err != nil {
		return err
	}
	return nil
}
