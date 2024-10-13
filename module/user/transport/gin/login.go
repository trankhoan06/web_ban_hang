package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/common"
	"main.go/component/tokenProvider"
	"main.go/module/user/biz"
	"main.go/module/user/model"
	"main.go/module/user/storage"
	"net/http"
)

func Login(db *gorm.DB, provider tokenProvider.Provider) func(*gin.Context) {
	return func(c *gin.Context) {
		var login model.LoginUser
		err := c.ShouldBind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store:=storage.NewSqlModel(db)
		hash:=common.NewSha256Hash()
		business:=biz.NewLoginBiz(store,hash, provider)
		verify, err:=business.NewLogin(c.Request.Context(), &login, 60*60*30*24)
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": verify})

	}
}