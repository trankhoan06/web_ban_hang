package storage

import (
	"context"
	"gorm.io/gorm"
)

func (s *SqlModel) CoutItem(ctx context.Context, itemId, userId, amount int) error {
	if err := s.db.Table("cart_user").Where("user_id=? and item_id=?", userId, itemId).Update("amount", gorm.Expr("amount + ?", amount)).Error; err != nil {
		return err
	}
	return nil
}
