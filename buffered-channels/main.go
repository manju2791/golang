package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 3)
	threshold := 10
	go worker(ch, 1)
	go worker(ch, 2)
	go worker(ch, 3)

	go generateNumbers(ch, threshold)

	time.Sleep(10 * time.Second)

}

func worker(ch chan int, id int) {
	for i := range ch {
		fmt.Printf("Worker %d received data: %d\n", id, i)
		// Added sleep to mimic some processing being done by workers
		time.Sleep(1 * time.Second)
	}
}

func generateNumbers(ch chan int, threshold int) {
	for i := 1; i <= threshold; i++ {
		ch <- i
		fmt.Printf("Sent data: %d --->\n", i)
	}
	close(ch)
}
