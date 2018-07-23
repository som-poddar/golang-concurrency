package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	for i := 35; i <= 45; i++ {
		fmt.Printf("fib(%d) : %d. \n", i, fibonacci(i))
	}

	fmt.Printf("duration: %v .\n", time.Since(startTime))
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
