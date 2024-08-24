package storage

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/module/cart/model"
)

func (s *SqlModel) FindItem(ctx context.Context, itemId, userId int) (*model.CartUser, error) {
	var order model.CartUser
	db := s.db.Preload("Owner")
	if err := db.Table("cart_user").Where("user_id=? and item_id=?", userId, itemId).First(&order).Error; err != nil {
		return nil, common.ErrCart(err)
	}
	if *order.Status == model.StatusRemove {
		return &order, common.ErrCart(errors.New("no item of item have been deleted"))
	}
	return &order, nil
}
