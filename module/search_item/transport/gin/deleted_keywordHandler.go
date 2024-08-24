package ginSearch

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/search_item/biz"
	"main.go/module/search_item/model"
	"main.go/module/search_item/storage"
	"net/http"
	"strings"
)

func DeletedKeyword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var filter model.Filter
		if err := c.ShouldBindJSON(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userId := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		filter.UserId = &userId
		filter.Keyword = strings.TrimSpace(filter.Keyword)
		store := storage.NewSqlModel(db)
		business := biz.NewSearchBiz(store)
		err := business.DeletedKeyword(c.Request.Context(), &filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
