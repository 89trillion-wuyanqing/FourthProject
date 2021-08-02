package service

import (
	"ThirdProject/internal/model"
	utils2 "ThirdProject/internal/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type UserService struct {
}

func (this *UserService) UpdateUsers(users *model.Users) {
	collection := utils2.GetCollection()
	filter := bson.M{"id": users.Id}
	update := bson.M{"$set": users}
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
