package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("num: %d", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("main quit")
}
