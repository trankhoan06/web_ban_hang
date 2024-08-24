package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/item/biz"
	"main.go/module/item/storage"
	"net/http"
)

func ListOwnItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		store := storage.NewSqlModel(db)
		business := biz.NewListOwnItemBiz(store)
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		result, err := business.ListNewItem(c.Request.Context(), userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
