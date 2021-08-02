package ctrl

import (
	"ThirdProject/internal/handler"
	"ThirdProject/internal/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type GiftCodeController struct {
}

func (this *GiftCodeController) CreateGiftCodes() gin.HandlerFunc {
	return func(context *gin.Context) {
		giftCodes := &model.GiftCodes{}
		jsonStr, _ := context.GetPostForm("jsonStr")
		err := json.Unmarshal([]byte(jsonStr), giftCodes)

		giftHandler := handler.GiftCodeshandler{}
		if err != nil {
			context.JSON(200, map[string]interface{}{
				"code":    201,
				"message": "ERROR",
				"data":    err.Error(),
			})
		} else {

			_, e := giftHandler.CreateGiftCodes(giftCodes)
			if e != nil {
				context.JSON(200, map[string]interface{}{
					"code":    201,
					"message": "ERROR",
					"data":    e.Error(),
				})
			} else {
				context.JSON(200, map[string]interface{}{
					"code":    200,
					"message": "OK",
					"data":    "创建成功，礼品码是：" + giftCodes.GiftCode,
				})
			}
		}
	}
}

func (this *GiftCodeController) GetCiftCodes() gin.HandlerFunc {
	return func(context *gin.Context) {

		giftCode, _ := context.GetPostForm("giftCode")
		giftHandler := handler.GiftCodeshandler{}
		giftCodes, err := giftHandler.GetCiftCodes(giftCode)
		if err != nil {
			context.JSON(200, map[string]interface{}{
				"code":    201,
				"message": "ERROR",
				"data":    err.Error(),
			})
		} else {
			context.JSON(200, map[string]interface{}{
				"code":    200,
				"message": "OK",
				"data":    giftCodes,
			})
		}
	}
}

func (this *GiftCodeController) ActivateCode() gin.HandlerFunc {
	return func(context *gin.Context) {
		giftCode, _ := context.GetPostForm("giftCode")
		userId, _ := context.GetPostForm("userId")
		giftHandler := handler.GiftCodeshandler{}
		giftList, _ := giftHandler.ActivateCodeNew(giftCode, userId)
		jsonStr, _ := json.Marshal(giftList)
		context.Writer.Write(jsonStr)
		/*if err != nil {
			context.Writer.Write([]byte(err.Error()))
		} else {
			context.Writer.Write(jsonStr)
		}*/
	}
}
