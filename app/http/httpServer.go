package http

import (
	"ThirdProject/internal/router"
	"github.com/gin-gonic/gin"
)

func ServerInit() {
	engine := gin.Default()
	router := router.GiftCodesRouter{}
	router.CreateGiftCodes(engine)
	engine.Run(":8000")
}
