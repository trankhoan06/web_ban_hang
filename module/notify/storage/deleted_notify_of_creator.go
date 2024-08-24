package storage

import (
	"context"
	"time"
)

func (s *SqlModel) DeletedNotifyOfCreator(ctx context.Context, userId int, message string, CreateAt time.Time) error {
	if err := s.db.Table("notify").Where("creator_id=? and message=?  and create_at=?", userId, message, CreateAt).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
