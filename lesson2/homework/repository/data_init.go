package repository

//处理数据的初始化过程，此处是没有用数据库，用文件把它一次性读到的内存

import (
	"bufio"
	"encoding/json"
	"os"
)

var (
	topicIndexMap map[int64]*Topic //暴露给外界查询的mapping
	postIndexMap  map[int64][]*Post
	LastPostId    int64 //存下最大的一个PostId，方便Post id的生成
)

func Init(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := initPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	//一行一行的进行扫描
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		if LastPostId < post.Id { //更新lastPostId
			LastPostId = post.Id
		}
		posts, ok := postTmpMap[post.ParentId]
		if !ok {
			postTmpMap[post.ParentId] = []*Post{&post}
			continue
		}
		posts = append(posts, &post)
		postTmpMap[post.ParentId] = posts //根据parentId进行索引存储
	}
	postIndexMap = postTmpMap
	return nil
}
