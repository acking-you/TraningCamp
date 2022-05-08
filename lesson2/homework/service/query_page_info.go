package service

import (
	"errors"
	"github.com/ACking-you/TraningCamp/lesson2/homework/repository"
	"sync"
)

// PageInfo 一个页面的信息包括，topic和它上面的post言论
type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

func QueryPageInfo(id int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(id).Do()
}

// QueryPageInfoFlow 为了防止Service构造层的高耦合度的构造PageInfo，可以用像如下通过流式处理
type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic *repository.Topic
	posts []*repository.Post
}

func NewQueryPageInfoFlow(id int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{topicId: id}
}

// Do 逐个流式初始化
func (q *QueryPageInfoFlow) Do() (*PageInfo, error) {
	//对id进行合法性验证
	if err := q.checkNum(q.topicId); err != nil {
		return nil, err
	}
	//准备好生成PageInfo的数据
	if err := q.prepareInfo(); err != nil {
		return nil, err
	}
	//打包最终的PageInfo
	if err := q.packPageInfo(); err != nil {
		return nil, err
	}
	return q.pageInfo, nil
}

func (q *QueryPageInfoFlow) checkNum(id int64) error {
	if id <= 0 {
		return errors.New("topic must larger than 0")
	}
	return nil
}

//这两个过程，由于是毫无关联的，可以用go协程进行并发处理
func (q *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	//获取Topic
	go func() {
		defer wg.Done()
		q.topic = repository.NewTopicDao().QueryTopicFromId(q.topicId)
	}()
	//获取Posts
	go func() {
		defer wg.Done()
		q.posts = repository.NewPostDao().QueryPostsFromParentId(q.topicId)
	}()

	wg.Wait()
	return nil
}

//更新最终的PageInfo
func (q *QueryPageInfoFlow) packPageInfo() error {
	q.pageInfo = &PageInfo{
		Topic:    q.topic,
		PostList: q.posts,
	}
	return nil
}
