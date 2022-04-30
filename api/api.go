package api

import (
	"fmt"
	"github.com/faramarz-hosseini/anonymousEmailSender/rabbitmq"
	"github.com/streadway/amqp"
	"log"
	"net/http"

	"github.com/faramarz-hosseini/anonymousEmailSender/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	GinServer  *gin.Engine
	RabbitCon  *amqp.Connection
	RabbitChan *amqp.Channel
}

func InitializeServer(r *gin.Engine) *Server {
	cfg, err := config.LoadConfig("")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	rabbitCon := rabbitmq.GetRabbitMQ(cfg.RabbitHost)
	rabbitChan, err := rabbitCon.Channel()
	if err != nil {
		log.Fatalf("could not get rabbit channel: %v", err)
	}

	server := Server{
		GinServer:  r,
		RabbitCon:  rabbitCon,
		RabbitChan: rabbitChan,
	}

	r.GET("/", index)
	r.POST("/send-email", server.sendEmailRequest)

	return &server
}

func index(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		map[string]string{"hello": "world"},
	)
}

func (s *Server) sendEmailRequest(c *gin.Context) {
	if c.Query("content") == "" || c.Query("receiver") == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	request := `{"content": "%s", "receiver": "%s"}`
	request = fmt.Sprintf(request, c.Query("content"), c.Query("receiver"))
	err := s.RabbitChan.Publish(
		"",
		"email-request",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(request),
		},
	)
	if err != nil {
		log.Fatalf("could not publish to email-request queue: %v", err)
	}
}
