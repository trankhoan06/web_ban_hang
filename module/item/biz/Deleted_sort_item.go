package biz

import (
	"context"
)

func (biz *SortItemBiz) NewDeletedSortItem(ctx context.Context, id int) error {
	if err := biz.storeSort.DeletedSortItem(ctx, id); err != nil {
		return err
	}
	return nil
}
