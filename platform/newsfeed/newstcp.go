package newsfeed

type tcpAdder interface {
	Add(atmrequesttransfer ATMrequestTransfer)
}

type ATMrequestTransfer struct {
	No1    string `json:"no1" binding:"required"`
	Date1  string `json:"date1" binding:"required"`
	No2    string `json:"no2" binding:"required"`
	Date2  string `json:"date2" binding:"required"`
	No3    string `json:"no3" binding:"required"`
	Date3  string `json:"date3" binding:"required"`
	Date4  string `json:"date4" binding:"required"`
	No4    string `json:"no4" binding:"required"`
	Rquid1 string `json:"rquid1" binding:"required"`
	No5    string `json:"no5" binding:"required"`
	Brcd   string `json:"brcd" binding:"required"`
	No6    string `json:"no6" binding:"required"`
	Rquid2 string `json:"rquid2" binding:"required"`
	No7    string `json:"no7" binding:"required"`
	Accid1 string `json:"accid1" binding:"required"`
	No8    string `json:"no8" binding:"required"`
	Accid2 string `json:"accid2" binding:"required"`
	No9    string `json:"no9" binding:"required"`
}

type tcpRepo struct {
	ATMrequestTransfers []ATMrequestTransfer
}

func tcpNew() *tcpRepo {
	return &tcpRepo{
		ATMrequestTransfers: []ATMrequestTransfer{},
	}
}

func (r *tcpRepo) Add(atmrequesttransfer ATMrequestTransfer) {
	r.ATMrequestTransfers = append(r.ATMrequestTransfers, atmrequesttransfer)
}

func (r *tcpRepo) tcpGetAll() []ATMrequestTransfer {
	return r.ATMrequestTransfers
}
