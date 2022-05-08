package controller

import (
	"github.com/ACking-you/TraningCamp/lesson2/homework/service"
	"strconv"
)

// PageData 最终发送给客户端的json数据对应的结构体，我们需要错误码，以及对应错误码对应的消息，最后再是数据(用空接口实现泛型
type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// QueryPageINfo 真正和客户端进行交互的函数，需要注意客户端发来的流量都是字符串形式
func QueryPageINfo(topicIdStr string) *PageData {
	pageId, err := strconv.Atoi(topicIdStr)
	if err != nil {
		return &PageData{Code: 1, Msg: err.Error(), Data: nil}
	}
	pageInfo, err := service.QueryPageInfo(int64(pageId))
	if err != nil {
		return &PageData{Code: 2, Msg: err.Error(), Data: nil}
	}
	return &PageData{Code: 0, Msg: "success", Data: pageInfo}
}
