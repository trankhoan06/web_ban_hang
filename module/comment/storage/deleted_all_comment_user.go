package storage

import (
	"context"
	modelComment "main.go/module/comment/model"
)

func (s *SqlModel) DeletedAllCommentUser(ctx context.Context, userId int) error {
	if err := s.db.Table("comment").Where("user_id=?", userId).Update("status", modelComment.CommentStatusRemove).Error; err != nil {
		return err
	}
	return nil
}
