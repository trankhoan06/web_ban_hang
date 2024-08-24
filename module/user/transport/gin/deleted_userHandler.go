package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	bizcomment "main.go/module/comment/biz"
	storageComment "main.go/module/comment/storage"
	storageFollow "main.go/module/follow/storage"
	bizItem "main.go/module/item/biz"
	storageitem "main.go/module/item/storage"
	"main.go/module/user/biz"
	"main.go/module/user/storage"
	"net/http"
)

func DeletedUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		store2 := storageFollow.NewSQLModel(db)
		business := biz.NewDeletedUserBiz(store, store2)
		if err := business.NewDeletedUser(c.Request.Context(), userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store1 := storageitem.NewSqlModel(db)
		businessItem := bizItem.NewDeletedAllItemBiz(store1)
		items, err1 := businessItem.DeleteAllItem(c.Request.Context(), userId)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		}
		storeComment := storageComment.NewSqlModel(db)
		for _, item := range *items {
			businessListComment := bizcomment.NewListCommentBiz(storeComment, store1)
			root, err := businessListComment.NewListComment(c.Request.Context(), item.Id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			businessDeletedComment := bizcomment.NewDeletedAllCommentBiz(storeComment)
			if err := businessDeletedComment.NewDeletedAllComment(c.Request.Context(), root); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		businessComment := bizcomment.NewDeletedAllCommentUserBiz(storeComment)
		if err := businessComment.NewDeletedAllCommentUser(c.Request.Context(), userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"deleted": true})
	}
}
