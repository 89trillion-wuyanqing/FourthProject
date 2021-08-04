package model

type Result struct {
	Code string      //状态码
	Msg  string      //返回消息
	Data interface{} //返回结果
}
