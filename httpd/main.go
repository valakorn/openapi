package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valakorn/openapi/httpd/handler"
	"github.com/valakorn/openapi/platform/newsfeed"
)

func main() {

	//fmt.Println("test")

	feed := newsfeed.New()

	router := gin.Default()

	router.GET("/", handler.HomePage)
	router.POST("/", handler.PostHomePage)
	router.GET("/query", handler.QueryStrings) // query?name=sunsu&age=24
	router.GET("/ping", handler.PingGet())
	router.GET("/newsfeed", handler.NewsfeedGet(feed))
	router.GET("/newsfeedv1", handler.NewsfeedGetv1(feed))
	router.POST("/newsfeedv1", handler.NewsfeedPostV1(feed))
	router.POST("/newsfeed", handler.NewsfeedPost(feed))

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	//r.Run(":8080")

	// feed := newsfeed.New()
	// fmt.Println(feed)
	// feed.Add(newsfeed.Item{"Hello", "How ya doing mate?"})
	// fmt.Println(feed)
}
