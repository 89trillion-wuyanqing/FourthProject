package http

import (
	"ThirdProject/internal/router"
	"ThirdProject/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func ServerInit() {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()
	engine := gin.Default()
	router := router.GiftCodesRouter{}
	router.CreateGiftCodes(engine)
	httpPort := utils.GetVal("server", "HttpPort")
	err := engine.Run(":" + httpPort)
	if err != nil {
		log.Println("http服务启动失败")
		panic("http服务启动失败")
	}
}
