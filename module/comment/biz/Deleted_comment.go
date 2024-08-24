package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/module/comment/model"
)

type DeletedCommentStorage interface {
	DeletedComment(ctx context.Context, id int) error
}
type DeletedCommentBiz struct {
	store  DeletedCommentStorage
	store1 GetItemStorage
}

func NewDeletedCommentBiz(store DeletedCommentStorage, store1 GetItemStorage) *DeletedCommentBiz {
	return &DeletedCommentBiz{
		store:  store,
		store1: store1,
	}
}
func (biz *DeletedCommentBiz) DeleteChildComment(ctx context.Context, root *common.TreeComment) {
	for _, v := range root.Child {
		if v.Child != nil {
			biz.DeleteChildComment(ctx, v)
		}
		comment := v.Val.(model.CommentUser)
		_ = biz.store.DeletedComment(ctx, comment.Id)
	}
}

func (biz *DeletedCommentBiz) NewDeletedComment(ctx context.Context, userId int, root *common.TreeComment, ownerId int) error {
	comment := root.Val.(model.CommentUser)
	if comment.UserId != userId && comment.OwnerItem != ownerId {
		return common.ErrUserUpdate(errors.New("No Permission"))
	}
	if err := biz.store.DeletedComment(ctx, comment.Id); err != nil {
		return err
	}
	biz.DeleteChildComment(ctx, root)
	return nil

}
