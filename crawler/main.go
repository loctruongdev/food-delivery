package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	numberOfRequests := 100
	maxWorkerNumber := 5
	queueChan := make(chan int, numberOfRequests)
	doneChane := make(chan int, maxWorkerNumber)

	for i := 1; i <= maxWorkerNumber; i++ {
		go func(name string) {
			for url := range queueChan {
				go crawler(name, url)
			}

			fmt.Printf("%s is done\n", name)

			doneChane <- 1
		}(strconv.Itoa(i))

	}

	for i := 1; i <= numberOfRequests; i++ {
		queueChan <- i
	}

	close(queueChan) // MUST HAVE!!!
	// time.Sleep(time.Second * 5)

	for i := 1; i <= maxWorkerNumber; i++ {
		<-doneChane
	}
}

func crawler(name string, url int) {
	fmt.Printf("Worker %s is crawling %d\n", name, url)
	time.Sleep(time.Second / 100)
}
