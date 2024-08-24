package biz

import (
	"context"
	"errors"
	"main.go/module/user/model"
)

type DeletedUserStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	UpdateStatusUser(ctx context.Context, userId int, status model.StatusUser) error
}
type DeletedFollowStorage interface {
	UnfollowAll(ctx context.Context, userId int) error
}
type DeletedUserBiz struct {
	store  DeletedUserStorage
	store2 DeletedFollowStorage
}

func NewDeletedUserBiz(store DeletedUserStorage, store2 DeletedFollowStorage) *DeletedUserBiz {
	return &DeletedUserBiz{store: store, store2: store2}
}
func (biz *DeletedUserBiz) NewDeletedUser(ctx context.Context, userId int) error {
	if data, err := biz.store.FindUser(ctx, map[string]interface{}{"id": userId}); err != nil {
		if data.Status == model.StatusUserInactive {
			return errors.New("account has been deleted")
		}
	}
	if err := biz.store.UpdateStatusUser(ctx, userId, model.StatusUserInactive); err != nil {
		return err
	}

	if err := biz.store2.UnfollowAll(ctx, userId); err != nil {
		return err
	}
	return nil
}
