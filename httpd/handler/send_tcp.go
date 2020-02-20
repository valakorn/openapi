package handler

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ATMrequestTransfer struct {
	No1    string `json:"no1" binding:"required"`    //ISO0160000700200F23EE49528E080000000080006000004164499320010003830400000000000000060
	Date1  string `json:"date1" binding:"required"`  //0220
	No2    string `json:"no2" binding:"required"`    //102743042764102743
	Date2  string `json:"date2" binding:"required"`  //0220
	No3    string `json:"no3" binding:"required"`    //0908
	Date3  string `json:"date3" binding:"required"`  //0220
	Date4  string `json:"date4" binding:"required"`  //0220
	No4    string `json:"no4" binding:"required"`    //6011000000  00000000000000000009123456789374499320010003830=23081261744963236801
	Rquid1 string `json:"rquid1" binding:"required"` //403000000001{SPR_3}{Random}
	No5    string `json:"no5" binding:"required"`    //K53400          0006
	Brcd   string `json:"brcd" binding:"required"`   //000456
	No6    string `json:"no6" binding:"required"`    //KKKKK THAI BANK LOAD TEST    010TH764
	Rquid2 string `json:"rquid2" binding:"required"` //403000000001
	No7    string `json:"no7" binding:"required"`    //16
	Accid1 string `json:"accid1" binding:"required"` //4560093393
	No8    string `json:"no8" binding:"required"`    //16
	Accid2 string `json:"accid2" binding:"required"` //8090480160
	No9    string `json:"no9" binding:"required"`    //260& 0000300260! Q100124                     00060006     0000000000000000000000000000000000000000000000000000000000000000000000000000FM20! QC00104   VC00000000

}

const (
	//message:= make([]byte, 496)
	message       = "\x01\xF0ISO0160000700200F23EE49528E080000000080006000004164499320010003830400000000000000060022010274304276410274302200908022002206011000000  00000000000000000009123456789374499320010003830=23081261744963236801553962174296K53400          0006           001265KKKKK THAI BANK LOAD TEST    010TH764553911371229166780071064      169804103923      260& 0000300260! Q100124                     00060006     0000000000000000000000000000000000000000000000000000000000000000000000000000FM20! QC00104   VC00000000"
	StopCharacter = "\r\n\r\n"
)

func Sendtcp(feed newsfeed.tcpAdder) gin.HandlerFunc {
	return func(c *gin.Context) {
		var atmrequestransfer ATMrequestTransfer
		//var newsFeed ATMrequestTransfer

		// if c.ShouldBindQuery(&atmrequestransfer) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(atmrequestransfer.No1)
		log.Println(atmrequestransfer.Date1)

		var (
			ip   = "10.9.182.131"
			port = 27002
		)
		outputtcp := SocketClient(ip, port)
		c.JSON(http.StatusOK, outputtcp)

		// }
	}
}
func SocketClient(ip string, port int) string { //[]byte
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	var x string

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer conn.Close()

	//s := make([]byte, 5)

	conn.Write([]byte(message))
	//conn.Write([]byte(StopCharacter))
	//log.Printf("Send: %s", message)

	buff := make([]byte, 683)
	n, _ := conn.Read(buff)
	outputtcp := buff[:n]
	x = string(outputtcp)
	//log.Printf("Receive: %s", x)

	return x

}

// func response_tcp(c *gin.Context, httpCode, code int, data interface{}) {
// 	c.JSON(httpCode, gin.H{
// 		"code": code,
// 		"data": data,
// 	})
// 	return
// }
