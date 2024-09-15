package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/item/biz"
	"main.go/module/item/model"
	"main.go/module/item/storage"
	"net/http"
	"strconv"
)

func UpdateSortItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		itemId, err := strconv.Atoi(c.Query("item_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id_item don't know"})
			return
		}
		var data model.UpdateSortItem
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.ItemId = itemId
		store := storage.NewSqlModel(db)
		bussines := biz.NewSortItemBiz(store)
		if err := bussines.NewUpdateSortItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
