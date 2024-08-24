package biz

import (
	"context"
	"main.go/common"
	"main.go/module/comment/model"
	modelitem "main.go/module/item/model"
)

type ListCommentStorage interface {
	ListParentComment(ctx context.Context, itemId int, moreKeyItem []string, moreKey ...string) (*[]model.CommentUser, error)
	ListCommentChild(ctx context.Context, itemId, parentId int, moreKeyItem []string, moreKey ...string) (*[]model.CommentUser, error)
}
type ListCommentBiz struct {
	store  ListCommentStorage
	store1 GetItemStorage
}

func NewListCommentBiz(store ListCommentStorage, store1 GetItemStorage) *ListCommentBiz {
	return &ListCommentBiz{store: store, store1: store1}
}

func (biz *ListCommentBiz) ArrChild(root *common.TreeComment, childs *[]model.CommentUser, ctx context.Context) {
	for _, child := range *childs {
		rootChild := common.NewNode(child)
		root.Child = append(root.Child, rootChild)
		child1, err := biz.store.ListCommentChild(ctx, child.ItemId, child.Id, []string{"OwnerItemUser"}, "Owner")
		if err == nil {
			biz.ArrChild(rootChild, child1, ctx)
		}
	}
}

func (biz *ListCommentBiz) NewListComment(ctx context.Context, itemId int) (*common.TreeComment, error) {
	item, err := biz.store1.GetItem(ctx, map[string]interface{}{modelitem.NameItem: itemId})
	if err != nil {
		return nil, err
	}
	var root common.TreeComment
	root.Val = item
	comment, err := biz.store.ListParentComment(ctx, itemId, []string{"OwnerItemUser"}, "Owner")
	if err != nil {
		return nil, err
	}
	biz.ArrChild(&root, comment, ctx)
	return &root, nil
}
