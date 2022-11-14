package main

import (
	"fmt"
	"time"
)

func main() {
	results := queryFirst("server1", "server2", "server3")
	fmt.Println(<-results)
}

func query(url string) string {
	time.Sleep(time.Second)
	return url
}

func queryFirst(servers ...string) <-chan string {
	c := make(chan string)

	for _, serv := range servers {
		go func(s string) {
			c <- query(s)
		}(serv)
	}
	return c

}

// Pattern nay chi lay ra 1 gia tri dau tien, vi vay 2 rountines con lai se bi leak, nen can xu ly them!!
