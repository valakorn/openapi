package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valakorn/openapi/platform/newsfeed"
)

type CbsinfoservicesOfRequestHeader struct {
	ChannelId string `json:"channelId" binding:"required"`
	TransCd   string `json:"transCd" binding:"required"`
}

type CbsinfoservicesItem struct {
	CbsinfoservicesOfRequestHeader CbsinfoservicesOfRequestHeader `json:"cbsinfoservicesOfRequestHeader"`
	FuncFlag                       string                         `json:"funcFlag" binding:"required"`
	EffDate                        time.Time                      `json:"effDate" binding:"required" time_format:"2006-01-02"`
	CIFNo                          string                         `json:"cIFNo" binding:"required"`
	AcctNo                         string                         `json:"acctNo" binding:"required"`
	MobileNo                       string                         `json:"mobileNo" binding:"required"`
	RefNo                          string                         `json:"refNo" binding:"required"`
	TransAmt                       string                         `json:"transAmt" binding:"required"`
	ExpireTime                     string                         `json:"expireTime" binding:"required"`
}

type CbsinfoservicesItemPostRep struct {
	ResponseCode int    `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
}

func Func_cbsinfoservices(feed newsfeed.CbsinfoservicesItemAdder) gin.HandlerFunc {
	//return func(c *gin.Context) {
	return func(c *gin.Context) {
		var newsFeed CbsinfoservicesItem
		//var x interface{}

		if err := c.ShouldBindJSON(&newsFeed); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			//today := time.Now()
			//day_string := today.Format("2006-01-02 15:04:05")

			res := CbsinfoservicesItemPostRep{ResponseDesc: "Test"}

			jsonPessoal, _ := json.Marshal(res)

			personMap := make(map[string]interface{})
			errunma := json.Unmarshal(jsonPessoal, &personMap)
			if errunma != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, personMap)

			//========================Update Data In DB==================================

			//============================================================================

		}

	}
}
