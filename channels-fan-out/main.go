package main

import "sync"

func main() {
	channel := make(chan int)
	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	channels := SplitChannel(channel, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for val := range channels[0] {
			println("Channel 1 received:", val)
		}
	}()
	go func() {
		defer wg.Done()
		for val := range channels[1] {
			println("Channel 2 received:", val)
		}
	}()
	go func() {
		defer wg.Done()
		for val := range channels[2] {
			println("Channel 3 received:", val)
		}
	}()

	wg.Wait()
}

func SplitChannel[T any](inoutCh <-chan T, n int) []<-chan T {
	outChs := make([]chan T, n)
	for i := 0; i < n; i++ {
		outChs[i] = make(chan T)
	}

	go func() {
		idx := 0
		for val := range inoutCh {
			outChs[idx] <- val // this is blocking operation (if needed, can be replaced to non-blocking solution)

			// round-robin algorithm:
			idx = (idx + 1) % n
		}

		for _, ch := range outChs {
			close(ch)
		}
	}()

	// cast []chan T to []<-chanT
	resultChs := make([]<-chan T, n)
	for i := 0; i < n; i++ {
		resultChs[i] = outChs[i]
	}

	return resultChs
}
