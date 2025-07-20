package main

import "fmt"

func Filter[T any](inputCh <-chan T, filterFunc func(T) bool) <-chan T {
	outputCh := make(chan T)

	go func() {
		defer close(outputCh)
		for v := range inputCh {
			if filterFunc(v) {
				outputCh <- v
			}
		}
	}()

	return outputCh
}

func main() {
	// Example usage
	inputCh := make(chan int)
	filterFunc := func(v int) bool {
		return v%2 == 0 // Filter even numbers
	}

	go func() {
		defer close(inputCh)
		for i := 0; i < 10; i++ {
			inputCh <- i
		}

	}()

	for v := range Filter(inputCh, filterFunc) {
		fmt.Println(v) // Should print even numbers: 0, 2, 4, 6, 8
	}
}
