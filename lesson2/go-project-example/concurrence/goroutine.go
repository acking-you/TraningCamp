package concurrence

import (
	"fmt"
	"sync"
)

func hello() {
	fmt.Println("hello")
}
func ManyGo() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			hello()
		}()
	}
	wg.Wait()
}
