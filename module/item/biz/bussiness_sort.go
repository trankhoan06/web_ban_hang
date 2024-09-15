package biz

import (
	"context"
	"main.go/module/item/model"
)

type SortItemStorage interface {
	UpdateSortItem(ctx context.Context, data *model.UpdateSortItem) error
	DeletedSortItem(ctx context.Context, id int) error
	CreateSortItem(ctx context.Context, data *model.CreateSortItem) error
	ListSortItem(ctx context.Context, idItem int) (*[]model.SortItem, error)
}

type SortItemBiz struct {
	storeSort SortItemStorage
}

func NewSortItemBiz(storeSort SortItemStorage) *SortItemBiz {
	return &SortItemBiz{storeSort: storeSort}
}
