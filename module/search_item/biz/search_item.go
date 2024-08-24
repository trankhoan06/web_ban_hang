package biz

import (
	"context"
	"errors"
	modelitem "main.go/module/item/model"
	"main.go/module/search_item/model"
	"strings"
)

func (biz *SearchBiz) SearchItem(ctx context.Context, keyword *model.Filter, categorySearch *model.CategorySearch) (*[]modelitem.TodoList, error) {
	keyword.Keyword = strings.TrimSpace(keyword.Keyword)
	if keyword.Keyword == "" {
		return nil, errors.New("keyword is require")
	}
	items, err := biz.store.SearchItem(ctx, keyword.Keyword, categorySearch, "Owner")
	if err != nil {
		return nil, err
	}
	if item, errFind := biz.store.FindSearch(ctx, keyword); errFind == nil {
		if item != nil && item.Status == model.StatusSearchInactive {
			if errStatus := biz.store.UpdateStatusSearch(ctx, keyword, model.StatusSearchActive); errStatus != nil {
				return nil, errStatus
			}
			if errUpdate := biz.store.UpdateTimeSearch(ctx, keyword); errUpdate != nil {
				return nil, errUpdate
			}
		} else {
			if errUpdate := biz.store.UpdateTimeSearch(ctx, keyword); errUpdate != nil {
				return nil, errUpdate
			}
		}
	} else {
		if errCreate := biz.store.CreateSearch(ctx, keyword); errCreate != nil {
			return nil, errCreate
		}
	}
	return items, nil
}
