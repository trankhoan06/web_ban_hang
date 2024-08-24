package biz

import (
	"context"
	"time"
)

func (biz *NotifyBiz) DeletedNotifyOfCreator(ctx context.Context, creatorId int, message string, CreateAt time.Time) error {
	if err := biz.store.DeletedNotifyOfCreator(ctx, creatorId, message, CreateAt); err != nil {
		return err
	}
	return nil
}
