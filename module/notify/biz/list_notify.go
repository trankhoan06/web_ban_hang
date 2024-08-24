package biz

import (
	"context"
	"main.go/module/notify/model"
)

func (biz *NotifyBiz) ListNotify(ctx context.Context, userId int) (*[]model.Notify, error) {
	notifys, err := biz.store.ListNotify(ctx, userId)
	if err != nil {
		return nil, err
	}
	return notifys, nil
}
