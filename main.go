package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func incr(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	counter++
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for index :=0; index < 1000; index++ {
		wg.Add(1)
		go incr(&wg, &mutex)
	}

	wg.Wait()
  fmt.Printf("Final value %d", counter)
}
