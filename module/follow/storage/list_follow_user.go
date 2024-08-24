package storage

import (
	"context"
	"main.go/module/follow/model"
)

func (s *sqlModel) ListFollowUser(ctx context.Context, column string, userID int, moreKey ...string) (*[]model.Follower, error) {
	var followUser []model.Follower
	db := s.db.Where(column, userID)
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	if err := db.Find(&followUser).Error; err != nil {
		return nil, err
	}
	return &followUser, nil
}
