package service

import (
	"github.com/ACking-you/TraningCamp/lesson2/homework/repository"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//TestMain 是在测试进行前初始化运行的代码，且最后退出也是由他进行
func TestMain(m *testing.M) {
	repository.Init("../data/")

	os.Exit(m.Run())
}

func TestQueryPageInfo(t *testing.T) {
	pageInfo, _ := QueryPageInfo(1)
	assert.NotEqual(t, nil, pageInfo)
	assert.Equal(t, "小姐姐，快到碗里来~", pageInfo.Topic.Content)
}
