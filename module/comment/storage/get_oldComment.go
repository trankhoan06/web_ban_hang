package storage

import (
	"context"
	modelComment "main.go/module/comment/model"
)

func (s *SqlModel) GetOldComment(ctx context.Context, id int) (*[]modelComment.OldComment, error) {
	var result []modelComment.OldComment
	if err := s.db.Where("original_id=?", id).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
