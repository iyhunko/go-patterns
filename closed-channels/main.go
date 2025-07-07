package main

import "fmt"

func main() {
	//sendPanics()

	receiveZeroValue()
}

func sendPanics() {
	var c = make(chan int, 100)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			close(c)
		}()
	}
	for i := range c {
		fmt.Println(i)
	}
}

func receiveZeroValue() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	for i := 0; i < 4; i++ {
		res, closed := <-c
		fmt.Printf("%d %v \n", res, closed) // prints 1 2 3 0
	}
}
