package ginSearch

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/search_item/biz"
	"main.go/module/search_item/model"
	"main.go/module/search_item/storage"
	"net/http"
)

func SearchItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var filter model.Filter
		var category model.CategorySearch
		if err := c.ShouldBindJSON(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		category.Name = c.Query("name")
		category.Arrangement = c.Query("arrangement")
		userID := c.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSqlModel(db)
		business := biz.NewSearchBiz(store)
		filter.UserId = &userID
		items, err := business.SearchItem(c.Request.Context(), &filter, &category)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": items})
	}
}
