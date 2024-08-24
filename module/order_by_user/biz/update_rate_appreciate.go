package biz

import "context"

func (biz *OrderBiz) NewUpdateAppreciateItemBiz(ctx context.Context, itemId int) error {
	result, err := biz.store.GetListItemOrder(ctx, itemId)
	if err != nil {
		return err
	}
	var count float64
	var total float64
	for _, val := range *result {
		if val.Appreciate > 0 {
			count++
			total += float64(val.Appreciate)
		}
	}
	if err := biz.store1.UpdateAppreciateItem(ctx, itemId, total/count); err != nil {
		return err
	}
	return nil
}
