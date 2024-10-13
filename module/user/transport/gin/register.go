package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/module/user/biz"
	"main.go/module/user/model"
	"main.go/module/user/storage"
	"net/http"
)

func Register(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data model.CreateUser
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		S := model.StatusUserDoing
		data.Status = &S
		store := storage.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewRegisterBiz(store, hash)
		verify, errVerify := business.NewRegister(c.Request.Context(), &data)
		if verify == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errVerify.Error()})
			return
		} else if errVerify != nil {
			c.JSON(http.StatusOK, gin.H{"data": verify, "err": errVerify.Error()})

		}
		c.JSON(http.StatusOK, gin.H{"data": verify})
	}
}
