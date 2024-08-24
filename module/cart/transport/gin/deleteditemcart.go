package gincart

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/cart/biz"
	"main.go/module/cart/storage"
	"net/http"
	"strconv"
)

func DeletedItemCart(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		itemId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requester := c.MustGet(common.Current_user).(common.Requester)
		store := storage.NewSqlModel(db)
		busines := biz.NewCartDeletedAmountItemBiz(store)
		if err := busines.NewCartDeletedAmountItem(c.Request.Context(), itemId, requester.GetUserId()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"update": "sucess"})

	}
}
