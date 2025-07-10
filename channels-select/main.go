package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	timer := time.NewTimer(1 * time.Second)

	select {
	case v := <-ch1:
		fmt.Println("v = ", v, "from ch1")
	case v := <-ch2:
		fmt.Println("v = ", v, "from ch2")
	case <-timer.C:
		fmt.Println("exited by timer")
	}
}
