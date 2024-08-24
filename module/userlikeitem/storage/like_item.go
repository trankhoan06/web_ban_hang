package storage

import (
	"context"
	"main.go/module/userlikeitem/model"
)

func (s *SqlModel) ListLike(ctx context.Context, cond map[string]interface{}) (*[]model.UserLikeItem, error) {
	var data []model.UserLikeItem
	db := s.db.Preload("Owner")
	if err := db.Where(cond).Find(&data).Error; err != nil {
		return nil, err

	}
	return &data, nil
}
