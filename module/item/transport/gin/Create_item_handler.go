package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/item/biz"
	"main.go/module/item/model"
	"main.go/module/item/storage"
	"net/http"
)

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoCreateItem
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		request := c.MustGet(common.Current_user).(common.Requester)
		data.UserId = request.GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewCreateItemBiz(store)
		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
			"user": request.GetUserId(),
		})
	}
}
