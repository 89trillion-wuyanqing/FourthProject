package router

import (
	"ThirdProject/internal/ctrl"
	"github.com/gin-gonic/gin"
)

type GiftCodesRouter struct {
}

func (this *GiftCodesRouter) CreateGiftCodes(engine *gin.Engine) {
	crtl := ctrl.GiftCodeController{}
	engine.POST("/createGiftCodes", crtl.CreateGiftCodes())
	engine.POST("/getCiftCodes", crtl.GetCiftCodes())
	engine.POST("/activateCode", crtl.ActivateCode())
	userctrl := ctrl.UsersController{}
	engine.POST("/registerUser", userctrl.RegisterUser())

}
