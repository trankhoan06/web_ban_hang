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

func GetListParentAndChild(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		itemId, err := strconv.Atoi(c.Param("item_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		id, err1 := strconv.Atoi(c.Param("id"))
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": err1.Error(),
			})
			return
		}
		store := storageComment.NewSqlModel(db)
		store1 := storageItem.NewSqlModel(db)
		busines := biz.NewGetListParentAndChildBiz(store, store1)
		root, err1 := busines.NewGetListParentAndChild(c.Request.Context(), itemId, id)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err1.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": root,
		})
	}
}
