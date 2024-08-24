package storage

import "context"

func (s *SqlModel) DeletedNotify(ctx context.Context, id, userID int) error {
	if err := s.db.Table("notify").Where("id=?and user_id=?", id, userID).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
