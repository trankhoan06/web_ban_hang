package storage

import (
	"context"
	"main.go/module/user/model"
)

func (s *SqlModel) FindSendCode(ctx context.Context, cond map[string]interface{}) (*model.SendCode, error) {
	var send model.SendCode
	if err := s.db.Table("send_code_email").Where(cond).Last(&send).Error; err != nil {
		return nil, err
	}
	return &send, nil
}
