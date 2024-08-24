package biz

import (
	"context"
	"errors"
	"main.go/module/search_item/model"
)

func (biz *SearchBiz) DeletedKeyword(ctx context.Context, keyword *model.Filter) error {
	if keywordDel, err := biz.store.FindSearch(ctx, keyword); err == nil {
		if keywordDel.Status == model.StatusSearchInactive {
			return errors.New("keyword has been deleted")
		}
	}
	if err := biz.store.UpdateStatusSearch(ctx, keyword, model.StatusSearchInactive); err != nil {
		return err
	}
	return nil
}
