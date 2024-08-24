package storage

import (
	"context"
)

func (s *SqlModel) UpdateReadNotify(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("notify").Where(cond).Update("is_read", true).Error; err != nil {
		return err
	}
	return nil
}
