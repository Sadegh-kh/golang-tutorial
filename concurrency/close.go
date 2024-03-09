package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	doneCh := make(chan int)

	go workerNum(doneCh)

	for value := range doneCh {
		fmt.Println("value: ", value)
	}
	randNum := <-doneCh
	fmt.Println("mynum", randNum)

	//when chanel closed can't writ num on that
	doneCh <- 3

}

func workerNum(doneCh chan<- int) {
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)
		if n > 50 {
			doneCh <- n
			time.Sleep(1 * time.Second)
		} else {
			close(doneCh)
			return
		}
	}

}
