package ginOrder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/order_by_user/biz"
	"main.go/module/order_by_user/model"
	"main.go/module/order_by_user/storage"
	"net/http"
	"strconv"
)

func UpdateStatusSell(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var status model.StatusOrderItem
		if err := c.ShouldBind(&status); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		business := biz.NewOrderUserBiz(store)
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		item, err := business.UpdateStatusSell(c.Request.Context(), id, userId, status.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if status.Status == model.StatusOrderDone {
			if err := business.NewUpdateAmountSoldItem(c.Request.Context(), item.ItemId, item.Amount); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": status})
	}
}
