package handler

import (
	"ThirdProject/internal/model"
	utils2 "ThirdProject/internal/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type UsersHandler struct {
}

/**
用户注册
*/
func (this *UsersHandler) RegisterUser(id string) model.Result {

	collection := utils2.GetCollection()
	filter := bson.M{"id": id}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return model.Result{Code: "202", Msg: "mongodb查询失败", Data: nil}
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Println(err)
		}
	}()
	//这里的结果遍历可以使用另外一种更方便的方式：
	var results []model.Users
	err = cursor.All(context.TODO(), &results)
	if err != nil {

		return model.Result{Code: "203", Msg: "mongodb游标获取数据失败", Data: nil}
	}
	if len(results) > 0 {
		//已经注册过了
		return model.Result{Code: "204", Msg: "该用户已注册", Data: nil}
	} else {
		//注册新用户
		//插入某一条数据
		//var iResult *mongo.InsertOneResult
		insertUser := &model.Users{}
		random := utils2.RandomCode{}
		insertUser.UID = random.RandomObjectId()
		insertUser.Id = id
		insertUser.JewelNum = 0
		insertUser.GoldNum = 0

		_, e := collection.InsertOne(context.TODO(), insertUser)
		if e != nil {

			return model.Result{Code: "205", Msg: "注册新用户失败", Data: nil}
		}

		return model.Result{Code: "200", Msg: "成功", Data: insertUser}
	}

}
