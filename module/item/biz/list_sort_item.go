package biz

import (
	"context"
	"main.go/module/item/model"
)

func (biz *SortItemBiz) NewListSortItem(ctx context.Context, itemId int) (*[]model.SortItem, error) {
	sort, err := biz.storeSort.ListSortItem(ctx, itemId)
	if err != nil {
		return nil, err
	}
	return sort, nil
}
