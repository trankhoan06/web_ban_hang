package gincart

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/cart/biz"
	"main.go/module/cart/model"
	"main.go/module/cart/storage"
	"net/http"
	"strconv"
)

func UpdateItemCart(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		itemId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var data model.CartUpdateUser
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requester := c.MustGet(common.Current_user).(common.Requester)
		data.ItemId = itemId
		data.UserId = requester.GetUserId()
		store := storage.NewSqlModel(db)
		busines := biz.NewCartUpdateAmountItemBiz(store)
		if err := busines.NewCartUpdateAmountItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"update": "sucess"})

	}
}
