package storage

import (
	"context"
	"main.go/module/notify/model"
)

func (s *SqlModel) DeletedNotifyComment(ctx context.Context, itemId, id int) error {
	if err := s.db.Table("notify").Where("item_id=? and comment_id=? and type_message=?", itemId, id, model.TypeComment).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
