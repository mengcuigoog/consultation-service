package services

import (
	"consultation-service/dao"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetAllPatients(c *gin.Context) {
	dao := dao.PatientInfoDao{}
	all, err := dao.All()
	if err != nil {
		fmt.Printf("err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_code": 1002,
			"error_msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"patients": all,
	})
}
