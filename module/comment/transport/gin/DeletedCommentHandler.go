package ginComment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/comment/biz"
	storageComment "main.go/module/comment/storage"
	bizItem "main.go/module/item/biz"
	storageItem "main.go/module/item/storage"
	storageNotify "main.go/module/notify/storage"
	"net/http"
	"strconv"
)

func DeletedComment(db *gorm.DB) func(*gin.Context) {
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
		busines1 := biz.NewGetListParentAndChildBiz(store, store1)
		root, err1 := busines1.NewGetListParentAndChild(c.Request.Context(), itemId, id)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err1.Error(),
			})
			return
		}
		busines := biz.NewDeletedCommentBiz(store, store1)
		requester := c.MustGet(common.Current_user).(common.Requester)
		businessItem := bizItem.NewGetItemBiz(store1, requester)
		item, err := businessItem.GetNewItem(c.Request.Context(), itemId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := busines.NewDeletedComment(c.Request.Context(), requester.GetUserId(), root, item.UserId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storeNotify := storageNotify.NewSQLModel(db)
		if err := storeNotify.DeletedNotifyComment(c.Request.Context(), itemId, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
