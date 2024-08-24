package ginOrder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	bizItem "main.go/module/item/biz"
	storageItem "main.go/module/item/storage"
	bizNotify "main.go/module/notify/biz"
	modelNotify "main.go/module/notify/model"
	storageNotify "main.go/module/notify/storage"
	"main.go/module/order_by_user/biz"
	modelOrder "main.go/module/order_by_user/model"
	storageOrder "main.go/module/order_by_user/storage"
	"net/http"
	"strconv"
)

func CreateOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data modelOrder.CreateOrder
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if _, err := strconv.Atoi(data.Telephone); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storageOrder.NewSqlModel(db)
		store1 := storageItem.NewSqlModel(db)
		requester := c.MustGet(common.Current_user).(common.Requester)
		data.UserId = requester.GetUserId()
		business := biz.NewOrderBiz(store, store1)
		if err := business.CreateOrder(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		businessItem := bizItem.NewUpdateAmountItemBiz(store1)
		if err := businessItem.NewUpdateAmountItem(c.Request.Context(), data.ItemId, data.UserId, data.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		businessItem1 := bizItem.NewGetItemBiz(store1, requester)
		item, errItem := businessItem1.GetNewItem(c.Request.Context(), data.ItemId)
		if errItem != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errItem.Error()})
			return
		}
		storeNotify := storageNotify.NewSQLModel(db)
		businessNotify := bizNotify.NewNotifyBiz(storeNotify)
		notify := &modelNotify.CreateNotify{
			UserId:      item.UserId,
			ItemId:      &data.ItemId,
			CreatorId:   data.UserId,
			Message:     fmt.Sprintf("%s has been like order this item ", item.Owner.LastName),
			TypeMessage: modelNotify.TypeOrder,
		}
		if err := businessNotify.CreateNotify(c.Request.Context(), notify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
		return
	}
}
