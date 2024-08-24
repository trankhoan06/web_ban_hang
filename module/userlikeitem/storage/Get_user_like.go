package storage

import (
	"context"
	"main.go/module/userlikeitem/model"
)

func (s *SqlModel) GetUserLike(ctx context.Context, itemId int, userId int) (*model.LikeItem, error) {
	var like model.LikeItem
	if err := s.db.Where("item_id=? and user_id=?", itemId, userId).First(&like).Error; err != nil {
		return nil, err
	}
	return &like, nil
}
