package storage

import (
	"golang.org/x/net/context"
	"main.go/module/cart/model"
)

func (s *SqlModel) UpdateStatusItem(ctx context.Context, item *model.CartUser) error {
	if err := s.db.Where("item_id = ? and user_id=?", item.ItemId, item.UserId).Updates(item).Error; err != nil {
		return err
	}
	return nil
}
