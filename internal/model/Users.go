package model

type Users struct {
	UID      string `bson:"UID"`
	Id       string `bson:"id"`       //用户唯一识别码
	GoldNum  int    `bson:"goldNum"`  //金币数
	JewelNum int    `bson:"jewelNum"` //钻石数

}
