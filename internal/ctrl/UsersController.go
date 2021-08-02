package ctrl

import (
	"ThirdProject/internal/handler"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
}

func (this *UsersController) RegisterUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := context.GetPostForm("id")
		userHandler := handler.UsersHandler{}
		user, err := userHandler.RegisterUser(id)
		if err != nil {
			context.JSON(200, map[string]interface{}{
				"code":    201,
				"message": "ERROR",
				"data":    err.Error(),
			})
		} else {
			context.JSON(200, map[string]interface{}{
				"code":    200,
				"message": "OK",
				"data":    user,
			})
		}
	}

}
