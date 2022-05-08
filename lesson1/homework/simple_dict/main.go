package main

import (
	"bufio"
	"fmt"
	"github.com/ACking-you/TraningCamp/lesson1/homework/simple_dict/src"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//if len(os.Args) != 2 {
	//	fmt.Fprint(os.Stderr, "usage: simpleDict [WORD] example: simpleDict hello")
	//	os.Exit(1)
	//}
	input := bufio.NewReader(os.Stdin)
	fmt.Println("请输入需要翻译的单词，或句子")
	word, err := input.ReadString('\n')

	if err != nil {
		panic(err)
	}
	word = strings.TrimRight(word, "\n")
	src.QueryDeepl(word)

	wg.Add(3)

	go func() {
		defer wg.Done()
		src.QueryYoudao(word)
	}()

	go func() {
		defer wg.Done()
		src.QueryCaiyun(word)
	}()
	go func() {
		defer wg.Done()
		src.QueryDeepl(word)
	}()

	wg.Wait()
}
