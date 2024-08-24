package storage

import (
	"context"
	"main.go/module/follow/model"
)

func (s *sqlModel) Unfollow(ctx context.Context, user *model.CreateFollower) error {
	if err := s.db.Table("follow").Where("user_id=? and by_user_id=?", user.UserId, user.ByUserId).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
