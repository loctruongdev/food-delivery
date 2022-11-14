package main

import (
	"fmt"
	"time"
)

func main() {
	responseChan := make(chan string)
	var results []string

	go func() { responseChan <- fetchAPI("users") }()
	// go func(m string) { responseChan <- fetchAPI(m) }("users")
	go func() { responseChan <- fetchAPI("categories") }()
	go func() { responseChan <- fetchAPI("products") }()

	for i := 1; i <= 3; i++ {
		results = append(results, <-responseChan)
	}

	fmt.Println("Results", results)

}

func fetchAPI(model string) string {
	time.Sleep(time.Second)
	return model
}

// Pattern nay ok. Thuong dung trong "search" API, "service" API
