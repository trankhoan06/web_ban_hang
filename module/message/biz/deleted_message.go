package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/module/message/model"
)

func (biz *MessageBiz) NewDeletedMessage(ctx context.Context, id, sender int) error {
	message, err := biz.store.GetMessage(ctx, id)
	if err != nil {
		return err
	}
	if message.SenderId != sender {
		return common.ErrNoPermiss
	}
	if message.IsStatusSender == model.StatusDeleted {
		return errors.New("message has been deleted")
	}
	if message.IsStatusReceive == model.StatusRemove && message.IsStatusSender == model.StatusRemove {
		return errors.New("message has been removed")
	}
	if message.IsStatusReceive == model.StatusDoing {
		if err := biz.store.DeletedMessage(ctx, id, "is_status_receive"); err != nil {
			return err
		}
	}
	if message.IsStatusSender == model.StatusDoing {
		if err := biz.store.DeletedMessage(ctx, id, "is_status_sender"); err != nil {
			return err
		}
	}
	return nil
}
