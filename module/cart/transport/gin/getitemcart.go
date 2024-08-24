package gincart

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/cart/biz"
	storagecart "main.go/module/cart/storage"
	"net/http"
	"strconv"
)

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		itemId, err := strconv.Atoi(c.Param("itemId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storagecart.NewSqlModel(db)
		busines := biz.NewFindItemBiz(store)
		requester := c.MustGet(common.Current_user).(common.Requester)
		data, err1 := busines.NewFindItem(c.Request.Context(), itemId, requester)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
