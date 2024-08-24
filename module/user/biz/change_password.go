package biz

import (
	"context"
	"main.go/common"
	"main.go/module/user/model"
)

type ChangePasswordbstorage interface {
	ChangePasswordsto(ctx context.Context, newPassword string, userID map[string]interface{}) error
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
}
type ChangeBiz struct {
	store ChangePasswordbstorage
	hash  Hasher
}

func NewChangeBiz(store ChangePasswordbstorage, hash Hasher) *ChangeBiz {
	return &ChangeBiz{
		store: store,
		hash:  hash,
	}
}
func (biz *ChangeBiz) NewChangePasswordBiz(ctx context.Context, requester common.Requester, changeUser *model.UpdatePass) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"id": requester.GetUserId()})
	if err != nil {
		return common.ErrEmailOfPass(err)
	}
	changeUser.PassWord = biz.hash.Hash(changeUser.PassWord + user.Salt)
	changeUser.NewPassWord = biz.hash.Hash(changeUser.NewPassWord + user.Salt)
	if changeUser.PassWord != user.PassWord {
		return common.ErrPass(err)
	}
	if err := biz.store.ChangePasswordsto(ctx, changeUser.NewPassWord, map[string]interface{}{"id": user.UserId}); err != nil {
		return err
	}
	return nil
}
