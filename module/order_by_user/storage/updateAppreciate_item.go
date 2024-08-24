package storage

import "context"

func (s *SqlModel) UpdateAppreciateOrder(ctx context.Context, id, userId, appreciatePoint int) error {
	if err := s.db.Table("order_by_user").Where("id=? and user_id=?", id, userId).Update("appreciate", appreciatePoint).Error; err != nil {
		return err
	}
	return nil
}
