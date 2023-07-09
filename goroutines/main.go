package main

import "fmt"

func main() {
	ch_fib := make(chan int)
	ch_multi := make(chan int)
	quit_a := make(chan bool)
	quit_b := make(chan bool)
	n := 20

	go printOut(n, ch_fib, quit_a)
	go printOut(n, ch_multi, quit_b)

	fibonacci(ch_fib, quit_a)
	multples(ch_multi, quit_b, 3)

}

func printOut(n int, ch chan int, quit chan bool) {
	for i := 0; i < n; i++ {
		fmt.Println(<-(ch))
	}
	(quit) <- false
}
