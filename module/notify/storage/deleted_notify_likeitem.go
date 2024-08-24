package storage

import (
	"context"
	"main.go/module/notify/model"
)

func (s *SqlModel) DeletedNotifyLikeItem(ctx context.Context, creatorId int, itemId int) error {
	if err := s.db.Table("notify").Where("item_id=? and creator_id=? and type_message=?", itemId, creatorId, model.TypeLikeItem).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
