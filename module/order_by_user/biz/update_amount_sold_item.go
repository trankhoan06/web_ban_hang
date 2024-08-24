package biz

import "context"

func (biz *OrderUserBiz) NewUpdateAmountSoldItem(ctx context.Context, id, amount int) error {
	if err := biz.store.UpdateAmountSoldItem(ctx, id, amount); err != nil {
		return err
	}
	return nil
}
