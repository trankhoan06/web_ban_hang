package biz

import (
	"context"
	"main.go/module/item/model"
)

func (biz *SortItemBiz) NewUpdateSortItem(ctx context.Context, data *model.UpdateSortItem) error {
	if err := biz.storeSort.UpdateSortItem(ctx, data); err != nil {
		return err
	}
	return nil
}
