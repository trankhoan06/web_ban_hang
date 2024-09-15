package biz

import (
	"context"
	"main.go/module/item/model"
)

func (biz *SortItemBiz) NewCreateSortItem(ctx context.Context, data *model.CreateSortItem) error {
	if err := biz.storeSort.CreateSortItem(ctx, data); err != nil {
		return err
	}
	return nil
}
