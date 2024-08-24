package biz

import (
	"context"
	"errors"
	"main.go/common"
	modelComment "main.go/module/comment/model"
	modelitem "main.go/module/item/model"
)

type GetOldCommentStorage interface {
	GetOldComment(ctx context.Context, id int) (*[]modelComment.OldComment, error)
	FindComment(ctx context.Context, id int) (*modelComment.CommentUser, error)
}
type GetOldCommentBiz struct {
	store  GetOldCommentStorage
	store1 GetItemStorage
}

func NewGetOldCommentBiz(store GetOldCommentStorage, store1 GetItemStorage) *GetOldCommentBiz {
	return &GetOldCommentBiz{store: store, store1: store1}
}
func (biz *GetOldCommentBiz) NewGetOldComment(ctx context.Context, id int, itemID int) (*[]modelComment.OldComment, error) {
	if _, err := biz.store1.GetItem(ctx, map[string]interface{}{modelitem.NameItem: itemID}); err != nil {
		return nil, err
	}
	comment, err := biz.store.FindComment(ctx, id)
	if err != nil {
		return nil, err
	}
	if *comment.IsUpdate == false {
		return nil, common.ErrUneditedUpdate(errors.New("Unedited comments"))
	}
	result, err := biz.store.GetOldComment(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
