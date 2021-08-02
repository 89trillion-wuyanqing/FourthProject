package model

type Gifts struct {
	ID  int `json:"id"`
	Num int `json:"num"`
}

type Record struct {
	Userid      string `json:"userid"`
	PullTime    int64  `json:"pullTime"`
	PullTimeStr string `json:"pullTimeStr"`
	//GiftCodes string `json:"giftCodes"`
}

type GiftCodes struct {
	CreateUserId  string   `json:"createUserId"`
	CreateTime    int64    `json:"createTime"`
	GiftDescribe  string   `json:"giftDescribe"`
	GiftList      []Gifts  `json:"giftList"`
	GiftCodeType  string   `json:"giftCodeType"`
	GiftPullNum   int      `json:"giftPullNum"`
	ValidityStr   string   `json:"validityStr"`
	Validity      int64    `json:"validity"` //有效期
	GiftPulledNum int      `json:"giftPulledNum"`
	GiftCode      string   `json:"giftCode"`
	RecordList    []Record `json:"recordList"`
	GiftPullUser  string   `json:"giftPullUser"` //领取人 限A类礼品码
}
