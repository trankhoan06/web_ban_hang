package ginOrder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	bizItem "main.go/module/item/biz"
	storageitem "main.go/module/item/storage"
	bizNotify "main.go/module/notify/biz"
	modelNotify "main.go/module/notify/model"
	storageNotify "main.go/module/notify/storage"
	"main.go/module/order_by_user/biz"
	"main.go/module/order_by_user/storage"
	"net/http"
	"strconv"
)

func CancelOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		business := biz.NewOrderUserBiz(store)
		order, err := business.NewGetOrder(c.Request.Context(), id, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := business.NewCancelOrder(c.Request.Context(), id, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storeItem := storageitem.NewSqlModel(db)
		businessItem := bizItem.NewGetItemBiz(storeItem, c.MustGet(common.Current_user).(common.Requester))
		item, errItem := businessItem.GetNewItem(c.Request.Context(), order.ItemId)
		if errItem != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errItem.Error()})
			return
		}
		storeNotify := storageNotify.NewSQLModel(db)
		businessNotify := bizNotify.NewNotifyBiz(storeNotify)
		notify := &modelNotify.CreateNotify{
			UserId:    item.UserId,
			ItemId:    &item.Id,
			CreatorId: userId,
			Message:   fmt.Sprintf("%s has been cancel this item ", item.Owner.LastName),
		}
		if err := businessNotify.CreateNotify(c.Request.Context(), notify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})

	}
}
