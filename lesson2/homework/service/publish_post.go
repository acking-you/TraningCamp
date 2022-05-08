package service

import (
	"errors"
	"github.com/ACking-you/TraningCamp/lesson2/homework/repository"
	"time"
	"unicode/utf8"
)

func PublishPost(topicId, userId int64, content string) (int64, error) {
	return NewPublishPostFlow(topicId, userId, content).Do()
}

func NewPublishPostFlow(topicId, userId int64, content string) *PublishPostFlow {
	return &PublishPostFlow{
		userId:  userId,
		content: content,
		topicId: topicId,
	}
}

type PublishPostFlow struct {
	userId  int64
	content string
	topicId int64

	postId int64
}

func (f *PublishPostFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.postId, nil
}

func (f *PublishPostFlow) checkParam() error {
	if f.userId <= 0 {
		return errors.New("userId id must be larger than 0")
	}
	if utf8.RuneCountInString(f.content) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f *PublishPostFlow) publish() error {
	post := &repository.Post{
		ParentId:   f.topicId,
		UserId:     f.userId,
		Content:    f.content,
		CreateTime: time.Now().Unix(),
	}
	repository.LastPostId++
	post.Id = repository.LastPostId
	if err := repository.NewPostDao().AddPost(post); err != nil {
		return err
	}
	f.postId = post.Id
	return nil
}
