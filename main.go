package main

import (
	"encoding/json"
	"net/http"

	"kafka_ex_01/consumer"
	"kafka_ex_01/producer"
	"kafka_ex_01/utils"

	"github.com/gin-gonic/gin"
)

// test kafka
// service to create messages to a topic
// a consumer in the same place to proccess the message
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// to create a comment on the kafka queue
	router.POST("/comment", CreateComment)
	router.Run(":8080")

	// consumer
	consumer.InitConsumer()
}

func CreateComment(c *gin.Context) {
	// Instantiate new Message struct
	cmt := new(utils.Comment)
	if err := c.BindJSON(cmt); err != nil {
		c.IndentedJSON(400, &utils.ResponseAPI{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	// convert body into bytes and send it to kafka
	cmtInBytes, err := json.Marshal(cmt)
	if err != nil {
		c.IndentedJSON(500, &utils.ResponseAPI{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	err = producer.PushCommentToQueue("comments", cmtInBytes)
	if err != nil {
		c.IndentedJSON(500, &utils.ResponseAPI{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	// Return Comment in JSON format
	c.IndentedJSON(http.StatusOK, &utils.ResponseAPI{
		Success: true,
		Message: "Comment pushed successfully",
		Com:     *cmt,
	})
}
