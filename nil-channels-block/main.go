package main

import (
	"fmt"
	"time"
)

func printAndWait() {
	fmt.Println("1")
	time.Sleep(5 * time.Second)
	fmt.Println("2")
}

func main() {

	go printAndWait()

	// send to nil channel
	//var c1 chan string
	//c1 <- "let's get started" // deadlock

	// receive from nil channel
	var c2 chan string
	fmt.Println(<-c2) // deadlock

	fmt.Println("Hello, World")
}
