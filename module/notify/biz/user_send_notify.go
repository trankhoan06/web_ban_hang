package biz

import (
	"context"
	modelFollow "main.go/module/follow/model"
	"main.go/module/notify/model"
)

func (biz *NotifyBiz) SendNotify(ctx context.Context, data *model.CreateNotify, users *[]modelFollow.Follower) error {
	for _, follower := range *users {
		NewData := data
		NewData.UserId = follower.ByUserId
		if err := biz.store.SendNotify(ctx, data); err != nil {
			return err
		}
	}
	return nil
}
