package ginUserlikeItem

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	bizItem "main.go/module/item/biz"
	storageItem "main.go/module/item/storage"
	bizNotify "main.go/module/notify/biz"
	modelNotify "main.go/module/notify/model"
	storageNotify "main.go/module/notify/storage"
	"main.go/module/userlikeitem/biz"
	"main.go/module/userlikeitem/model"
	"main.go/module/userlikeitem/storage"
	"net/http"
	"strconv"
	"time"
)

func CreateLikeItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		now := time.Now().UTC()
		request := c.MustGet(common.Current_user).(common.Requester)
		data := model.LikeItem{
			UserId:   request.GetUserId(),
			ItemId:   id,
			CreateAt: &now,
		}
		store := storage.NewSqlModel(db)
		business := biz.NewUserLikeBiz(store)
		err = business.NewCreateUserLike(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storeItem := storageItem.NewSqlModel(db)
		businessItem := bizItem.NewGetItemBiz(storeItem, request)
		item, errItem := businessItem.GetNewItem(c.Request.Context(), id)
		if errItem != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errItem.Error()})
			return
		}
		storeNotify := storageNotify.NewSQLModel(db)
		businessNotify := bizNotify.NewNotifyBiz(storeNotify)
		notify := &modelNotify.CreateNotify{
			UserId:      item.UserId,
			ItemId:      &id,
			CreatorId:   request.GetUserId(),
			Message:     fmt.Sprintf("%s has been like regarding your item ", item.Owner.LastName),
			TypeMessage: modelNotify.TypeLikeItem,
		}
		if err := businessNotify.CreateNotify(c.Request.Context(), notify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
