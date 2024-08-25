package ginOrder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	storageItem "main.go/module/item/storage"
	bizNotify "main.go/module/notify/biz"
	modelNotify "main.go/module/notify/model"
	storageNotify "main.go/module/notify/storage"
	"main.go/module/order_by_user/biz"
	"main.go/module/order_by_user/model"
	"main.go/module/order_by_user/storage"
	"net/http"
	"strconv"
)

func AppreciateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var appreciate model.UpdateOrder
		if err := c.ShouldBind(&appreciate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewOrderUserBiz(store)
		item, err := business.NewGetOrder(c.Request.Context(), id, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		appreciate.Id = id
		appreciate.UserId = userId
		itemId, err := business.NewAppreciateOrder(c.Request.Context(), &appreciate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store1 := storageItem.NewSqlModel(db)
		business1 := biz.NewOrderBiz(store, store1)
		if err := business1.NewUpdateAppreciateItemBiz(c.Request.Context(), itemId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storeNotify := storageNotify.NewSQLModel(db)
		businessNotify := bizNotify.NewNotifyBiz(storeNotify)
		notify := &modelNotify.CreateNotify{
			UserId:      item.SellId,
			ItemId:      &item.ItemId,
			CreatorId:   appreciate.UserId,
			Message:     fmt.Sprintf("%s has been appreciate %d point ", item.OwnerUser.LastName),
			TypeMessage: modelNotify.TypeOrder,
		}
		if err := businessNotify.CreateNotify(c.Request.Context(), notify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"appreciate": appreciate})
	}
}
