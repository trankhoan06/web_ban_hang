package biz

import (
	"context"
	"errors"
)

func (biz *SearchBiz) DeletedAllKeyword(ctx context.Context, userId int) error {
	searchs, err := biz.store.ListSearch(ctx, userId)
	if err != nil {
		return err
	}
	if len(*searchs) == 0 {
		return errors.New("All keyword has been deleted")
	}
	if err := biz.store.DeletedAllSearch(ctx, userId); err != nil {
		return err
	}
	return nil
}
