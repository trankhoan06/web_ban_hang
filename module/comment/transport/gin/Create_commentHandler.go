package ginComment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/comment/biz"
	"main.go/module/comment/model"
	storageComment "main.go/module/comment/storage"
	storageItem "main.go/module/item/storage"
	bizNotify "main.go/module/notify/biz"
	modelNotify "main.go/module/notify/model"
	storageNotify "main.go/module/notify/storage"
	"net/http"
	"strconv"
)

func CreateComment(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		idItem, err := strconv.Atoi(c.Param("item_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var data model.CreateComment
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requester := c.MustGet(common.Current_user).(common.Requester)
		store := storageComment.NewSqlModel(db)
		store1 := storageItem.NewSqlModel(db)
		data.ItemId = idItem
		data.UserId = requester.GetUserId()
		business := biz.NewCreateCommentBiz(store, store1)
		comment, err := business.NewCreateComment(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storeNotify := storageNotify.NewSQLModel(db)
		businessNotify := bizNotify.NewNotifyBiz(storeNotify)
		notify := &modelNotify.CreateNotify{
			UserId:      comment.OwnerItem,
			ItemId:      &idItem,
			CommentId:   &comment.Id,
			TypeMessage: modelNotify.TypeComment,
			CreatorId:   requester.GetUserId(),
			Message:     fmt.Sprintf("%s has been comment regarding your item ", comment.Owner.LastName),
		}
		if err := businessNotify.CreateNotify(c.Request.Context(), notify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": comment})
	}
}
