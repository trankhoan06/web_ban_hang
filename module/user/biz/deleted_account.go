package biz

import "context"

func (biz *UserBiz) NewDeletedAccount(ctx context.Context, userId int) error {
	if err := biz.store.DeletedAccount(ctx, userId); err != nil {
		return err
	}
	return nil
}
