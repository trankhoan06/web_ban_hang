package storage

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/module/comment/model"
)

func (s *SqlModel) FindComment(ctx context.Context, id int) (*model.CommentUser, error) {
	var data model.CommentUser
	db := s.db.Preload("Owner")

	if err := db.Where("id=? ", id).First(&data).Error; err != nil {
		return nil, err
	}
	if *data.Status == model.CommentStatusRemove {
		return nil, common.ErrCommonDeleted(errors.New("Comment has been deleted or no exist"))
	}
	return &data, nil
}
