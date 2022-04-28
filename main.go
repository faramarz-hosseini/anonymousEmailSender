package main

import (
	"github.com/faramarz-hosseini/anonymousEmailSender/api"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	api.SetAPIHandlers(server)
	service.Run()
}
