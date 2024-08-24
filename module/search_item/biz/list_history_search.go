package biz

import (
	"context"
	"main.go/module/search_item/model"
)

func (biz *SearchBiz) ListHistorySearch(ctx context.Context, keyword *model.Filter) (*[]model.Search, error) {
	if keyword.Keyword == "" {
		search, errSearch := biz.store.ListSearch(ctx, *keyword.UserId)
		if errSearch != nil {
			return nil, errSearch
		}
		return search, nil
	} else {
		search, err := biz.store.ListFilterSearch(ctx, keyword.Keyword)
		if err != nil {
			return nil, err
		}
		return search, nil
	}
}
