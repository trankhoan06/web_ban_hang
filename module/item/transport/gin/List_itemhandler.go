package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/item/biz"
	"main.go/module/item/model"
	"main.go/module/item/storage"
	"net/http"
)

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var paging model.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlModel(db)
		business := biz.NewListItemBiz(store)
		result, err := business.ListNewItem(c.Request.Context(), &paging, &filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data":         result,
			"paging total": paging,
			"filter":       filter,
		})
	}
}
