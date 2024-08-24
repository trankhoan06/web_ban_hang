package biz

import (
	"context"
	"errors"
	"main.go/common"
	modelComment "main.go/module/comment/model"
	modelitem "main.go/module/item/model"
)

type UpdateCommentStorage interface {
	UpdateOldComment(ctx context.Context, data *modelComment.OldComment) error
	UpdateComment(ctx context.Context, data *modelComment.UpdateComment) error
	FindComment(ctx context.Context, id int) (*modelComment.CommentUser, error)
}
type UpdateCommentBiz struct {
	store  UpdateCommentStorage
	store1 GetItemStorage
}

func NewUpdateCommentBiz(store UpdateCommentStorage, store1 GetItemStorage) *UpdateCommentBiz {
	return &UpdateCommentBiz{
		store:  store,
		store1: store1,
	}
}
func (biz *UpdateCommentBiz) NewUpdateComment(ctx context.Context, data *modelComment.UpdateComment) error {
	_, err := biz.store1.GetItem(ctx, map[string]interface{}{modelitem.NameItem: data.ItemId})
	if err != nil {
		return err
	}
	comment, err := biz.store.FindComment(ctx, data.Id)
	if err != nil {
		return err
	}
	if comment.UserId != data.UserId {
		return common.ErrUserUpdate(errors.New("No Permission"))
	}
	oldComment := &modelComment.OldComment{
		UserId:     comment.UserId,
		ItemId:     comment.ItemId,
		Content:    comment.Content,
		Status:     modelComment.CommentStatusRemove,
		ParentId:   comment.ParentId,
		CreateAt:   comment.CreateAt,
		UpdateAt:   comment.UpdateAt,
		OriginalId: comment.Id,
		IsUpdate:   true,
	}
	if err := biz.store.UpdateOldComment(ctx, oldComment); err != nil {
		return err
	}
	if err := biz.store.UpdateComment(ctx, data); err != nil {
		return err
	}
	return nil

}
