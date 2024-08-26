package biz

import (
	"context"
	"errors"
)

func (biz *MessageBiz) NewDeletedUserMessage(ctx context.Context, sender, receive int) error {
	_, err := biz.store.FindUserMessage(ctx, sender, receive, "OwnerReceiverId")
	if err == nil {
		return errors.New("message deleted or no exist")
	}
	if err := biz.store.DeletedUserMessage(ctx, sender, receive); err != nil {
		return err
	}
	return nil
}
