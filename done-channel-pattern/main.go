package main

import "fmt"

func main() {
	closeCh := make(chan struct{})
	closeDoneCh := process(closeCh)

	close(closeCh)
	<-closeDoneCh

	fmt.Println("Channels closed, exiting program.")
}

func process(closeCh <-chan struct{}) <-chan struct{} {
	closeDoneCh := make(chan struct{})

	go func() {
		defer close(closeDoneCh)

		for {
			select {
			case <-closeCh:
				fmt.Println("Received close signal, cleaning up...")
				return
			default:
				// Simulate some work
				fmt.Println("Processing...")
			}
		}
	}()

	return closeDoneCh
}
