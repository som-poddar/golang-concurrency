package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	for i := 35; i <= 45; i++ {
		go printFib(i)
	}

	fmt.Printf("duration: %v .\n", time.Since(startTime))
	time.Sleep(7 * time.Second)
}

func printFib(i int) {
	fmt.Printf("fib(%d) : %d. \n", i, fibonacci(i))
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
