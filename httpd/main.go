package main

import (
	"github.com/gin-gonic/gin"
	"github.com/valakorn/openapi/httpd/handler"
	"github.com/valakorn/openapi/platform/newsfeed"
)

func main() {

	//fmt.Println("test")

	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/", handler.HomePage)
	r.POST("/", handler.PostHomePage)
	r.GET("/query", handler.QueryStrings) // query?name=sunsu&age=24
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	r.Run()

	// feed := newsfeed.New()
	// fmt.Println(feed)
	// feed.Add(newsfeed.Item{"Hello", "How ya doing mate?"})
	// fmt.Println(feed)
}
