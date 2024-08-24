package biz

import (
	"context"
	"main.go/common"
	"main.go/module/comment/model"
	modelitem "main.go/module/item/model"
)

type GetListParentAndChildStorage interface {
	FindComment(ctx context.Context, id int) (*model.CommentUser, error)
	ListCommentChild(ctx context.Context, itemId, parentId int, moreKeyItem []string, moreKey ...string) (*[]model.CommentUser, error)
}
type GetListParentAndChildBiz struct {
	store  GetListParentAndChildStorage
	store1 GetItemStorage
}

func NewGetListParentAndChildBiz(store GetListParentAndChildStorage, store1 GetItemStorage) *GetListParentAndChildBiz {
	return &GetListParentAndChildBiz{
		store:  store,
		store1: store1,
	}
}

func (biz *GetListParentAndChildBiz) ArrChild(root *common.TreeComment, childs *[]model.CommentUser, ctx context.Context) {
	for _, child := range *childs {
		rootChild := common.NewNode(child)
		root.Child = append(root.Child, rootChild)
		child1, err := biz.store.ListCommentChild(ctx, child.ItemId, child.Id, []string{"OwnerItemUser"}, "Owner")
		if err == nil {
			biz.ArrChild(rootChild, child1, ctx)
		}
	}
}
func (biz *GetListParentAndChildBiz) NewGetListParentAndChild(ctx context.Context, itemId int, id int) (*common.TreeComment, error) {
	_, err := biz.store1.GetItem(ctx, map[string]interface{}{modelitem.NameItem: itemId})
	if err != nil {
		return nil, err
	}
	data, err := biz.store.FindComment(ctx, id)
	if err != nil {
		return nil, err
	}
	var root common.TreeComment
	root.Val = *data
	child, err := biz.store.ListCommentChild(ctx, data.ItemId, data.Id, []string{"OwnerItemUser"}, "Owner")
	if err != nil {
		return nil, err
	}
	biz.ArrChild(&root, child, ctx)
	return &root, nil
}
