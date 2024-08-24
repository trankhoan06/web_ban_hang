package storage

import (
	"context"
)

func (s *sqlModel) UnfollowAll(ctx context.Context, userId int) error {
	if err := s.db.Table("follow").Where("user_id=? or by_user_id=? ", userId).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
