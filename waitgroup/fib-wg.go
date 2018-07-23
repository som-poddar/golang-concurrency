package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup

	for i := 35; i <= 45; i++ {
		wg.Add(1)
		go printFib(&wg, i)
	}

	wg.Wait()
	fmt.Printf("duration: %v .\n", time.Since(startTime))
}

func printFib(wait *sync.WaitGroup, i int) {
	defer wait.Done()
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
