package biz

import (
	"context"
	"errors"
)

func (biz *NotifyBiz) NewReadStatus(ctx context.Context, id, userId int) error {
	_, err := biz.store.FindNotify(ctx, id)
	if err != nil {
		return errors.New("notify has been deleted or no exist")
	}
	if err := biz.store.UpdateReadNotify(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
