package main

// This example demonstrates a simple generator pattern using channels in Go.
func writer() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i + 1
		}
		close(ch)
	}()
	return ch
}

func main() {
	ch := writer()
	for value := range ch {
		println(value) // Output: 1, 2, ..., 10
	}
}
