package biz

import (
	"golang.org/x/net/context"
	"main.go/common"
)

type RegisterRoleStorage interface {
	UpdateRole(ctx context.Context, userID int, role string) error
}
type RegisterRoleBiz struct {
	store RegisterRoleStorage
}

func NewRegisterRoleBiz(store RegisterRoleStorage) *RegisterRoleBiz {
	return &RegisterRoleBiz{store: store}
}
func (biz *RegisterRoleBiz) NewRegisterRole(ctx context.Context, userId int, roleName string) error {
	if roleName == "admin" {
		return common.ErrNoPermiss
	}
	if err := biz.store.UpdateRole(ctx, userId, roleName); err != nil {
		return err
	}
	return nil
}
