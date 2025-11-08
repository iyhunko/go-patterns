package main

import "fmt"

func main() {
	s := []int{1, 2}

	// index of value that must be removed
	idxToRemove := 0

	fmt.Println(removeFromSliceByIndex(s, idxToRemove))

	fmt.Println(s)
}

func removeFromSliceByIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
