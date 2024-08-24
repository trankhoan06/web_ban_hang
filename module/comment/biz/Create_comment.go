package biz

import (
	"context"
	"main.go/common"
	modelComment "main.go/module/comment/model"
	modelitem "main.go/module/item/model"
)

type CreateCommentStorage interface {
	CreateComment(ctx context.Context, comment *modelComment.CreateComment) error
	FindComment(ctx context.Context, id int) (*modelComment.CommentUser, error)
}

type CreateCommentBiz struct {
	store  CreateCommentStorage
	store1 GetItemStorage
}

func NewCreateCommentBiz(store CreateCommentStorage, store1 GetItemStorage) *CreateCommentBiz {
	return &CreateCommentBiz{
		store:  store,
		store1: store1,
	}
}

func (biz *CreateCommentBiz) NewCreateComment(ctx context.Context, comment *modelComment.CreateComment) (*modelComment.CommentUser, error) {
	item, err := biz.store1.GetItem(ctx, map[string]interface{}{modelitem.NameItem: comment.ItemId})
	comment.OwnerItem = item.UserId
	if err != nil {
		return nil, common.ErrFindItem
	}
	err = biz.store.CreateComment(ctx, comment)
	if err != nil {
		return nil, err
	}
	comment1, err := biz.store.FindComment(ctx, *comment.Id)
	if err != nil {
		return nil, err
	}
	return comment1, nil

}
