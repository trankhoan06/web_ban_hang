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

func ListComment(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		itemId, err := strconv.Atoi(c.Param("item_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storageComment.NewSqlModel(db)
		store1 := storageItem.NewSqlModel(db)
		busines := biz.NewListCommentBiz(store, store1)
		root, err := busines.NewListComment(c.Request.Context(), itemId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": root})

	}
}
