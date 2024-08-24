package biz

import (
	"context"
	"main.go/module/notify/model"
	modelUser "main.go/module/user/model"
	"time"
)

type NotifyStorage interface {
	DeletedNotify(ctx context.Context, id, userID int) error
	UpdateReadNotify(ctx context.Context, cond map[string]interface{}) error
	FindNotify(ctx context.Context, id int) (*model.Notify, error)
	ListNotify(ctx context.Context, userId int) (*[]model.Notify, error)
	SendNotify(ctx context.Context, data *model.CreateNotify) error
	DeletedNotifyOfCreator(ctx context.Context, userId int, message string, CreateAt time.Time) error
	DeletedNotifyComment(ctx context.Context, itemId, id int) error
	DeletedNotifyLikeItem(ctx context.Context, creatorId int, itemId int) error
}
type GetUserBiz interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
}
type NotifyBiz struct {
	store NotifyStorage
}
type NotifyUserBiz struct {
	store  NotifyStorage
	store1 GetUserBiz
}

func NewNotifyBiz(store NotifyStorage) *NotifyBiz {
	return &NotifyBiz{store}
}
func NewNotifyUserBiz(store NotifyStorage, store1 GetUserBiz) *NotifyUserBiz {
	return &NotifyUserBiz{store: store, store1: store1}
}
