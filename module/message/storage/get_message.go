package storage

import (
	"context"
	"main.go/module/message/model"
)

func (s *SqlModel) GetMessage(ctx context.Context, id int) (*model.Message, error) {
	var data model.Message
	if err := s.db.Table("message").Where("id=? ", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
