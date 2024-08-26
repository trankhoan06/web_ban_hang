package storage

import (
	"context"
	"main.go/module/message/model"
)

func (s *SqlModel) ListUserMessage(ctx context.Context, sender int, moreKey ...string) (*[]model.UserMessage, error) {
	var result []model.UserMessage
	db := s.db.Select("sender_id, receiver_id").Where("(sender_id=?  and is_status_sender<>?) or (receiver_id=?  and is_status_receive<>?)", sender, model.StatusDeleted, sender, model.StatusDeleted).Group("sender_id, receiver_id")
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
