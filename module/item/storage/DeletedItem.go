package storage

import "context"

func (s *SqlModel) DeletedItem(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("todo_items").Where(cond).Updates(map[string]interface{}{
		"status": "Deleted",
	}).Error; err != nil {
		return err
	}
	return nil
}
