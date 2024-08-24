package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/item/biz"
	"main.go/module/item/model"
	"main.go/module/item/storage"
	"net/http"
	"strconv"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data model.TodoUpdateItem
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		request := c.MustGet(common.Current_user).(common.Requester)
		data.UserId = request.GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewUpdateItemBiz(store, request)
		if err := business.UpdateNewItem(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
