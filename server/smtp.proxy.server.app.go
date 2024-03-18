package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	client "github.com/iluyuns/smtp.proxy/client"
)

var token = "Bearer " + os.Getenv("SMTP_PROXY_TOKEN")

func main() {
	r := gin.New()
	port, err := strconv.Atoi(os.Getenv("SMTP_PROXY_PORT"))
	if err != nil {
		panic(err)
	}
	Client := &client.Client{
		Identity: "",
		Username: os.Getenv("SMTP_PROXY_USERNAME"),
		Password: os.Getenv("SMTP_PROXY_PASSWORD"),
		Host:     os.Getenv("SMTP_PROXY_HOST"),
		Port:     port,
	}

	type Mail struct {
		From    string   `json:"from" binding:"required"`
		To      []string `json:"to"`
		Subject string   `json:"subject" binding:"required"`
		Body    string   `json:"body" binding:"required"`
	}
	r.POST("/api/mail/send", func(c *gin.Context) {
		if token != c.GetHeader("Authorization") {
			c.JSON(401, gin.H{"error": "unauthorized"})
			return
		}
		var mail Mail
		// 检查请求参数是否合法
		if err := c.ShouldBindJSON(&mail); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := Client.SendMailWithTLS(mail.From, mail.To, mail.Subject, mail.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "send mail success"})
	})
	r.Run(":80")
}
