package storage

import (
	"context"
	"main.go/module/cart/model"
)

func (s *SqlModel) DeletedItemCart(ctx context.Context, itemId, userId int) error {
	if err := s.db.Table("cart_user").Where("user_id=? and item_id=?", userId, itemId).Update("status", model.StatusRemove).Error; err != nil {
		return err
	}
	return nil
}
