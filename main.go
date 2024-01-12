package main

import (
	"consultation-service/dao"
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
	defaultEngine.POST("/logout", services.HandleLogout)

	fmt.Println("start at [:9090]")
	dao.Init()

	err := defaultEngine.Run(":9090")

	if err != nil {
		fmt.Printf("err:%v\n", err)
		panic("start server failed")
	}
}
