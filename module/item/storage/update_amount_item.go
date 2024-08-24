package storage

import (
	"context"
	"gorm.io/gorm"
)

func (s *SqlModel) UpdateAmountItem(ctx context.Context, itemId, userId, amount int) error {
	if err := s.db.Table("todo_items").Where("id=? and user_id=?", itemId, userId).Update("amount_item", gorm.Expr("amount_item-?", amount)).Error; err != nil {
		return err
	}
	return nil
}
