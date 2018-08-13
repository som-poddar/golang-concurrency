package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	c := make(chan string, 11)

	for i := 35; i <= 45; i++ {
		// insert arg into chan
		go printFib(c, i)
	}

	// receive from chan
	for rcvd, ok = range c {
		// fibPrint, ok := <-c
		// if ok {
		// fmt.Printf("%s, ok: %v", fibPrint, ok)
		fmt.Printf("%s,", rcvd)
		// }
	}

	close(c)
	fmt.Printf("duration: %v.", time.Since(startTime))
}

func printFib(c chan string, i int) {
	c <- fmt.Sprintf("fib(%d) : %d.\n", i, fibonacci(i))
}

func fibonacci(x int) int {
	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	return fibonacci(x-1) + fibonacci(x-2)
}
