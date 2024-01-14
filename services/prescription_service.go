package services

import (
	"consultation-service/dao"
	"consultation-service/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetPresciption(c *gin.Context) {
	fmt.Printf("########################## ")
	pId, ok := c.GetPostForm("id")
	fmt.Printf("pid:%s, ok:%t\n", pId, ok)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":      1002,
			"error_msg": "no have patinent id",
		})
		return
	}
	//
	pIdInt, err := strconv.ParseInt(pId, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":      1002,
			"error_msg": err.Error(),
		})
		return
	}
	dao := dao.NewPrescriptionDao()
	m := models.Prescription{
		PatientInfoId: uint64(pIdInt),
	}
	err = dao.GetOne(&m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":      1002,
			"error_msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"prescription": m,
	})
}
