package biz

import (
	"context"
	"main.go/common"
	"main.go/module/comment/model"
)

type DeletedAllCommentBiz struct {
	store DeletedCommentStorage
}

func NewDeletedAllCommentBiz(store DeletedCommentStorage) *DeletedAllCommentBiz {
	return &DeletedAllCommentBiz{
		store: store,
	}
}
func (biz *DeletedAllCommentBiz) DeleteChildComment(ctx context.Context, root *common.TreeComment) {
	for _, v := range root.Child {
		if v.Child != nil {
			biz.DeleteChildComment(ctx, v)
		}
		comment := v.Val.(model.CommentUser)
		_ = biz.store.DeletedComment(ctx, comment.Id)
	}
}

func (biz *DeletedAllCommentBiz) NewDeletedAllComment(ctx context.Context, root *common.TreeComment) error {
	biz.DeleteChildComment(ctx, root)
	return nil

}
