package ctrl

import (
	"ThirdProject/internal/handler"
	"ThirdProject/internal/model"
	"github.com/gin-gonic/gin"
)

/**
centroller层
*/
type UsersController struct {
}

func (this *UsersController) RegisterUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := context.GetPostForm("id")

		if id == "" || id == " " {
			context.JSON(200, model.Result{Code: "201", Msg: "用户唯一标识不能为空"})
		}
		userHandler := handler.UsersHandler{}
		result := userHandler.RegisterUser(id)
		context.JSON(200, result)
	}

}
