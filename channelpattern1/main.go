package main

import (
	"fmt"
	"time"
)

func main() {
	sender := startSender("Ti")

	for i := 1; i <= 5; i++ {
		fmt.Println(<-sender)
	}
}

func startSender(name string) <-chan string {
	c := make(chan string)

	for i := 1; i <= 5; i++ {
		go func() {
			c <- (name + " hello")
			time.Sleep(time.Second)
		}()
	}
	return c
}
