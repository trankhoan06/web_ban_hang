package storage

import "context"

func (s *SqlModel) DeletedUserLike(ctx context.Context, itemId int, userId int) error {
	if err := s.db.Table("userlikeitem").Where("item_id=? and user_id=?", itemId, userId).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
