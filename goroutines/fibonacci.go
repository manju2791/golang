package main

import "fmt"

func fibonacci(ch chan int, quit chan bool) {
	fmt.Println("Print fibonacci: ")

	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("End fibonacci.")
			return
		}
	}
}

func multples(ch chan int, quit chan bool, num int) {
	fmt.Printf("Print multiples of %d \n", num)

	x := 2
	for {
		select {
		case ch <- x * num:
			x = x + 1
		case <-quit:
			fmt.Println("End multiples.")
			return
		}
	}
}
