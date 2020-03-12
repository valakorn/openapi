package newsfeed

type CbsinfoservicesItemAdder interface {
	Add(cbsinfoservicesItem CbsinfoservicesItem)
}

type CbsinfoservicesOfRequestHeader struct {
	ChannelId string `json:"channelId"`
	TransCd   string `json:"transCd"`
}

type CbsinfoservicesItem struct {
	CbsinfoservicesOfRequestHeader CbsinfoservicesOfRequestHeader `json:"cbsinfoservicesOfRequestHeader"`
	FuncFlag                       string                         `json:"funcFlag"`
	EffDate                        string                         `json:"effDate"`
	CIFNo                          string                         `json:"cIFNo"`
	AcctNo                         string                         `json:"acctNo"`
	MobileNo                       string                         `json:"mobileNo"`
	RefNo                          string                         `json:"refNo"`
	TransAmt                       string                         `json:"transAmt"`
	ExpireTime                     string                         `json:"expireTime"`
}

type CbsinfoservicesItemRepo struct {
	CbsinfoservicesItems []CbsinfoservicesItem
}

func CbsinfoservicesItemRepoNew() *CbsinfoservicesItemRepo {
	return &CbsinfoservicesItemRepo{
		CbsinfoservicesItems: []CbsinfoservicesItem{},
	}
}

func (r *CbsinfoservicesItemRepo) Add(cbsinfoservicesItem CbsinfoservicesItem) {
	r.CbsinfoservicesItems = append(r.CbsinfoservicesItems, cbsinfoservicesItem)
}

func (r *CbsinfoservicesItemRepo) CbsinfoservicesItemRepoGetAll() []CbsinfoservicesItem {
	return r.CbsinfoservicesItems
}
