package ginComment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	emailSend "main.go/email"
	"net/http"
)

func SendEmail(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		email := c.Param("email")
		sender := emailSend.NewGmailSender()
		title := "A test email"
		content := `
<h1>Hello word</h1>
<p>This is a test email.<a href="facebook.com"> facebook</a></p>
`
		to := []string{email}
		attackFile := []string{"./main.go"}
		err := sender.SendEmail(title, content, to, nil, nil, attackFile)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
