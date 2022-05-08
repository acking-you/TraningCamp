package repository

import "sync"

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

// TopicDao 定义一个空的结构体，为了让请求的函数不会被重名之类的（相当于给函数加个namespace）
type TopicDao struct {
}

// 定义全局变量实现单例模式，其中sync.Once类型能够通过Do方法使得代码块只执行一次
var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDao() *TopicDao {
	topicOnce.Do(func() {
		topicDao = new(TopicDao) //初始化指针
	})
	return topicDao
}

func (_ *TopicDao) QueryTopicFromId(id int64) *Topic {
	return topicIndexMap[id]
}
