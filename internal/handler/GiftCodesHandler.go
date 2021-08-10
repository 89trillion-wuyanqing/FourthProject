package handler

import (
	"ThirdProject/internal/model"
	"ThirdProject/internal/service"
	utils2 "ThirdProject/internal/utils"
	"context"
	"encoding/json"
	"github.com/go-redis/redis"

	"github.com/golang/protobuf/proto"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"time"
)

type GiftCodeshandler struct {
}

/**
创建礼品码业务处理
*/
func (this *GiftCodeshandler) CreateGiftCodes(giftCodes *model.GiftCodes) model.Result {
	giftCodes.GiftCode = new(utils2.RandomCode).RandomString()
	giftCodes.CreateTime = time.Now().Unix()
	giftCodes.GiftPulledNum = 0
	giftService := service.GiftCodesService{}
	result := giftService.ValPullNum(giftCodes)
	var valErr model.Result
	if result != valErr {
		return result
	}
	jsonStr, err := json.Marshal(giftCodes)
	if err != nil {
		return model.Result{Code: "212", Msg: "后台数据序列化出错", Data: nil}
	}

	r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
	if r != nil {
		return model.Result{Code: "213", Msg: "redis存储失败", Data: nil}
	}
	return model.Result{Code: "200", Msg: "成功", Data: giftCodes.GiftCode}
}

/**
获取礼品码信息业务处理
*/
func (this *GiftCodeshandler) GetCiftCodes(giftCode string) model.Result {
	result, r := utils2.StringPull(giftCode)
	if r != nil {
		if r == redis.Nil {
			return model.Result{Code: "214", Msg: "redis中不存在该礼品码", Data: nil}
		} else {
			return model.Result{Code: "215", Msg: "redis获取数据失败", Data: nil}
		}

	}
	giftCodes := &model.GiftCodes{}
	err := json.Unmarshal([]byte(result), giftCodes)
	if err != nil {
		return model.Result{Code: "202", Msg: "后台反序列化出错", Data: nil}
	}
	return model.Result{Code: "200", Msg: "成功", Data: giftCodes}
}

func (this *GiftCodeshandler) ActivateCodeNew(giftCode string, userId string) []byte {
	//先验证用户是否存在
	collection := utils2.GetCollection()
	/*currntUser := &model.Users{}

	client :=utils2.GetRedisClient()*/
	filter := bson.M{"id": userId}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		bytes, _ := proto.Marshal(&model.GeneralReward{Code: 226, Msg: "mongodb查询失败"})
		return bytes
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Println(err)

		}
	}()

	var results []model.Users
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err)
	}
	if len(results) <= 0 {
		bytes, _ := proto.Marshal(&model.GeneralReward{Code: 225, Msg: "用户没有注册登陆"})
		return bytes
	}

	//先验证验证码是否存在
	giftCodes := &model.GiftCodes{}
	result, r := utils2.StringPull(giftCode)
	if r != nil {
		if r == redis.Nil {
			bytes, _ := proto.Marshal(&model.GeneralReward{Code: 214, Msg: "redis中不存在该礼品码"})
			return bytes
		} else {
			bytes, _ := proto.Marshal(&model.GeneralReward{Code: 215, Msg: "redis获取数据失败"})
			return bytes
		}

	}

	json.Unmarshal([]byte(result), giftCodes)
	//先验证礼品码是否过期
	CurrentTime := time.Now().Unix()

	if CurrentTime > giftCodes.Validity {
		bytes, _ := proto.Marshal(&model.GeneralReward{Code: 216, Msg: "该礼品码已过期"})
		return bytes
	}
	//验证验证码是哪一类验证码
	if giftCodes.GiftCodeType == "A" {
		//查看限制人
		if giftCodes.GiftPullUser == userId {
			//查看可领取次数
			if giftCodes.GiftPullNum >= 1 {
				//领取
				byte := UpdateGift(giftCodes, userId, results[0])
				return byte
			} else {
				//已领取过
				bytes, _ := proto.Marshal(&model.GeneralReward{Code: 217, Msg: "礼品码已领取"})
				return bytes
			}

		} else {
			//不是该领取码的限制人
			bytes, _ := proto.Marshal(&model.GeneralReward{Code: 218, Msg: "你不可使用该礼品码"})
			return bytes
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
				byte := UpdateGift(giftCodes, userId, results[0])
				return byte
			} else {
				//有领取记录 使用查看是否领取过
				for i, v := range records {
					if v.Userid == userId {
						//领取列表存在该用户领取记录
						bytes, _ := proto.Marshal(&model.GeneralReward{Code: 219, Msg: "礼品码已领取"})
						return bytes
						break
					}
					if i == len(records)-1 {
						//可以领取礼品
						//增加领取记录
						//领取
						byte := UpdateGift(giftCodes, userId, results[0])
						return byte

					}
				}

			}

		} else {
			bytes, _ := proto.Marshal(&model.GeneralReward{Code: 219, Msg: "该礼品码已被领取完"})
			return bytes
		}

	} else if giftCodes.GiftCodeType == "C" {
		records := giftCodes.RecordList
		if len(records) <= 0 {
			//没有领取记录，则增加一条领取记录
			//可以领取礼品
			//增加领取记录
			//领取
			byte := UpdateGift(giftCodes, userId, results[0])
			return byte
		} else {
			//有领取记录 使用查看是否领取过
			for i, v := range records {
				if v.Userid == userId {
					//领取列表存在该用户领取记录
					bytes, _ := proto.Marshal(&model.GeneralReward{Code: 217, Msg: "礼品码已领取"})
					return bytes
					break
				}
				if i == len(records)-1 {
					//可以领取礼品
					//增加领取记录
					//领取
					byte := UpdateGift(giftCodes, userId, results[0])
					return byte
				}
			}

		}

	}
	bytes, _ := proto.Marshal(&model.GeneralReward{Code: 220, Msg: "礼品码无效"})
	return bytes
}

func UpdateGift(giftCodes *model.GiftCodes, userId string, user model.Users) []byte {
	currntUser := &model.Users{}

	client := utils2.GetRedisClient()

	//redis事务开始
	// 开启一个TxPipeline事务
	pipe := client.TxPipeline()

	// 执行事务操作，可以通过pipe读写redis
	_ = pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)
	if giftCodes.GiftCodeType == "C" {

	} else {
		giftCodes.GiftPullNum -= 1
	}

	giftCodes.GiftPulledNum += 1
	list := giftCodes.RecordList
	m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2006-01-02 15:04:05")}
	list = append(list, m1)
	giftCodes.RecordList = list

	jsonStr, err := json.Marshal(giftCodes)
	if err != nil {
		bytes, _ := proto.Marshal(&model.GeneralReward{Code: 223, Msg: "protobuf序列化出错"})
		return bytes
	}
	r := pipe.Set(giftCodes.GiftCode, string(jsonStr), 0).Err()
	//r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
	userService := service.UserService{}
	reward := model.GeneralReward{}
	reward.Code = 200
	reward.Msg = "成功"
	reward.Changes = make(map[uint32]uint64)
	reward.Balance = make(map[uint32]uint64)
	reward.Counter = make(map[uint32]uint64)
	for _, v := range giftCodes.GiftList {
		if v.ID == 1001 || v.ID == 1002 { //金币  钻石
			reward.Changes[uint32(v.ID)] = uint64(v.Num)
		}

	}
	reward.Balance[1001] = uint64(user.GoldNum)
	reward.Balance[1002] = uint64(user.JewelNum)
	reward.Counter[1001] = reward.Changes[1001] + reward.Balance[1001]
	reward.Counter[1002] = reward.Changes[1002] + reward.Balance[1002]
	currntUser.Id = user.Id
	currntUser.UID = user.UID
	currntUser.GoldNum = int(reward.Counter[1001])
	currntUser.JewelNum = int(reward.Counter[1002])

	bytes, _ := proto.Marshal(&reward)

	if r == nil {
		r2 := userService.UpdateUsers(currntUser)
		if r2 == nil {
			//提交事务
			_, _ = pipe.Exec()
			log.Println("提交事务")
			return bytes
		} else {
			pipe.Discard()
			bytes, _ := proto.Marshal(&model.GeneralReward{Code: 224, Msg: "mongodb更新失败"})
			return bytes
		}

	} else {
		pipe.Discard()
		bytes, _ := proto.Marshal(&model.GeneralReward{Code: 213, Msg: "redis存储失败"})
		return bytes
	}
}
