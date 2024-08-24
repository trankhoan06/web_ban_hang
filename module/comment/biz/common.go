package biz

import (
	"context"
	modelitem "main.go/module/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*modelitem.TodoList, error)
}
