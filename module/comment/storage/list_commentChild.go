package storage

import (
	"context"
	modelcart "main.go/module/cart/model"
	"main.go/module/comment/model"
)

func (s *SqlModel) ListCommentChild(ctx context.Context, itemId, parentId int, moreKeyItem []string, moreKey ...string) (*[]model.CommentUser, error) {
	var result []model.CommentUser
	db := s.db.Table("comment")
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	for i := range moreKeyItem {
		db = db.Preload(moreKeyItem[i])
	}
	if err := db.Where("item_id=? and parent_id=? and status<>?", itemId, parentId, modelcart.StatusRemove).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
