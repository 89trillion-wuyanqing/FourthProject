package ctrl

import (
	"ThirdProject/internal/handler"
	"ThirdProject/internal/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

/**
controller层
*/
type GiftCodeController struct {
}

func (this *GiftCodeController) CreateGiftCodes() gin.HandlerFunc {
	return func(context *gin.Context) {
		giftCodes := &model.GiftCodes{}
		jsonStr, _ := context.GetPostForm("jsonStr")
		if jsonStr == "" {
			context.JSON(200, model.Result{Code: "201", Msg: "请输入创建礼品码信息", Data: nil})
			return
		}
		err := json.Unmarshal([]byte(jsonStr), giftCodes)
		if err != nil {
			context.JSON(200, model.Result{Code: "202", Msg: "后台反序列化出错", Data: nil})
			return
		}
		giftHandler := handler.GiftCodeshandler{}
		result := giftHandler.CreateGiftCodes(giftCodes)
		context.JSON(200, result)

	}
}

func (this *GiftCodeController) GetCiftCodes() gin.HandlerFunc {
	return func(context *gin.Context) {

		giftCode, _ := context.GetPostForm("giftCode")
		if giftCode == "" {
			context.JSON(200, model.Result{Code: "201", Msg: "请输入参数", Data: nil})
			return
		}
		giftHandler := handler.GiftCodeshandler{}
		result := giftHandler.GetCiftCodes(giftCode)
		context.JSON(200, result)

	}
}

func (this *GiftCodeController) ActivateCode() gin.HandlerFunc {
	return func(context *gin.Context) {
		giftCode, _ := context.GetPostForm("giftCode")

		if giftCode == "" {
			context.JSON(200, model.Result{Code: "221", Msg: "请输入礼品码参数", Data: nil})
			return
		}
		userId, _ := context.GetPostForm("userId")
		if giftCode == "" {
			context.JSON(200, model.Result{Code: "222", Msg: "请输入用户id参数", Data: nil})
			return
		}
		giftHandler := handler.GiftCodeshandler{}
		giftList := giftHandler.ActivateCodeNew(giftCode, userId)

		context.Writer.Write(giftList)

	}
}
