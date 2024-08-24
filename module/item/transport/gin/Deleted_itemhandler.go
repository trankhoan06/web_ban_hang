package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	bizcomment "main.go/module/comment/biz"
	storageComment "main.go/module/comment/storage"
	bizItem "main.go/module/item/biz"
	storageItem "main.go/module/item/storage"
	"net/http"
	"strconv"
)

func DeletedItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		request := c.MustGet(common.Current_user).(common.Requester)

		store := storageItem.NewSqlModel(db)
		storeComment := storageComment.NewSqlModel(db)
		businessListComment := bizcomment.NewListCommentBiz(storeComment, store)
		root, err := businessListComment.NewListComment(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		businessDeletedComment := bizcomment.NewDeletedAllCommentBiz(storeComment)
		if err := businessDeletedComment.NewDeletedAllComment(c.Request.Context(), root); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		business := bizItem.NewDeletedItemBiz(store, request)
		if err := business.DeleteItem(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
