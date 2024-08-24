package gincart

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/cart/biz"
	"main.go/module/cart/storage"
	"net/http"
)

func ListItemCart(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := storage.NewSqlModel(db)
		requester := c.MustGet(common.Current_user).(common.Requester)
		busines := biz.NewListItemCartBiz(store)
		result, err := busines.NewCartItemCart(c.Request.Context(), requester)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
