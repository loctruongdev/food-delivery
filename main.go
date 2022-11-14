package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int = 0
	lock := new(sync.RWMutex)

	for i := 1; i <= 5; i++ {
		go func() {
			for i := 1; i <= 10000; i++ {
				lock.Lock()
				count += 1
				lock.Unlock()
			}
		}()
	}

	time.Sleep(time.Second * 7)
	fmt.Println("Count", count)

}
