package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	mgoCli     *mongo.Client
	collection *mongo.Collection
)

func init() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

//暴露获取mongodb连接，单例模式
func GetMgoCli() *mongo.Client {

	return mgoCli
}

func GetCollection() *mongo.Collection {
	GetMgoCli()
	collection = mgoCli.Database("Users").Collection("users")
	return collection
}
