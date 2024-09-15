package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/item/biz"
	"main.go/module/item/storage"
	"net/http"
	"strconv"
)

func DeletedSortItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		business := biz.NewSortItemBiz(store)
		if err := business.NewDeletedSortItem(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
