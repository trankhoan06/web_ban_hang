package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/item/biz"
	"main.go/module/item/model"
	"main.go/module/item/storage"
	"net/http"
)

func CreateSortItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.CreateSortItem
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		bussines := biz.NewSortItemBiz(store)
		if err := bussines.NewCreateSortItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
