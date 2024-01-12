package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientInfoService struct {
}

func HandleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
