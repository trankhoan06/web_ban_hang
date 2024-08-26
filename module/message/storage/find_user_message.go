package storage

import (
	"context"
	"main.go/module/message/model"
)

func (s *SqlModel) FindUserMessage(ctx context.Context, sender int, receive int, moreKey ...string) (*model.UserMessage, error) {
	var result model.UserMessage
	db := s.db.Select("sender_id, receiver_id").Where("(sender_id=? and receiver_id=? and is_status_sender<>?) or (sender_id=? and receiver_id=?  and is_status_receive<>?)", sender, receive, model.StatusDeleted, sender, receive, model.StatusDeleted).Group("sender_id, receiver_id")
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	if err := db.First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
