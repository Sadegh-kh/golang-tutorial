package main

import (
	"fmt"
	"sync"
)

// if receiver c <-chan int
// sender c chan<- int
func sum1(numbers []int, c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	var sum int
	for _, value := range numbers {
		sum += value
	}
	fmt.Println("sum ", sum)
	c <- sum
}

func addToMap(myMap *sync.Map, key, value string, wg *sync.WaitGroup) {
	myMap.Store(key, value)
	wg.Done()
}

func main() {
	numbers := []int{6, 7, 4, 3, -1, 5, 10, 2}
	chanel := make(chan int, 2)

	var wg sync.WaitGroup
	wg.Add(2)
	go sum1(numbers[:len(numbers)/2], chanel, &wg)
	go sum1(numbers[len(numbers)/2:], chanel, &wg)

	wg.Wait()
	// concurrent map writes error
	//m := map[string]string{}
	m := sync.Map{}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go addToMap(&m, fmt.Sprintf("counter %v", i), "fruit", &wg)
		go addToMap(&m, fmt.Sprintf("counter 0"), "bnana", &wg)
	}
	value, _ := m.Load("counter 0")
	fmt.Println("my value of counter 1", value)
	wg.Wait()
	//sum := <-chanel
	//
	//fmt.Println("sum1:", sum)
	//
	//sum = <-chanel
	//
	//fmt.Println("sum2:", sum)

}
