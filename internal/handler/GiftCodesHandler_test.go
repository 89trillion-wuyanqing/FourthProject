package handler

import (
	"ThirdProject/internal/model"
	"ThirdProject/internal/utils"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestGiftCodeshandler_CreateGiftCodes(t *testing.T) {
	giftHandler := GiftCodeshandler{}
	gifts := []model.Gifts{}
	gifts = append(gifts, model.Gifts{ID: 1001, Num: 11})
	gifts = append(gifts, model.Gifts{ID: 1002, Num: 11})
	var giftCodes = &model.GiftCodes{GiftCodeType: "B", CreateUserId: "100001", GiftList: gifts, GiftPullNum: 10, ValidityStr: "2021-08-12 02:03:45"}
	tests := []struct {
		name string
		args *model.GiftCodes
		want int
	}{
		{"test1", giftCodes, 8},
		{"test2", giftCodes, 8},
		{"test3", giftCodes, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := giftHandler.CreateGiftCodes(tt.args)
			if op, ok := got.Data.(string); ok && len(op) != tt.want {
				t.Errorf("CreateGiftCodes() = %v, want %v", got, tt.want)
			}
			fmt.Println(got.Data)
		})
	}

}

func TestGiftCodeshandler_GetCiftCodes(t *testing.T) {
	//test1---对应指定用户一次性领取
	giftCode1 := "LN2RJV7O"
	pullStr1, _ := utils.StringPull(giftCode1)
	giftCodes1 := &model.GiftCodes{}
	json.Unmarshal([]byte(pullStr1), giftCodes1)

	//test2 --- 测试-限定次数、不指定用户
	giftCode2 := "PWPIJT8N"
	pullStr2, _ := utils.StringPull(giftCode2)
	giftCodes2 := &model.GiftCodes{}
	json.Unmarshal([]byte(pullStr2), giftCodes2)

	//test1 --- 测试-不限次数、不限用户
	giftCode3 := "VA20V0PR"
	pullStr3, _ := utils.StringPull(giftCode3)
	giftCodes3 := &model.GiftCodes{}
	json.Unmarshal([]byte(pullStr3), giftCodes3)

	tests := []struct {
		name string
		args string
		want model.GiftCodes
	}{
		{"test1", giftCode1, *giftCodes1},
		{"test2", giftCode2, *giftCodes2},
		{"test3", giftCode3, *giftCodes3},
	}
	giftHandler := GiftCodeshandler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := giftHandler.GetCiftCodes(tt.args); !reflect.DeepEqual(got.Data, tt.want) {
				t.Errorf("GetCiftCodes() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestGiftCodeshandler_ActivateCodeNew(t *testing.T) {
	type args struct {
		giftCode string
		userId   string
	}

	//test1---对应指定用户一次性领取
	giftCode1 := "LN2RJV7O"
	pullStr1, _ := utils.StringPull(giftCode1)
	giftCodes1 := &model.GiftCodes{}
	json.Unmarshal([]byte(pullStr1), giftCodes1)

	//test2 --- 测试-限定次数、不指定用户
	giftCode2 := "PWPIJT8N"
	pullStr2, _ := utils.StringPull(giftCode2)
	giftCodes2 := &model.GiftCodes{}
	json.Unmarshal([]byte(pullStr2), giftCodes2)

	//test1 --- 测试-不限次数、不限用户
	giftCode3 := "VA20V0PR"
	pullStr3, _ := utils.StringPull(giftCode3)
	giftCodes3 := &model.GiftCodes{}
	json.Unmarshal([]byte(pullStr3), giftCodes3)

	tests := []struct {
		name string
		args args
		want model.GiftCodes
	}{
		{"test1", args{giftCode1, "100001"}, *giftCodes1},
		{"test2", args{giftCode2, "1001"}, *giftCodes2},
		{"test3", args{giftCode3, "1234"}, *giftCodes3},
	}
	giftHandler := GiftCodeshandler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := giftHandler.ActivateCodeNew(tt.args.giftCode, tt.args.userId); !reflect.DeepEqual(got, tt.want.GiftList) {
				t.Errorf("ActivateCodeNew() = %v, want %v", got, tt.want)
			}
		})
	}

}
