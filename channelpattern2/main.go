package main

import (
	"fmt"
	"time"
)

func main() {
	sender1 := startSender("Messsage 1: ")
	sender2 := startSender("Messsage 2: ")

	for i := 1; i <= 10; i++ {
		select {
		case msgSender1 := <-sender1:
			fmt.Println(msgSender1)

		case msgSender2 := <-sender2:
			fmt.Println(msgSender2)
		}

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
