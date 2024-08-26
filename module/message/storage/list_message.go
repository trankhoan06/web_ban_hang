package storage

import (
	"context"
	"main.go/module/message/model"
)

func (s *SqlModel) ListMessage(ctx context.Context, sender, receiver int) (*[]model.Message, error) {
	var result []model.Message
	if err := s.db.Table("message").Where("(sender_id=? and receiver_id=? and is_status_sender<>?) or (sender_id=? and receiver_id=?and is_status_receiver<>?)", sender, receiver, model.StatusDeleted, receiver, sender, model.StatusDeleted).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
