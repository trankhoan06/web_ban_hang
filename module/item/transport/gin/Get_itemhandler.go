package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/item/biz"
	"main.go/module/item/storage"
	"net/http"
	"strconv"
)

func Getitem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		request := c.MustGet(common.Current_user).(common.Requester)
		store := storage.NewSqlModel(db)
		business := biz.NewGetItemBiz(store, request)
		data, err := business.GetNewItem(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
