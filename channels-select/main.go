package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	timer := time.NewTimer(1 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Microsecond)
	defer cancel()

	select {
	case v := <-ch1:
		fmt.Println("v = ", v, "from ch1")
	case v := <-ch2:
		fmt.Println("v = ", v, "from ch2")
	case <-timer.C:
		fmt.Println("exited by timer")
	case <-ctx.Done():
		fmt.Println("exited by context:", ctx.Err())
	}
}
