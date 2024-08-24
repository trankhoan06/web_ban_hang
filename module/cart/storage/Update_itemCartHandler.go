package storage

import (
	"context"
	"main.go/module/cart/model"
)

func (s *SqlModel) UpdateItemCart(ctx context.Context, update *model.CartUpdateUser) error {
	if err := s.db.Table("cart_user").Where("user_id=? and item_id=?", update.UserId, update.ItemId).Update("amount", update.Amount).Error; err != nil {
		return err
	}
	return nil
}
