package main

import "time"
import "fmt"

func main() {
	reader(double(writer()))
}

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

func double(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range ch {
			time.Sleep(500 * time.Millisecond) // Simulate some processing delay
			out <- v * 2
		}
		close(out)
	}()

	return out
}

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
