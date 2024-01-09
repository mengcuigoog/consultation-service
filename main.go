package main

import (
	"consultation-service/dao"
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("###:%s\n", r.Host)
	w.Write([]byte("<h1>Hello World</h1>"))
}

func main() {
	http.HandleFunc("/hello", Hello)
	fmt.Println("start at [:9090]")
	// demo()
	dao.Init()
	err := http.ListenAndServe(":9090", nil)
	// err := http.ListenAndServeTLS(":9090", "", "", nil)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		panic("start server failed")
	}
}
