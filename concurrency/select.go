package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	doneCh := make(chan bool)
	errCh := make(chan error)

	go worker(doneCh, errCh)

	for {
		select {
		case <-doneCh:
			fmt.Println("done")
		case e := <-errCh:
			fmt.Println("err:", e)
		}
		time.Sleep(2 * time.Second)

	}

}

func worker(done chan<- bool, err chan<- error) {
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10)
		if n < 5 {
			done <- true
			time.Sleep(3 * time.Second)
		} else {
			err <- fmt.Errorf("this less than 5")
			time.Sleep(2 * time.Second)
		}
	}

}
