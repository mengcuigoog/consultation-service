package plugins

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("强制退出")
		token := c.GetHeader("token")
		fmt.Printf("token:%s len:%d\n", token, len(token))
		if len(token) <= 0 || token == "null" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "未登录",
			})
			c.Abort()
			return
		}
		c.Set("token", token)
	}
}
