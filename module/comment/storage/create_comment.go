package storage

import (
	"context"
	modelComment "main.go/module/comment/model"
)

func (s *SqlModel) CreateComment(ctx context.Context, comment *modelComment.CreateComment) error {
	db := s.db.Begin()
	if err := db.Create(comment).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
