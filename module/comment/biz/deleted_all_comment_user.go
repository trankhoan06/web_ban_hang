package biz

import "context"

type DeletedAllCommentUserStorage interface {
	DeletedAllCommentUser(ctx context.Context, userId int) error
}

type DeletedAllCommentUserBiz struct {
	store DeletedAllCommentUserStorage
}

func NewDeletedAllCommentUserBiz(store DeletedAllCommentUserStorage) *DeletedAllCommentUserBiz {
	return &DeletedAllCommentUserBiz{store: store}
}
func (biz *DeletedAllCommentUserBiz) NewDeletedAllCommentUser(ctx context.Context, userId int) error {
	if err := biz.store.DeletedAllCommentUser(ctx, userId); err != nil {
		return err
	}
	return nil
}
