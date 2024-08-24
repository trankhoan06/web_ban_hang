package ginSearch

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/search_item/biz"
	"main.go/module/search_item/storage"
	"net/http"
)

func DeletedAllKeyword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewSearchBiz(store)
		err := business.DeletedAllKeyword(c.Request.Context(), userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
