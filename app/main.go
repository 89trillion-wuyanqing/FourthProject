package main

import (
	"ThirdProject/app/http"
	utils2 "ThirdProject/internal/utils"
	"fmt"
	"log"
)

func main() {
	rediserror := utils2.InitClient()
	if rediserror != nil {
		fmt.Println("连接失败")
		log.Println("redis服务连接失败")
	}

	http.ServerInit()

}
