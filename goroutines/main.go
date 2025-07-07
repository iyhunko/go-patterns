package main

import (
	"fmt"
	"time"
)

func print1() {
	for {
		fmt.Println("1")
	}
}

func print2() {
	for {
		fmt.Println("2")
	}
}

func main() {

	go print1()
	go print2()

	time.Sleep(5 * time.Second)
}
