package biz

import (
	"context"
	"main.go/module/notify/model"
)

func (biz *NotifyBiz) CreateNotify(ctx context.Context, notify *model.CreateNotify) error {
	if err := biz.store.SendNotify(ctx, notify); err != nil {
		return err
	}
	return nil
}
