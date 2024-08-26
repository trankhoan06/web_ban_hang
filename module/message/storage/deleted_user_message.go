package storage

import (
	"context"
	"main.go/module/message/model"
)

func (s *SqlModel) DeletedUserMessage(ctx context.Context, sender, receive int) error {
	db := s.db.Table("message").Where("sender_id = ? and receiver_id =? and is_status_sender<>?", sender, receive, model.StatusDeleted)
	if err := db.Update("is_status_sender", model.StatusDeleted).Error; err != nil {
		return err
	}

	db1 := s.db.Table("message").Where("sender_id = ? and receiver_id =? and is_status_receive<>?", receive, sender, model.StatusDeleted)
	if err := db1.Update("is_status_receive", model.StatusDeleted).Error; err != nil {
		return err
	}

	return nil
}
