package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var token = "1234567"

func HandleLogin(c *gin.Context) {
	username, ok := c.GetPostForm("username")
	fmt.Printf(">>>>:%v ok=%t\n", username, ok)
	if !ok || len(username) <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1001, "message": "no have user name"})
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "token": token, "message": "login success"})
}

func HandleLogout(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"code": 200, "message": "logout success"})
	c.Redirect(http.StatusMovedPermanently, "/index")
}
