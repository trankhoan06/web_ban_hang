package ginUserlikeItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/module/userlikeitem/biz"
	"main.go/module/userlikeitem/storage"
	"net/http"
	"strconv"
)

func ListLike(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlModel(db)
		business := biz.NewListLikeBiz(store)
		data, err := business.NewListLike(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"amount": len(*data),
			"data":   data,
		})
	}
}
