package benchmark

import (
	"github.com/bytedance/gopkg/lang/fastrand"
	"math/rand"
)

var ServerIndex [10]int

// InitServerIndex 初始化服务器的描述符
func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}

// RandSelect 随机选择一个服务器
func RandSelect() int {
	return ServerIndex[rand.Intn(10)]
}

// FastRandSelect 用外部的fast包
func FastRandSelect() int {
	return ServerIndex[fastrand.Intn(10)]
}
