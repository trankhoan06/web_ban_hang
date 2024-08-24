package storage

import (
	"context"
	"main.go/module/notify/model"
)

func (s *SqlModel) ListNotifyOfCreator(ctx context.Context, createId int) (*[]model.CreatorNotify, error) {
	var result []model.CreatorNotify
	if err := s.db.Table("notify").Select("message, create_at").Group("message, create_at").Where("creator_id=? and type_message=?", createId, model.TypeEvent).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
