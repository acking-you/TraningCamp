package repository

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	UserId     int64  `json:"user_id"`
}

type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDao() *PostDao {
	postOnce.Do(func() {
		postDao = new(PostDao)
	})
	return postDao
}

func (_ *PostDao) QueryPostsFromParentId(id int64) []*Post {
	return postIndexMap[id]
}

func (d *PostDao) AddPost(post *Post) error {
	//加锁保证同时请求的并发安全
	lock := sync.Mutex{}
	lock.Lock()
	posts, ok := postIndexMap[post.ParentId]
	if !ok {
		return errors.New("post invalid,not exist parent id")
	}
	//注意更新map里的数据，go切片并不像C++里的Vector，可能append后操作的就不是同一片 底层数组了

	postIndexMap[post.ParentId] = append(posts, post)
	err := fileDataInsertPost("./lesson2/homework/data/", post)
	if err != nil {
		return err
	}

	lock.Unlock()
	return nil
}

func fileDataInsertPost(filePath string, post *Post) error {
	open, err := os.OpenFile(filePath+"post", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(open)

	data, err := json.Marshal(*post)
	if err != nil {
		return err
	}
	writer.WriteString("\r\n")
	writer.Write(data)
	writer.Flush()
	return nil
}
