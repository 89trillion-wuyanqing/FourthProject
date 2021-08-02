package handler

import (
	"ThirdProject/internal/model"
	utils2 "ThirdProject/internal/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type UsersHandler struct {
}

func (this *UsersHandler) RegisterUser(id string) (model.Users, error) {

	collection := utils2.GetCollection()
	filter := bson.M{"id": id}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	//这里的结果遍历可以使用另外一种更方便的方式：
	var results []model.Users
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	if len(results) > 0 {
		//已经注册过了
		return results[0], nil
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

			return model.Users{}, errors.New("注册新用户失败")
		}

		return *insertUser, nil
	}

}
