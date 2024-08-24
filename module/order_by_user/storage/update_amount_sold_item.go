package storage

import (
	"context"
	"gorm.io/gorm"
)

func (s *SqlModel) UpdateAmountSoldItem(ctx context.Context, itemId, amountSold int) error {
	if err := s.db.Table("todo_items").Where("id=?", itemId).Update("amount_sold", gorm.Expr("amount_sold+?", amountSold)).Error; err != nil {
		return err
	}
	return nil
}
