package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/item/biz"
	"main.go/module/item/storage"
	"net/http"
	"strconv"
)

func ListSortItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		itemId, err := strconv.Atoi(c.Query("item_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		store := storage.NewSqlModel(db)
		bussiness := biz.NewSortItemBiz(store)
		sort, err := bussiness.NewListSortItem(c.Request.Context(), itemId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": sort})
	}
}
