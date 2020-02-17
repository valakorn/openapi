package handler

import (
	"bufio"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/valakorn/openapi/platform/newsfeed"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}

func NewsfeedGetv1(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		//results := feed.GetAll()
		//results := []string{"learn golang", "practise execercise", "make coffee"}
		results := []string{}
		results = readLines("D:\\Job\\go_module\\openapi\\httpd\\DBTEXT.txt")
		//results.Execute(w, readLines("todos.txt"))

		c.JSON(http.StatusOK, results)
	}
}

func readLines(name string) []string {
	f, err := os.Open(name)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return lines
}
