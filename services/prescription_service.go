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
		PatientInfoId: int(pIdInt),
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

func HandleFindPresciption(c *gin.Context) {
	fmt.Printf("########################## ")
	pId, ok := c.GetQuery("patientInfoId")
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
		PatientInfoId: int(pIdInt),
	}
	prescriptions, err := dao.Find(m, 0, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":      1002,
			"error_msg": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "prescription", gin.H{
		"prescriptions": prescriptions,
		"pname":         "nihao",
	})
}

func HandleCreatePresciption(c *gin.Context) {
	fmt.Printf("########################## ")
}

func HandleDeletePresciption(c *gin.Context) {
	fmt.Printf("########################## ")
	id, ok := c.GetPostForm("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code":         1002,
			"error_mesage": "no have id",
		})
		return
	}
	pIdInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":      1002,
			"error_msg": err.Error(),
		})
		return
	}

	p := models.Prescription{
		Id: int(pIdInt),
	}
	dao := dao.NewPrescriptionDao()
	err = dao.Del(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":         1002,
			"error_mesage": "del failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    100,
		"message": "del success",
	})
}
