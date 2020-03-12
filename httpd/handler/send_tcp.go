package handler

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/valakorn/openapi/platform/newsfeed"
)

type atmrequestTransfer struct {
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
	ip   = "10.9.182.131"
	port = 27001
)

func Sendtcp(feed newsfeed.TcpAdder) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody atmrequestTransfer

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		} else {

			output := requestBody.No1 + requestBody.Date1 +
				requestBody.No2 + requestBody.Date2 +
				requestBody.No3 + requestBody.Date3 +
				requestBody.Date4 + requestBody.No4 +
				requestBody.Rquid1 + requestBody.No5 +
				requestBody.Brcd + requestBody.No6 +
				requestBody.Rquid2 + requestBody.No7 +
				requestBody.Accid1 + requestBody.No8 +
				requestBody.Accid2 + requestBody.No9

			// var (
			// 	ip   = "10.9.182.131"
			// 	port = 27002
			// )
			newport := randomInt(27002, 27060) //

			countlen := len(output)
			covhex := fmt.Sprintf("%04x", countlen)
			decode, _ := hex.DecodeString(covhex)

			fmt.Println(newport)

			outputtcp := SocketClient(ip, newport, string(decode)+output)
			c.JSON(http.StatusOK, outputtcp)
			fmt.Println(substringTcp(output))

		}
	}
}
func SocketClient(ip string, port int, msg string) string { //[]byte
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	var x string
	fmt.Println(msg)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer conn.Close()
	conn.Write([]byte(msg))
	buff := make([]byte, 683) //683
	n, _ := conn.Read(buff)
	outputtcp := buff[:n]
	x = string(outputtcp)

	return x

}
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func substringTcp(resp string) string {
	// a := []rune(resp)
	//var queue []Queue
	para := [26]int{32, 16, 16, 6, 12, 10, 6, 6, 4, 4, 4}

	//para := [26]int{32, 16, 16, 6, 12, 10, 6, 6, 4, 4, 4, 3, 2, 9, 4, 37, 12, 16, 15, 40, 44, 3, 15, 12, 16, 50}

	n := 0
	stringxx := "xxxxx"

	for r, c := range para {
		switch r {
		case 0:
			mti := resp[n : c+n]
			n = n + c

			fmt.Println("mti: ", mti)

		case 1:
			secondarybitmap := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 1: ", secondarybitmap)

		case 2:
			processingCode := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 2: ", processingCode)

		case 3:
			transactionAmount := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 3: ", transactionAmount)

		case 4:
			transmissionDateandTime := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 4: ", transmissionDateandTime)

		case 5:
			systemTraceAuditNumber := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 7: ", systemTraceAuditNumber)

		case 6:
			localTransactionTime := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 11: ", localTransactionTime)
		case 7:
			localTransactionDate := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 12: ", localTransactionDate)

		case 8:
			settlement := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 13: ", settlement)
		case 9:
			captureDate := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 17: ", captureDate)

		case 10:
			mcc := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 18: ", mcc)
		case 11:
			cardSequenceNumber := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 22: ", cardSequenceNumber)
		case 12:
			pointofServiceConditionCode := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 25: ", pointofServiceConditionCode)
		case 13:
			transactionFeeAmount := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 28: ", transactionFeeAmount)
		case 14:
			acquiringInstitutionID := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 32: ", acquiringInstitutionID)

		case 15:
			track2Data := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 35: ", track2Data)

		case 16:
			retrievalReferenceNumber := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 37: ", retrievalReferenceNumber)

		case 17:
			authorisationIDResponse := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 41: ", authorisationIDResponse)

		case 18:
			responseCode := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 42: ", responseCode)

		case 19:
			cardAcceptorTerminalID := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 43: ", cardAcceptorTerminalID)

		case 20:
			cardAcceptorIdentificationCode := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 48: ", cardAcceptorIdentificationCode)

		case 21:
			cardAcceptorNameLocation := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 49: ", cardAcceptorNameLocation)

		case 22:
			additionalResponseData := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 60: ", additionalResponseData)

		case 23:
			accountID1 := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 85: ", accountID1)

		case 24:
			accountID2 := resp[n : c+n]
			n = n + c
			fmt.Println("Bit 102: ", accountID2)
		case 25:
			bit25 := resp[n:]
			n = n + c
			fmt.Println("Bit 126: ", bit25)

		default:
			//fmt.Println("pass")
		}

		//fmt.Println(r, c)
	}

	// mti := resp[n:4]
	// n = n + 4
	// primaryBitmap := resp[n : 16+n]
	// n = n + 16
	// secondarybitmap := resp[n : 16+n]
	// n = n + 16
	// processingCode := resp[n : 6+n]
	// n = n + 6
	// transactionAmount := resp[n : 12+n]

	fmt.Println(para)

	//n = 0
	return (stringxx)

}

// func IntToHex(num int64) []byte {
// 	buff := new(bytes.Buffer)
// 	err := binary.Write(buff, binary.BigEndian, num)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return buff.Bytes()
// }

// func Splithex(hex string) string {
// 	//var v int
// 	//var c float64 = '\x01'

// 	str := hex
// 	if len(str) < 2 {
// 		fmt.Println(str)
// 	}
// 	var strBuf bytes.Buffer // A Buffer needs no initialization.
// 	indx := str[0]

// 	strBuf.WriteByte(indx)
// 	for _, count := range []byte(str[1:]) {
// 		if count != indx {
// 			strBuf.WriteString(" ")
// 		}
// 		strBuf.WriteByte(count)
// 		indx = count
// 	}

// 	output := strBuf.String()
// 	s := strings.Split(output, " ")
// 	fmt.Println(s)

// 	//msghex := '\x00'
// 	//ascii := int(msghex)
// 	//fmt.Println(msghex)

// 	//msghexrep := strings.Replace(string(msghex), "", "", -1)
// 	fmt.Println(selectFArrays(s[1]))

// 	output = selectFArrays(s[1]) + selectFArrays(s[2])

// 	return output

// }
// func selectFArrays(carry string) string {
// 	switch carry {
// 	case "1":
// 		carry = "\x01"
// 	case "2":
// 		carry = "\x02"
// 	case "3":
// 		carry = "\x03"
// 	case "4":
// 		carry = "\x04"
// 	case "5":
// 		carry = "\x05"
// 	case "6":
// 		carry = "\x06"
// 	case "7":
// 		carry = "\x07"
// 	case "8":
// 		carry = "\x08"
// 	case "9":
// 		carry = "\x09"
// 	case "0":
// 		carry = "\x00"

// 	case "a":
// 		carry = "\xa0"
// 	case "b":
// 		carry = "\xb0"
// 	case "c":
// 		carry = "\xc0"
// 	case "d":
// 		carry = "\xd0"
// 	case "e":
// 		carry = "\xe0"
// 	case "f":
// 		carry = "\xf0"

// 	default:
// 		//fmt.Println("It's a weekday")
// 	}
// 	return carry
// }

// func response_tcp(c *gin.Context, httpCode, code int, data interface{}) {
// 	c.JSON(httpCode, gin.H{
// 		"code": code,
// 		"data": data,
// 	})
// 	return
// }

//==============Backup Function ==============================================
// func Sendtcp(feed newsfeed.TcpAdder) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		//requestBody := newsfeedPostRequest{}
// 		//var newsfeed atmrequestTransfer

// 		//var newsFeed atmrequestTransfer
// 		//requestBody := atmrequestTransfer{}
// 		var requestBody atmrequestTransfer

// 		//c.Bind(&requestBody)

// 		if err := c.ShouldBindJSON(&requestBody); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		} else {
// 			//c.Bind(&requestBody)
// 			output := requestBody.No1 + requestBody.Date1 +
// 				requestBody.No2 + requestBody.Date2 +
// 				requestBody.No3 + requestBody.Date3 +
// 				requestBody.Date4 + requestBody.No4 +
// 				requestBody.Rquid1 + requestBody.No5 +
// 				requestBody.Brcd + requestBody.No6 +
// 				requestBody.Rquid2 + requestBody.No7 +
// 				requestBody.Accid1 + requestBody.No8 +
// 				requestBody.Accid2 + requestBody.No9

// 			//fmt.Println(output)

// 			var (
// 				ip   = "10.9.182.131"
// 				port = 27002
// 			)
// 			//IntToHex(int64(targetBits)),
// 			//fmt.print(int64(output))
// 			countlen := len(output)
// 			//fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
// 			//covhex := fmt.Sprintf("%04X", countlen)
// 			covhex := fmt.Sprintf("%04x", countlen)
// 			decode, _ := hex.DecodeString(covhex)

// 			// fmt.Printf("Hex conv of '%d' is '%s'\n", countlen, covhex)
// 			//========================================
// 			// countlen64 := int64(countlen)
// 			// fmt.Println(Splithex(string(covhex)))
// 			//	gethexnum := Splithex(string(covhex))
// 			//gethexnum := "\x01\xF0"
// 			//fmt.Println(gethexnum)

// 			//outputtcp := SocketClient(ip, port, Splithex(string(covhex))+output)
// 			outputtcp := SocketClient(ip, port, string(decode)+output)
// 			//outputtcp := SocketClient(ip, port, output)
// 			c.JSON(http.StatusOK, outputtcp)

// 		}
// 	}
// }

//==================================Note const Parameter ===================================================
// const (
// 	//message:= make([]byte, 496)
// 	//message = "\x01\xF0"
// 	//message       = "\x01\xF0ISO0160000700200F23EE49528E080000000080006000004164499320010003830400000000000000060022010274304276410274302200908022002206011000000  00000000000000000009123456789374499320010003830=23081261744963236801553962174296K53400          0006           001265KKKKK THAI BANK LOAD TEST    010TH764553911371229166780071064      169804103923      260& 0000300260! Q100124                     00060006     0000000000000000000000000000000000000000000000000000000000000000000000000000FM20! QC00104   VC00000000"

// 	ip   = "10.9.182.131"
// 	port = 27002
// 	//StopCharacter = "\r\n\r\n"
// )
//=================================Backup Sent Msg ==========================================================
// func Sendtcp(feed newsfeed.TcpAdder) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var requestBody atmrequestTransfer

// 		if err := c.ShouldBindJSON(&requestBody); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		} else {

// 			output := requestBody.No1 + requestBody.Date1 +
// 				requestBody.No2 + requestBody.Date2 +
// 				requestBody.No3 + requestBody.Date3 +
// 				requestBody.Date4 + requestBody.No4 +
// 				requestBody.Rquid1 + requestBody.No5 +
// 				requestBody.Brcd + requestBody.No6 +
// 				requestBody.Rquid2 + requestBody.No7 +
// 				requestBody.Accid1 + requestBody.No8 +
// 				requestBody.Accid2 + requestBody.No9

// 			// var (
// 			// 	ip   = "10.9.182.131"
// 			// 	port = 27002
// 			// )

// 			countlen := len(output)
// 			covhex := fmt.Sprintf("%04x", countlen)
// 			decode, _ := hex.DecodeString(covhex)

// 			outputtcp := SocketClient(ip, port, string(decode)+output)
// 			c.JSON(http.StatusOK, outputtcp)

// 		}
// 	}
// }

//=======================================Backup TCP=============================================
// func SocketClient(ip string, port int, msg string) string { //[]byte
// 	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
// 	conn, err := net.Dial("tcp", addr)
// 	var x string
// 	//msg = message
// 	fmt.Println(msg)

// 	if err != nil {
// 		log.Fatalln(err)
// 		os.Exit(1)
// 	}

// 	defer conn.Close()

// 	//s := make([]byte, 5)

// 	conn.Write([]byte(msg))
// 	// conn.Write([]byte(message + msg))
// 	//conn.Write([]byte(StopCharacter))
// 	//log.Printf("Send: %s", message)

// 	buff := make([]byte, 683)
// 	n, _ := conn.Read(buff)
// 	outputtcp := buff[:n]
// 	x = string(outputtcp)
// 	//log.Printf("Receive: %s", x)

// 	return x

// }
