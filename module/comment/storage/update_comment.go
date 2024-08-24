package storage

import (
	"context"
	modelComment "main.go/module/comment/model"
)

func (s *SqlModel) UpdateComment(ctx context.Context, data *modelComment.UpdateComment) error {
	if err := s.db.Table("comment").Where("id=?", data.Id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
