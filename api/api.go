package api

import (
	"fmt"
	"github.com/faramarz-hosseini/anonymousEmailSender/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	cfg *config.Config
	err error
)

func init() {
	cfg, err = config.LoadConfig("")
	if err != nil {
		fmt.Errorf("could not load config: %v", err)
	}
}

func SetAPIHandlers(r *gin.Engine) {
	r.GET("/", helloWorld)
	r.POST("/send-email", sendEmail)
}

func helloWorld(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		map[string]string{"hello": "world"},
	)
}

func sendEmail(c *gin.Context) {

}
