package biz

import (
	"context"
	"main.go/module/user/model"
)

type ListAllUserIDStorage interface {
	ListUserId(ctx context.Context) (*[]model.LIstUserId, error)
}
type ListAllUserIDBiz struct {
	store ListAllUserIDStorage
}

func NewListAllUserIDBiz(s ListAllUserIDStorage) *ListAllUserIDBiz {
	return &ListAllUserIDBiz{s}
}
func (biz *ListAllUserIDBiz) NewListAllUserID(ctx context.Context) (*[]model.LIstUserId, error) {
	result, err := biz.store.ListUserId(ctx)
	if err != nil {
		return nil, err

	}
	return result, nil
}
