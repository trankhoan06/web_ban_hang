package storage

import (
	"context"
	"main.go/module/comment/model"
)

func (s *SqlModel) ListParentComment(ctx context.Context, itemId int, moreKeyItem []string, moreKey ...string) (*[]model.CommentUser, error) {
	db := s.db.Table("comment").Where("item_id=? and parent_id IS NULL and status<>?", itemId, model.CommentStatusRemove)
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	for i := range moreKeyItem {
		db = db.Preload(moreKeyItem[i])
	}
	var result []model.CommentUser
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
