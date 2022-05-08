package main

import (
	"github.com/ACking-you/TraningCamp/lesson2/homework/controller"
	"github.com/ACking-you/TraningCamp/lesson2/homework/repository"
	"gopkg.in/gin-gonic/gin.v1"
	"os"
	"strings"
)

//最后再通过gin框架搭建服务器

func main() {
	if err := Init("./lesson2/homework/data/"); err != nil {
		os.Exit(-1)
	}

	r := gin.Default()
	r.GET("me:id", func(c *gin.Context) {
		topicId := c.Param("id")
		topicId = strings.TrimLeft(topicId, ":,")
		println(topicId)
		data := controller.QueryPageINfo(topicId)
		c.JSONP(200, data)
	})

	r.POST("/post/do", func(c *gin.Context) {
		uid, _ := c.GetPostForm("uid")
		println(uid)
		topicId, _ := c.GetPostForm("topic_id")
		println(topicId)
		content, _ := c.GetPostForm("content")
		println(content)
		data := controller.PublishPost(uid, topicId, content)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}

}

func Init(filepath string) error {
	err := repository.Init(filepath)
	if err != nil {
		return err
	}
	return nil
}
