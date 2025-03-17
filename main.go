package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	timeStart := time.Now()

	// Start channle to be used and we used "1024" as buffer size:
	responseChannle := make(chan int, 1024)

	// Make wait group to be controlle all golang rounines:
	wg := &sync.WaitGroup{}

	go maskanCalculation(20, responseChannle, wg)
	wg.Add(1) // should set 1 to make sure 1 process should be done.
	go goldCalculation(10, responseChannle, wg)
	wg.Add(1)

	// Order to wait group to wait untill all wait group be done:
	wg.Wait()

	// We should close the channle if not, channle always be in running and listinning:
	close(responseChannle)

	for responseChann := range responseChannle {
		fmt.Println(responseChann)
	}

	fmt.Println("GO Lang First Concurrent Training.")

	fmt.Println(time.Since(timeStart))
}

func maskanCalculation(price int, responseChannle chan int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println(price)

	responseChannle <- <-responseChannle

	/*
	 Note:
	 reterun statgement not used becuse no main thread to speak with
	 other thread and all and all execution result may in  chann :)
	*/

	wg.Done()
}

func goldCalculation(price int, responseChannle chan int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println(price)

	responseChannle <- price

	wg.Done()
}
