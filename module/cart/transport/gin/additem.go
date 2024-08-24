package gincart

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/cart/biz"
	"main.go/module/cart/model"
	storagecart "main.go/module/cart/storage"
	storageitem "main.go/module/item/storage"
	"net/http"
)

func AddItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CartCreateUser
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storagecart.NewSqlModel(db)
		store1 := storageitem.NewSqlModel(db)
		busines := biz.NewAddItemBiz(store, store1)
		requester := c.MustGet(common.Current_user).(common.Requester)
		data.UserId = requester.GetUserId()
		err := busines.NewAddItem(c.Request.Context(), &data, requester)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
