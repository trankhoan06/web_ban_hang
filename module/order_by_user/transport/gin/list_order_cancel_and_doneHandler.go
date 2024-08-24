package ginOrder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/order_by_user/biz"
	"main.go/module/order_by_user/storage"
	"net/http"
)

func ListOrderCancelAndDone(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		sellId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewOrderUserBiz(store)
		orders, err := business.NewListOrderCancelAndDone(c.Request.Context(), sellId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": orders})
	}
}
