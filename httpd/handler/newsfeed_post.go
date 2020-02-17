package handler

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valakorn/openapi/platform/newsfeed"
)

type newsfeedPostRequest struct {
	Title string `json:"title" binding:"required"`
	Post  string `json:"post" binding:"required"`
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

		// if err := c.ShouldBindJSON(&newsFeed); err == nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		// }

		if err := c.ShouldBindJSON(&newsFeed); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		} else {
			// if c.ShouldBind(&newsFeed) == nil {
			// 	log.Println(newsFeed.Title)
			// 	log.Println(newsFeed.Post)
			// }

			today := time.Now()
			c.JSON(http.StatusOK, gin.H{"status": "Passed", "Time": today})
			//fmt.Println(today.Format("2020-01-02 15:04:05").String())
			//fmt.Println(today.String())
			day_string := today.Format("2006-01-02 15:04:05")

			//========================Update Data In DB==================================
			var lines = []string{newsFeed.Title, "|", newsFeed.Post, "|", day_string}
			err := writeLines_main(lines, "D:\\Job\\go_module\\openapi\\httpd\\DBTEXT.txt")
			//err := writeLines(lines, "DBTEXT.txt")
			if err != nil {
				panic(err)
			}

			//============================================================================

		}

	}
}

func writeLines_main(lines []string, path string) error {
	// overwrite file if it exists
	//file, err := os.OpenFile("./DBFILE.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	file, err := os.OpenFile(path, os.O_APPEND, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	_, err = fmt.Fprintln(file, lines)
	if err != nil {
		return err
	}

	return w.Flush()
}
