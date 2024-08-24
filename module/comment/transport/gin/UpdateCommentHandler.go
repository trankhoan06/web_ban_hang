package ginComment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/comment/biz"
	"main.go/module/comment/model"
	storageComment "main.go/module/comment/storage"
	storageItem "main.go/module/item/storage"
	"net/http"
	"strconv"
)

func UpdateComment(db *gorm.DB) func(*gin.Context) {
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
		var data model.UpdateComment
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storageComment.NewSqlModel(db)
		store1 := storageItem.NewSqlModel(db)
		busines := biz.NewUpdateCommentBiz(store, store1)
		requester := c.MustGet(common.Current_user).(common.Requester)
		data.Id = id
		data.UserId = requester.GetUserId()
		data.ItemId = itemId
		data.IsUpdate = true
		if err := busines.NewUpdateComment(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
