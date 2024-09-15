package storage

import (
	"context"
)

func (s *SqlModel) DeletedSortItem(ctx context.Context, id int) error {
	if err := s.db.Table("sort_item").Where("id=?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
