package ginComment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/comment/biz"
	storageComment "main.go/module/comment/storage"
	storageItem "main.go/module/item/storage"
	"net/http"
	"strconv"
)

func GetOldComment(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		itemId, err := strconv.Atoi(c.Param("item_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storageComment.NewSqlModel(db)
		store1 := storageItem.NewSqlModel(db)
		busines := biz.NewGetOldCommentBiz(store, store1)
		result, err := busines.NewGetOldComment(c.Request.Context(), id, itemId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
