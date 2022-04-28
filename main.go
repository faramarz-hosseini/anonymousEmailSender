package main

import (
	"github.com/gin-gonic/gin"

	"github.com/faramarz-hosseini/anonymousEmailSender/api"
)

func main() {
	r := gin.Default()

	api.SetAPIHandlers(r)
	r.Run()
}
