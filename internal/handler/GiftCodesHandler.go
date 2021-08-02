package handler

import (
	"ThirdProject/internal/model"
	"ThirdProject/internal/service"
	utils2 "ThirdProject/internal/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"time"
)

type GiftCodeshandler struct {
}

func (this *GiftCodeshandler) CreateGiftCodes(giftCodes *model.GiftCodes) (bool, error) {
	giftCodes.GiftCode = new(utils2.RandomCode).RandomString()
	giftCodes.CreateTime = time.Now().Unix()
	giftCodes.GiftPulledNum = 0
	giftService := service.GiftCodesService{}
	_, valErr := giftService.ValPullNum(giftCodes)
	if valErr != nil {
		return false, valErr
	}
	jsonStr, err := json.Marshal(giftCodes)
	if err != nil {
		return false, err
	}

	r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
	if r != nil {
		return false, r
	}
	return true, nil
}

func (this *GiftCodeshandler) GetCiftCodes(giftCode string) (*model.GiftCodes, error) {
	result, r := utils2.StringPull(giftCode)
	if r != nil {
		return nil, r
	}
	giftCodes := &model.GiftCodes{}
	json.Unmarshal([]byte(result), giftCodes)
	return giftCodes, nil
}

func (this *GiftCodeshandler) ActivateCode(giftCode string, userId string) ([]model.Gifts, error) {
	//先验证验证码是否存在
	giftCodes := &model.GiftCodes{}
	result, r := utils2.StringPull(giftCode)
	if r != nil {
		return nil, r
	}

	json.Unmarshal([]byte(result), giftCodes)
	//先验证礼品码是否过期
	CurrentTime := time.Now().Unix()

	if CurrentTime > giftCodes.Validity {
		return nil, errors.New("该礼品码已过期")
	}
	//验证验证码是哪一类验证码
	if giftCodes.GiftCodeType == "A" {
		//查看限制人
		if giftCodes.GiftPullUser == userId {
			//查看可领取次数
			if giftCodes.GiftPullNum >= 1 {
				//领取
				giftCodes.GiftPullNum -= 1
				giftCodes.GiftPulledNum += 1
				list := giftCodes.RecordList
				m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2021-03-04 15:04:05")}
				list = append(list, m1)

				giftCodes.RecordList = list

				jsonStr, err := json.Marshal(giftCodes)
				if err != nil {
					return nil, errors.New("存储数据序列化出问题")
				}
				r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
				if r != nil {
					return nil, r
				}
				return giftCodes.GiftList, nil
			} else {
				//已领取过
				return nil, errors.New("该激活码已被领取")
			}

		} else {
			//不是该领取码的限制人
			return nil, errors.New("你不可使用该激活码")
		}
	} else if giftCodes.GiftCodeType == "B" {
		//不限用户  不限次数，用户是否用过
		//先判断可领次数是否可以领取
		//领取礼品
		if giftCodes.GiftPullNum > 0 {
			records := giftCodes.RecordList
			if len(records) <= 0 {
				//没有领取记录，则增加一条领取记录
				//可以领取礼品
				//增加领取记录
				//领取
				giftCodes.GiftPullNum -= 1
				giftCodes.GiftPulledNum += 1
				list := giftCodes.RecordList
				m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2021-03-04 15:04:05")}
				list = append(list, m1)

				giftCodes.RecordList = list

				jsonStr, err := json.Marshal(giftCodes)
				if err != nil {
					return nil, errors.New("存储数据序列化出问题")
				}
				r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
				if r != nil {
					return nil, r
				}
				return giftCodes.GiftList, nil
			} else {
				//有领取记录 使用查看是否领取过
				for i, v := range records {
					if v.Userid == userId {
						//领取列表存在该用户领取记录
						return nil, errors.New("你已经领取该礼品码")
						break
					}
					if i == len(records)-1 {
						//可以领取礼品
						//增加领取记录
						//领取
						giftCodes.GiftPullNum -= 1
						giftCodes.GiftPulledNum += 1
						list := giftCodes.RecordList
						m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2021-03-04 15:04:05")}
						list = append(list, m1)
						giftCodes.RecordList = list

						jsonStr, err := json.Marshal(giftCodes)
						if err != nil {
							return nil, errors.New("存储数据序列化出问题")
						}
						r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
						if r != nil {
							return nil, r
						}
						return giftCodes.GiftList, nil

					}
				}

			}

		} else {
			return nil, errors.New("该礼品码已被领取完")
		}

	} else if giftCodes.GiftCodeType == "C" {
		records := giftCodes.RecordList
		if len(records) <= 0 {
			//没有领取记录，则增加一条领取记录
			//可以领取礼品
			//增加领取记录
			//领取

			giftCodes.GiftPulledNum += 1
			list := giftCodes.RecordList
			m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2021-03-04 15:04:05")}
			list = append(list, m1)

			giftCodes.RecordList = list
			jsonStr, err := json.Marshal(giftCodes)
			if err != nil {
				return nil, errors.New("存储数据序列化出问题")
			}
			r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
			if r != nil {
				return nil, r
			}
			return giftCodes.GiftList, nil
		} else {
			//有领取记录 使用查看是否领取过
			for i, v := range records {
				if v.Userid == userId {
					//领取列表存在该用户领取记录
					return nil, errors.New("你已经领取该礼品码")
					break
				}
				if i == len(records)-1 {
					//可以领取礼品
					//增加领取记录
					//领取

					giftCodes.GiftPulledNum += 1
					list := giftCodes.RecordList
					m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2021-03-04 15:04:05")}
					list = append(list, m1)
					giftCodes.RecordList = list
					jsonStr, err := json.Marshal(giftCodes)
					if err != nil {
						return nil, errors.New("存储数据序列化出问题")
					}
					r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
					if r != nil {
						return nil, r
					}
					return giftCodes.GiftList, nil

				}
			}

		}

	}
	return nil, errors.New("礼品码无效")
}

func (this *GiftCodeshandler) ActivateCodeNew(giftCode string, userId string) ([]byte, error) {
	//先验证用户是否存在
	currntUser := &model.Users{}
	collection := utils2.GetCollection()
	filter := bson.M{"id": userId}
	cursor, err := collection.Find(context.TODO(), filter, options.Find().SetSkip(0), options.Find().SetLimit(10))
	if err != nil {
		log.Fatal(err)
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	var results []model.Users
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	if len(results) <= 0 {
		bytes, _ := proto.Marshal(&model.GeneralReward{Code: 201, Msg: "用户没有注册登陆"})
		return bytes, errors.New("该用户没有注册登陆")
	}

	//先验证验证码是否存在
	giftCodes := &model.GiftCodes{}
	result, r := utils2.StringPull(giftCode)
	if r != nil {

		bytes, _ := proto.Marshal(&model.GeneralReward{Code: 202, Msg: "礼品码不存在"})
		return bytes, r
	}

	json.Unmarshal([]byte(result), giftCodes)
	//先验证礼品码是否过期
	CurrentTime := time.Now().Unix()

	if CurrentTime > giftCodes.Validity {
		bytes, _ := proto.Marshal(&model.GeneralReward{Code: 203, Msg: "该礼品码已过期"})
		return bytes, errors.New("该礼品码已过期")
	}
	//验证验证码是哪一类验证码
	if giftCodes.GiftCodeType == "A" {
		//查看限制人
		if giftCodes.GiftPullUser == userId {
			//查看可领取次数
			if giftCodes.GiftPullNum >= 1 {
				//领取
				giftCodes.GiftPullNum -= 1
				giftCodes.GiftPulledNum += 1
				list := giftCodes.RecordList
				m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2021-03-04 15:04:05")}
				list = append(list, m1)

				giftCodes.RecordList = list

				jsonStr, err := json.Marshal(giftCodes)
				if err != nil {
					bytes, _ := proto.Marshal(&model.GeneralReward{Code: 204, Msg: "redis存储数据有问题"})
					return bytes, errors.New("存储数据序列化出问题")
				}
				r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
				if r != nil {
					bytes, _ := proto.Marshal(&model.GeneralReward{Code: 205, Msg: "礼品码领取过程出问题"})
					return bytes, r
				}
				reward := model.GeneralReward{}
				reward.Code = 200
				reward.Msg = giftCodes.GiftDescribe
				reward.Changes = make(map[uint32]uint64)
				reward.Balance = make(map[uint32]uint64)
				reward.Counter = make(map[uint32]uint64)
				for _, v := range giftCodes.GiftList {
					if v.ID == 1001 || v.ID == 1002 { //金币  钻石
						reward.Changes[uint32(v.ID)] = uint64(v.Num)
					}

				}
				reward.Balance[1001] = uint64(results[0].GoldNum)
				reward.Balance[1002] = uint64(results[0].JewelNum)

				reward.Counter[1001] = reward.Changes[1001] + reward.Balance[1001]
				reward.Counter[1002] = reward.Changes[1002] + reward.Balance[1002]
				currntUser.Id = results[0].Id
				currntUser.UID = results[0].UID
				currntUser.GoldNum = int(reward.Counter[1001])
				currntUser.JewelNum = int(reward.Counter[1002])
				userService := service.UserService{}
				userService.UpdateUsers(currntUser)
				bytes, _ := proto.Marshal(&reward)
				return bytes, nil
			} else {
				//已领取过
				bytes, _ := proto.Marshal(&model.GeneralReward{Code: 206, Msg: "礼品码已被领取"})
				return bytes, errors.New("该礼品码已被领取")
			}

		} else {
			//不是该领取码的限制人
			bytes, _ := proto.Marshal(&model.GeneralReward{Code: 207, Msg: "不可使用该礼品码"})
			return bytes, errors.New("你不可使用该礼品码")
		}
	} else if giftCodes.GiftCodeType == "B" {
		//不限用户  不限次数，用户是否用过
		//先判断可领次数是否可以领取
		//领取礼品
		if giftCodes.GiftPullNum > 0 {
			records := giftCodes.RecordList
			if len(records) <= 0 {
				//没有领取记录，则增加一条领取记录
				//可以领取礼品
				//增加领取记录
				//领取
				giftCodes.GiftPullNum -= 1
				giftCodes.GiftPulledNum += 1
				list := giftCodes.RecordList
				m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2021-03-04 15:04:05")}
				list = append(list, m1)

				giftCodes.RecordList = list

				jsonStr, err := json.Marshal(giftCodes)
				if err != nil {
					bytes, _ := proto.Marshal(&model.GeneralReward{Code: 204, Msg: "redis存储数据有问题"})
					return bytes, errors.New("存储数据序列化出问题")
				}
				r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
				if r != nil {
					bytes, _ := proto.Marshal(&model.GeneralReward{Code: 205, Msg: "礼品码领取过程出问题"})
					return bytes, r
				}

				reward := model.GeneralReward{}
				reward.Code = 200
				reward.Msg = giftCodes.GiftDescribe
				reward.Changes = make(map[uint32]uint64)
				reward.Balance = make(map[uint32]uint64)
				reward.Counter = make(map[uint32]uint64)
				for _, v := range giftCodes.GiftList {
					if v.ID == 1001 || v.ID == 1002 { //金币  钻石
						reward.Changes[uint32(v.ID)] = uint64(v.Num)
					}

				}
				reward.Balance[1001] = uint64(results[0].GoldNum)
				reward.Balance[1002] = uint64(results[0].JewelNum)

				reward.Counter[1001] = reward.Changes[1001] + reward.Balance[1001]
				reward.Counter[1002] = reward.Changes[1002] + reward.Balance[1002]
				currntUser.Id = results[0].Id
				currntUser.UID = results[0].UID
				currntUser.GoldNum = int(reward.Counter[1001])
				currntUser.JewelNum = int(reward.Counter[1002])
				userService := service.UserService{}
				userService.UpdateUsers(currntUser)
				bytes, _ := proto.Marshal(&reward)
				return bytes, nil
			} else {
				//有领取记录 使用查看是否领取过
				for i, v := range records {
					if v.Userid == userId {
						//领取列表存在该用户领取记录
						bytes, _ := proto.Marshal(&model.GeneralReward{Code: 206, Msg: "礼品码已被领取"})
						return bytes, errors.New("你已经领取该礼品码")
						break
					}
					if i == len(records)-1 {
						//可以领取礼品
						//增加领取记录
						//领取
						giftCodes.GiftPullNum -= 1
						giftCodes.GiftPulledNum += 1
						list := giftCodes.RecordList
						m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2021-03-04 15:04:05")}
						list = append(list, m1)
						giftCodes.RecordList = list

						jsonStr, err := json.Marshal(giftCodes)
						if err != nil {
							bytes, _ := proto.Marshal(&model.GeneralReward{Code: 204, Msg: "redis存储数据有问题"})
							return bytes, errors.New("存储数据序列化出问题")
						}
						r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
						if r != nil {
							bytes, _ := proto.Marshal(&model.GeneralReward{Code: 205, Msg: "礼品码领取过程出问题"})
							return bytes, r
						}
						reward := model.GeneralReward{}
						reward.Code = 200
						reward.Msg = giftCodes.GiftDescribe
						reward.Changes = make(map[uint32]uint64)
						reward.Balance = make(map[uint32]uint64)
						reward.Counter = make(map[uint32]uint64)
						for _, v := range giftCodes.GiftList {
							if v.ID == 1001 || v.ID == 1002 { //金币  钻石
								reward.Changes[uint32(v.ID)] = uint64(v.Num)
							}

						}
						reward.Balance[1001] = uint64(results[0].GoldNum)
						reward.Balance[1002] = uint64(results[0].JewelNum)

						reward.Counter[1001] = reward.Changes[1001] + reward.Balance[1001]
						reward.Counter[1002] = reward.Changes[1002] + reward.Balance[1002]
						currntUser.Id = results[0].Id
						currntUser.UID = results[0].UID
						currntUser.GoldNum = int(reward.Counter[1001])
						currntUser.JewelNum = int(reward.Counter[1002])
						userService := service.UserService{}
						userService.UpdateUsers(currntUser)
						bytes, _ := proto.Marshal(&reward)
						return bytes, nil

					}
				}

			}

		} else {
			bytes, _ := proto.Marshal(&model.GeneralReward{Code: 208, Msg: "该礼品码已被领取完"})
			return bytes, errors.New("该礼品码已被领取完")
		}

	} else if giftCodes.GiftCodeType == "C" {
		records := giftCodes.RecordList
		if len(records) <= 0 {
			//没有领取记录，则增加一条领取记录
			//可以领取礼品
			//增加领取记录
			//领取

			giftCodes.GiftPulledNum += 1
			list := giftCodes.RecordList
			m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2021-03-04 15:04:05")}
			list = append(list, m1)

			giftCodes.RecordList = list
			jsonStr, err := json.Marshal(giftCodes)
			if err != nil {
				bytes, _ := proto.Marshal(&model.GeneralReward{Code: 204, Msg: "redis存储数据有问题"})
				return bytes, errors.New("存储数据序列化出问题")
			}
			r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
			if r != nil {
				bytes, _ := proto.Marshal(&model.GeneralReward{Code: 205, Msg: "礼品码领取过程出问题"})
				return bytes, r
			}

			reward := model.GeneralReward{}
			reward.Code = 200
			reward.Msg = giftCodes.GiftDescribe
			reward.Changes = make(map[uint32]uint64)
			reward.Balance = make(map[uint32]uint64)
			reward.Counter = make(map[uint32]uint64)
			for _, v := range giftCodes.GiftList {
				if v.ID == 1001 || v.ID == 1002 { //金币  钻石
					reward.Changes[uint32(v.ID)] = uint64(v.Num)
				}

			}
			reward.Balance[1001] = uint64(results[0].GoldNum)
			reward.Balance[1002] = uint64(results[0].JewelNum)

			reward.Counter[1001] = reward.Changes[1001] + reward.Balance[1001]
			reward.Counter[1002] = reward.Changes[1002] + reward.Balance[1002]
			currntUser.Id = results[0].Id
			currntUser.UID = results[0].UID
			currntUser.GoldNum = int(reward.Counter[1001])
			currntUser.JewelNum = int(reward.Counter[1002])
			userService := service.UserService{}
			userService.UpdateUsers(currntUser)
			bytes, _ := proto.Marshal(&reward)
			return bytes, nil
		} else {
			//有领取记录 使用查看是否领取过
			for i, v := range records {
				if v.Userid == userId {
					//领取列表存在该用户领取记录
					bytes, _ := proto.Marshal(&model.GeneralReward{Code: 206, Msg: "礼品码已被领取"})
					return bytes, errors.New("你已经领取该礼品码")
					break
				}
				if i == len(records)-1 {
					//可以领取礼品
					//增加领取记录
					//领取

					giftCodes.GiftPulledNum += 1
					list := giftCodes.RecordList
					m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2021-03-04 15:04:05")}
					list = append(list, m1)
					giftCodes.RecordList = list
					jsonStr, err := json.Marshal(giftCodes)
					if err != nil {
						bytes, _ := proto.Marshal(&model.GeneralReward{Code: 204, Msg: "redis存储数据有问题"})
						return bytes, errors.New("存储数据序列化出问题")
					}
					r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
					if r != nil {
						bytes, _ := proto.Marshal(&model.GeneralReward{Code: 205, Msg: "礼品码领取过程出问题"})
						return bytes, r
					}
					reward := model.GeneralReward{}
					reward.Code = 200
					reward.Msg = giftCodes.GiftDescribe
					reward.Changes = make(map[uint32]uint64)
					reward.Balance = make(map[uint32]uint64)
					reward.Counter = make(map[uint32]uint64)
					for _, v := range giftCodes.GiftList {
						if v.ID == 1001 || v.ID == 1002 { //金币  钻石
							reward.Changes[uint32(v.ID)] = uint64(v.Num)
						}

					}
					reward.Balance[1001] = uint64(results[0].GoldNum)
					reward.Balance[1002] = uint64(results[0].JewelNum)

					reward.Counter[1001] = reward.Changes[1001] + reward.Balance[1001]
					reward.Counter[1002] = reward.Changes[1002] + reward.Balance[1002]
					currntUser.Id = results[0].Id
					currntUser.UID = results[0].UID
					currntUser.GoldNum = int(reward.Counter[1001])
					currntUser.JewelNum = int(reward.Counter[1002])
					userService := service.UserService{}
					userService.UpdateUsers(currntUser)
					bytes, _ := proto.Marshal(&reward)
					return bytes, nil
				}
			}

		}

	}
	bytes, _ := proto.Marshal(&model.GeneralReward{Code: 209, Msg: "礼品码无效"})
	return bytes, errors.New("礼品码无效")
}
