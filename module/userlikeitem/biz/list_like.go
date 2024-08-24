package biz

import (
	"context"
	"main.go/module/userlikeitem/model"
)

type ListLikeStorage interface {
	ListLike(ctx context.Context, cond map[string]interface{}) (*[]model.UserLikeItem, error)
}
type ListLikeBiz struct {
	store ListLikeStorage
}

func NewListLikeBiz(store ListLikeStorage) *ListLikeBiz {
	return &ListLikeBiz{store: store}
}
func (biz *ListLikeBiz) NewListLike(ctx context.Context, id int) (*[]model.UserLikeItem, error) {
	data, err := biz.store.ListLike(ctx, map[string]interface{}{"item_id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
