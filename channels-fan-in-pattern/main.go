package main

import (
	"context"
	"sync"
)

// fanIn takes data from multiple input channels and put them into single output channel
func fanIn(ctx context.Context, chans []chan int) chan int {
	out := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		for _, ch := range chans {
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
