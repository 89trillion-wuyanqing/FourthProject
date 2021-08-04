package http

import (
	"ThirdProject/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

func ServerInit() {
	engine := gin.Default()
	router := router.GiftCodesRouter{}
	router.CreateGiftCodes(engine)
	err := engine.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
