package main

import (
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"` 
}

func main() {

	client := gin.Default()


	client.GET("/ping", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	client.GET("/pong", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	client.POST("/message", func (c *gin.Context) {
		var msg Message

		if err := c.BindJSON(&msg); err != nil {
			c.JSON(400, gin.H{"erro": err.Error()})
		}	

		c.JSON(200, gin.H{
			"message": msg,
		})

	})

	client.Run()
}
