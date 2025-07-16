package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		for i := 0; i < 100; i++ {
			ch1 <- i
			ch2 <- i + 2
			ch3 <- i + 3
		}
	}()

	for v := range mergeChannels(ch1, ch2, ch3) {
		// Process the values from the fan-in channel
		// Here you can do whatever you need with the value v
		// For example, print it
		fmt.Println(v)
	}
}

func mergeChannels[T any](channels ...<-chan T) <-chan T {
	outCh := make(chan T)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		ch := ch
		go func() {
			defer wg.Done()
			for value := range ch {
				outCh <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()
	return outCh
}

// mergeChannels takes data from multiple input channels and put them into single output channel
func mergeChannelsWithCtx[T any](ctx context.Context, channels ...<-chan T) chan T {
	out := make(chan T)

	go func() {
		wg := &sync.WaitGroup{}
		for _, ch := range channels {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case v, ok := <-ch:
						if !ok {
							return
						}
						select {
						case out <- v:
						case <-ctx.Done():
							return
						}
					case <-ctx.Done():
						return
					}
				}
				for v := range ch {
					out <- v
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}
