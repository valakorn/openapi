package main

import (
	"github.com/gin-gonic/gin"
	"github.com/valakorn/openapi/httpd/handler"
)

func main() {
	//fmt.Println("test")

	r := gin.Default()
	r.GET("/", handler.HomePage)
	r.POST("/", handler.PostHomePage)
	r.GET("/query", handler.QueryStrings) // query?name=sunsu&age=24
	r.GET("/ping", handler.PingGet())
	r.Run()
}
