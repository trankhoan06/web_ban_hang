package biz

import (
	"context"
	modelitem "main.go/module/item/model"
	"main.go/module/search_item/model"
)

type SearchStorage interface {
	SearchItem(ctx context.Context, keyword string, category *model.CategorySearch, moreKey ...string) (*[]modelitem.TodoList, error)
	FindSearch(ctx context.Context, keyword *model.Filter) (*model.Search, error)
	UpdateStatusSearch(ctx context.Context, keyword *model.Filter, status model.StatusSearch) error
	DeletedAllSearch(ctx context.Context, userId int) error
	CreateSearch(ctx context.Context, keyword *model.Filter) error
	ListSearch(ctx context.Context, user int) (*[]model.Search, error)
	UpdateTimeSearch(ctx context.Context, keyword *model.Filter) error
	ListFilterSearch(ctx context.Context, keyword string) (*[]model.Search, error)
}
type SearchBiz struct {
	store SearchStorage
}

func NewSearchBiz(store SearchStorage) *SearchBiz {
	return &SearchBiz{store}
}
