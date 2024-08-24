package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/module/user/model"
	"strconv"
	"strings"
)

type UpdateUserStorage interface {
	UpdateUser(ctx context.Context, data *model.UpdateUser, cond map[string]interface{}) error
}
type UpdateUserbiz struct {
	store UpdateUserStorage
}

func NewUpdateUserBiz(store UpdateUserStorage) *UpdateUserbiz {
	return &UpdateUserbiz{store: store}
}
func (biz *UpdateUserbiz) NewUpdateUser(ctx context.Context, data *model.UpdateUser, request common.Requester) error {
	if data.Role == "admin" {
		return errors.New("No permiss role for admin")
	}
	_, err := strconv.Atoi(data.Phone)
	if err != nil {
		return err
	}
	data.FirtName = strings.Title(data.FirtName)
	data.LastName = strings.Title(data.LastName)
	if err := biz.store.UpdateUser(ctx, data, map[string]interface{}{"id": request.GetUserId()}); err != nil {
		return err
	}
	return nil
}
