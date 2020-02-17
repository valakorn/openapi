package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valakorn/openapi/platform/newsfeed"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)
		//============Unmarshal==================

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		//c.Status(http.StatusNoContent)
		c.JSON(http.StatusOK, gin.H{"status": "Passed"})
	}
}

func NewsfeedPostV1(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newsFeed newsfeedPostRequest

		if c.ShouldBind(&newsFeed) == nil {
			log.Println(newsFeed.Title)
			log.Println(newsFeed.Post)
		}

		today := time.Now()
		c.JSON(http.StatusOK, gin.H{"status": "Passed", "Time": today})

		//========================Update Data In DB==================================
		var lines = []string{newsFeed.Title}
		err := writeLines(lines, "DBTEXT.txt")
		//err := writeLines(lines, "DBTEXT.txt")
		if err != nil {
			panic(err)
		}

		//============================================================================

	}
}
