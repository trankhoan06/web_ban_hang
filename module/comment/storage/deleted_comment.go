package storage

import (
	"context"
	modelComment "main.go/module/comment/model"
)

func (s *SqlModel) DeletedComment(ctx context.Context, id int) error {
	if err := s.db.Table("comment").Where("id=?", id).Update("status", modelComment.CommentStatusRemove).Error; err != nil {
		return err
	}
	return nil
}
