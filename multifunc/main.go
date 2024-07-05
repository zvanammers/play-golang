package main

import (
	"fmt"
	"sync"
	"time"
)

func generateNumbers(total int, ch chan<- int) {
	for idx := 0; idx < total; idx++ {
		fmt.Printf("Sending %d to channel\n", idx)
		ch <- idx
	}
}

func printNumbers(idx int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("%d: reading %d from channel\n", idx, num)
	}
}

func main() {
	var wg sync.WaitGroup
	numberChan := make(chan int)

	for idx := 0; idx < 3; idx++ {
		wg.Add(1)
		go printNumbers(idx, numberChan, &wg)
	}
	go generateNumbers(10, numberChan)

	time.Sleep(2 * time.Second)
	close(numberChan)
	fmt.Println("Waiting for go routines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
