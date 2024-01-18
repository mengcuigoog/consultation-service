package main

import (
	"consultation-service/dao"
	"consultation-service/plugins"
	"consultation-service/services"
	"fmt"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func newRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("index", "templates/base.html", "templates/patients/index.html")
	r.AddFromFiles("presciption", "templates/base.html", "templates/presciption/index.html")
	return r
}

func main() {
	defaultEngine := gin.Default()
	defaultEngine.HTMLRender = newRender()
	defaultEngine.Static("/statics/", "./statics/")
	// defaultEngine.LoadHTMLGlob("templates/*")
	defaultEngine.GET("/index", services.HandleIndex)
	defaultEngine.POST("/login", services.HandleLogin)
	// 需要登录才能访问的接口
	authGroup := defaultEngine.Group("/").Use(plugins.HandleAuth())
	{
		authGroup.POST("/logout", services.HandleLogout)
		authGroup.POST("/patients", services.GetAllPatients)
		authGroup.POST("/updatepatients", services.UpdatePatients)

		authGroup.GET("/patientsdetails", services.HandleFindPresciption)
		authGroup.POST("/createpatientsdetail", services.HandleCreatePresciption)

		authGroup.POST("/delPresciption", services.HandleDeletePresciption)
	}

	fmt.Println("start at [:9090]")
	dao.Init()

	err := defaultEngine.Run(":9090")

	if err != nil {
		fmt.Printf("err:%v\n", err)
		panic("start server failed")
	}
}
