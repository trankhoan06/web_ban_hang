package storage

import (
	"context"
	"main.go/module/follow/model"
)

func (s *sqlModel) CreateFollowUser(ctx context.Context, user *model.CreateFollower) error {
	if err := s.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
