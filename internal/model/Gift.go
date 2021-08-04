package model

type Gifts struct {
	ID  int `json:"id"`  //礼品id
	Num int `json:"num"` //数量
}

type Record struct {
	Userid      string `json:"userid"`      //用户id
	PullTime    int64  `json:"pullTime"`    //用户领取时间戳
	PullTimeStr string `json:"pullTimeStr"` //用户领取时间字符串
	//GiftCodes string `json:"giftCodes"`
}

type GiftCodes struct {
	CreateUserId  string   `json:"createUserId"`  //创建人id
	CreateTime    int64    `json:"createTime"`    //创建时间戳
	GiftDescribe  string   `json:"giftDescribe"`  //礼品码描述
	GiftList      []Gifts  `json:"giftList"`      //礼品内容
	GiftCodeType  string   `json:"giftCodeType"`  //礼品码类型
	GiftPullNum   int      `json:"giftPullNum"`   //礼品码可领取次数
	ValidityStr   string   `json:"validityStr"`   //有效时间字符串
	Validity      int64    `json:"validity"`      //有效期时间戳
	GiftPulledNum int      `json:"giftPulledNum"` //已领取次数
	GiftCode      string   `json:"giftCode"`      //礼品码
	RecordList    []Record `json:"recordList"`    //领取记录
	GiftPullUser  string   `json:"giftPullUser"`  //领取人 限A类礼品码
}
