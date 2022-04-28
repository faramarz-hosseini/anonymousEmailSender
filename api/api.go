package api

import (
	"github.com/faramarz-hosseini/anonymousEmailSender/rabbitmq"
	"log"
	"net/http"

	"github.com/faramarz-hosseini/anonymousEmailSender/config"
	"github.com/gin-gonic/gin"
)

func init() {
	cfg, err := config.LoadConfig("")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	rabbit := rabbitmq.GetRabbitMQ(cfg.RabbitHost)
}

func SetAPIHandlers(r *gin.Engine) {
	r.GET("/", helloWorld)
	r.POST("/send-email", sendEmailRequest)
}

func helloWorld(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		map[string]string{"hello": "world"},
	)
}

func sendEmailRequest(c *gin.Context) {

}
