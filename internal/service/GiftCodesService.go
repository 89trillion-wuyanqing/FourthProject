package service

import (
	"ThirdProject/internal/model"
	"time"
)

type GiftCodesService struct {
}

//验证创建礼品码信息
func (this *GiftCodesService) ValPullNum(giftCodes *model.GiftCodes) model.Result {
	loc, _ := time.LoadLocation("Local")
	strToTime, timeerr := time.ParseInLocation("2006-01-02 15:04:05", giftCodes.ValidityStr, loc)
	if timeerr != nil {
		return model.Result{Code: "203", Msg: "有效期时间格式不正确"}
	}
	timeUnixx := strToTime.Unix()
	giftCodes.Validity = timeUnixx
	if giftCodes.CreateUserId == "" {
		return model.Result{Code: "204", Msg: "创建人不能为空"}
	}

	if len(giftCodes.GiftList) <= 0 {
		return model.Result{Code: "205", Msg: "礼品码中礼品内容不能为空"}
	}
	for _, v := range giftCodes.GiftList {
		if v.ID == 0 {
			return model.Result{Code: "206", Msg: "礼品内容中礼品id不能为空"}
		}
		if v.Num <= 0 {
			return model.Result{Code: "207", Msg: "礼品内容中礼品数量不能小于0"}
		}
	}

	if giftCodes.Validity <= giftCodes.CreateTime {
		return model.Result{Code: "208", Msg: "有效时间不能小于创建时间！"}
	}
	if giftCodes.GiftCodeType == "" {
		return model.Result{Code: "209", Msg: "请选择礼品码类别！"}
	}
	if giftCodes.GiftCodeType == "A" {
		if giftCodes.GiftPullUser == "" {
			return model.Result{Code: "210", Msg: "领取人不能为空！"}
		}
		giftCodes.GiftPullNum = 1 //A类码限制一个人领一次

	} else if giftCodes.GiftCodeType == "B" {
		if giftCodes.GiftPullNum == 0 {
			return model.Result{Code: "211", Msg: "可领取次数不能为空！"}
		}
	}
	giftCodes.GiftPulledNum = 0
	return model.Result{}
}
