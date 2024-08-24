package storage

import (
	"context"
	"main.go/module/follow/model"
)

func (s *sqlModel) GetUserFollow(ctx context.Context, user *model.CreateFollower) (*model.Follower, error) {
	var data model.Follower
	if err := s.db.Table("follow").Where("user_id=? and by_user_id=?", user.UserId, user.ByUserId).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
