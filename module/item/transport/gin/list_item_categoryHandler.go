package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/item/biz"
	"main.go/module/item/storage"
	"net/http"
)

func ListItemCategory(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		category := c.Query("category")
		if category == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category is required"})
			return
		}
		store := storage.NewSqlModel(db)
		business := biz.NewCategoryItemBiz(store)
		res, errRes := business.NewListItemCategory(c.Request.Context(), category)
		if errRes != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errRes})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
