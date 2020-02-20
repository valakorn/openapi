package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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

type RegisterOfpRespHeader struct {
	ResponseCode int    `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
	WsRefID      string `json:"wsRefId"`
}

type NewsfeedPostRep struct {
	RegisterOfpRespHeader RegisterOfpRespHeader `json:"registerOfpRespHeader"`
	Title                 string                `json:"title"`
	Post                  string                `json:"post" binding:"required"`
	Timestamp             string                `json:"timestamp"`
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
	//return func(c *gin.Context) {
	return func(c *gin.Context) {
		var newsFeed newsfeedPostRequest
		var x interface{}
		//var v interface{}
		//var newsfeedpostrep []NewsfeedPostRep
		//var x interface{}
		//var Getnewsfeedpostrep
		//var registerofprespheader []RegisterOfpRespHeader

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
			day_string := today.Format("2006-01-02 15:04:05")

			//newsfeedpostrep = append(newsfeedpostrep, NewsfeedPostRep{Timestamp: day_string})

			//registerofprespheader = append(registerofprespheader, RegisterOfpRespHeader{ResponseCode: 10})
			//newsfeedpostr
			//fmt.Printf("%q\n", newsfeedpostrep)
			//jsonPessoal, errr := json.Marshal(newsfeedpostrep)
			//jsonPessoal, errr := json.Marshal(NewsfeedPostRep{Timestamp: day_string}, RegisterOfpRespHeader{ResponseCode: 10})
			//result := strings.Join(newsfeedpostrep, ",")
			//func Marshal(v interface{}) ([]byte, error) {
			jsonPessoal, errr := json.Marshal(NewsfeedPostRep{Timestamp: day_string})
			//jsonPessoal, errr := json.Marshal(dataposrep)
			if errr != nil {
				log.Fatal(err)
			}

			//fmt.Fprintf(os.Stdout, "%s", jsonPessoal)
			//c.JSON(http.StatusOK, jsonPessoal)
			//JSON(code int, obj interface{}

			x = string(jsonPessoal)
			c.JSON(http.StatusOK, gin.H{"message": x})
			//c.JSON(http.StatusOK, x)
			//response(c, http.StatusOK, 200, "helloword")

			//fmt.Fprint(w, string(jsonPessoal)) // write response to ResponseWriter (w)

			//foo_marshalled, err := json.Marshal(Foo{Number: 1, Title: "test"})

			//fmt.Fprintf(os.Stdout, "%s", jsonPessoal)
			//c.Header()

			// c.JSON(http.StatusOK, gin.H{
			// 	"code":    http.StatusOK,
			// 	"message": string(jsonPessoal), // cast it to string before showing
			// })
			//c.JSON(http.StatusOK, gin.H{"status": "Passed", "Time": today})
			//fmt.Println(today.Format("2020-01-02 15:04:05").String())
			//fmt.Println(today.String())

			//========================Update Data In DB==================================
			var lines = []string{newsFeed.Title, "|", newsFeed.Post, "|", day_string}
			//err := writeLines_main(lines, "D:\\Job\\go_module\\openapi\\httpd\\DBTEXT.txt")
			err := writeLines_main(lines, "./DBTEXT.txt")
			//err := writeLines(lines, "DBTEXT.txt")
			if err != nil {
				panic(err)
			}

			//============================================================================

		}

	}
}

func response(c *gin.Context, httpCode, code int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"data": data,
	})
	return
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
