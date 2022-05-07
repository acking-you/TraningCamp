package main

import (
	"TrainingCamp/lesson1/homework/simple_dict/src"
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "usage: simpleDict [WORD] example: simpleDict hello")
		os.Exit(1)
	}
	word := os.Args[1]
	wg.Add(3)
	go func() {
		defer wg.Done()
		src.QueryYoudao(word)
	}()
	go func(){
		defer wg.Done()
		src.QueryCaiyun(word)
	}()
	go func(){
		defer wg.Done()
		src.QueryDeepl(word)
	}()
	wg.Wait()
}
