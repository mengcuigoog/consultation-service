package main

import (
	"consultation-service/dao"
	"consultation-service/plugins"
	"consultation-service/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	defaultEngine := gin.Default()
	defaultEngine.Static("/statics/", "./statics/")
	defaultEngine.LoadHTMLGlob("templates/*")

	defaultEngine.GET("/index", services.HandleIndex)
	defaultEngine.POST("/login", services.HandleLogin)
	// 需要登录才能访问的接口
	authGroup := defaultEngine.Group("/").Use(plugins.HandleAuth())
	{
		authGroup.POST("/logout", services.HandleLogout)
		authGroup.POST("/patients", services.GetAllPatients)
		authGroup.POST("/updatepatients", services.UpdatePatients)

		authGroup.POST("/getpatientsdetails", services.HandleGetPresciption)
	}

	fmt.Println("start at [:9090]")
	dao.Init()

	err := defaultEngine.Run(":9090")

	if err != nil {
		fmt.Printf("err:%v\n", err)
		panic("start server failed")
	}
}
