package biz

import (
	"context"
	modelitem "main.go/module/item/model"
	"main.go/module/order_by_user/model"
)

type Business interface {
	CancelOrder(ctx context.Context, id int) error
	CreateOrder(ctx context.Context, order *model.CreateOrder) error
	GetOrder(ctx context.Context, addressColumn string, id int, userId int, keyUser, keyItem string) (*model.Order, error)
	ListOrder(ctx context.Context, userId int, column string, moreKeyUser []string, moreKeyItem ...string) (*[]model.Order, error)
	ListOrderCancelAndDone(ctx context.Context, userId int, column string, moreKeyUser []string, moreKeyItem ...string) (*[]model.Order, error)
	UpdateOrder(ctx context.Context, data *model.UpdateOrder) error
	UpdateStatusOrder(ctx context.Context, id int, status model.StatusOrder) error
	UpdateAmountSoldItem(ctx context.Context, itemId, amountSold int) error
	GetListItemOrder(ctx context.Context, itemId int) (*[]model.AppreciateOrder, error)
}
type FindItemOrder interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*modelitem.TodoList, error)
	UpdateAppreciateItem(ctx context.Context, itemId int, point float64) error
}
type OrderBiz struct {
	store  Business
	store1 FindItemOrder
}

func NewOrderBiz(store Business, store1 FindItemOrder) *OrderBiz {
	return &OrderBiz{store: store, store1: store1}
}

type OrderUserBiz struct {
	store Business
}

func NewOrderUserBiz(store Business) *OrderUserBiz {
	return &OrderUserBiz{store: store}
}
