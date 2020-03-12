package newsfeed

type TcpAdder interface {
	Add(itemtcp Itemtcp)
}

type Itemtcp struct {
	No1    string `json:"no1"`
	Date1  string `json:"date1"`
	No2    string `json:"no2"`
	Date2  string `json:"date2"`
	No3    string `json:"no3"`
	Date3  string `json:"date3"`
	Date4  string `json:"date4"`
	No4    string `json:"no4"`
	Rquid1 string `json:"rquid1"`
	No5    string `json:"no5"`
	Brcd   string `json:"brcd"`
	No6    string `json:"no6"`
	Rquid2 string `json:"rquid2"`
	No7    string `json:"no7"`
	Accid1 string `json:"accid1"`
	No8    string `json:"no8"`
	Accid2 string `json:"accid2"`
	No9    string `json:"no9"`
}

type TcpRepo struct {
	Itemtcps []Itemtcp
}

func TcpNew() *TcpRepo {
	return &TcpRepo{
		Itemtcps: []Itemtcp{},
	}
}

func (r *TcpRepo) Add(itemtcp Itemtcp) {
	r.Itemtcps = append(r.Itemtcps, itemtcp)
}

func (r *TcpRepo) TcpGetAll() []Itemtcp {
	return r.Itemtcps
}
