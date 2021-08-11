package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64

	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				count++
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
	fmt.Println("ops:", ops)

}
