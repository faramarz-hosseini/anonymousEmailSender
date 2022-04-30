package main

import (
	"github.com/faramarz-hosseini/anonymousEmailSender/api"
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

	server := api.InitializeServer(ginServer)
	defer server.RabbitCon.Close()
	defer server.RabbitChan.Close()

	server.GinServer.Run()
}
