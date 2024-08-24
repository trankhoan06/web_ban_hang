package ginOrder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	storageItem "main.go/module/item/storage"
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
		 
		c.JSON(http.StatusOK, gin.H{"appreciate": appreciate})
	}
}
