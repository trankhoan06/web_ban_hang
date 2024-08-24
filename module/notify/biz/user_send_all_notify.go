package biz

import (
	"context"
	"main.go/common"
	"main.go/module/notify/model"
	modelUser "main.go/module/user/model"
)

func (biz *NotifyUserBiz) NewUserSendAllNotify(ctx context.Context, result *[]modelUser.LIstUserId, notify *model.CreateNotify) error {
	user, err := biz.store1.FindUser(ctx, map[string]interface{}{"id": notify.CreatorId})
	if err != nil {
		return err
	}
	if user.Role != "admin" {
		return common.ErrNoPermiss
	}
	for _, val := range *result {
		NewNotify := notify
		NewNotify.UserId = val.UserId
		if err := biz.store.SendNotify(ctx, NewNotify); err != nil {
			return err
		}
	}
	return nil
}
