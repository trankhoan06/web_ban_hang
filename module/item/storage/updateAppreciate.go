package storage

import "context"

func (s *SqlModel) UpdateAppreciateItem(ctx context.Context, itemId int, point float64) error {
	if err := s.db.Table("todo_items").Where("id=?", itemId).Update("appreciate", point).Error; err != nil {
		return err
	}
	return nil
}
