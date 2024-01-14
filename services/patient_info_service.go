package services

import (
	"consultation-service/dao"
	"consultation-service/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetAllPatients(c *gin.Context) {
	dao := dao.PatientInfoDao{}
	name, ok := c.GetPostForm("Name")
	if ok && len(name) > 0 {
		// dao.Find
		m := models.PatientInfo{
			Name: name,
		}

		result, err := dao.Find(m, 0, 0)
		if err != nil {
			fmt.Printf("err:%v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_code": 1002,
				"error_msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"patients": result,
		})
		return
	}

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

type PatientParam struct {
	Id      int    `form:"id"`
	Name    string `form:"name"`
	Age     int    `form:"age"`
	Sex     int    `form:"sex"`
	Address string `form:"address"`
	Tel     string `form:"tel"`
}

func UpdatePatients(c *gin.Context) {
	p := PatientParam{}
	err := c.ShouldBind(&p)
	fmt.Printf("err:%v, p=%+v\n", err, p)
	if err != nil {
		c.JSON(http.StatusPermanentRedirect, gin.H{
			"code":          123,
			"error_message": err.Error(),
		})
		return
	}
	dao := dao.PatientInfoDao{}
	patient := models.PatientInfo{
		Id:      uint(p.Id),
		Name:    p.Name,
		Age:     uint(p.Age),
		Sex:     uint(p.Sex),
		Address: p.Address,
		Tel:     p.Tel,
	}
	err = dao.Update(&patient)
	c.JSON(http.StatusOK, gin.H{
		"code": 100,
		"data": patient,
	})
}
