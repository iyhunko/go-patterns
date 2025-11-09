package main

import "fmt"

func main() {
	// Array declaration with fixed size
	var arr1 [3]int
	fmt.Printf("Zero-value array: %v\n", arr1)

	// Array initialization
	arr2 := [3]int{1, 2, 3}
	fmt.Printf("Initialized array: %v\n", arr2)

	// Array with inferred length
	arr3 := [...]int{10, 20, 30, 40}
	fmt.Printf("Inferred length array: %v, length: %d\n", arr3, len(arr3))

	// Arrays are value types - copying creates a new array
	arr4 := arr2
	arr4[0] = 100
	fmt.Printf("Original array after copy modified: %v\n", arr2)
	fmt.Printf("Modified copy: %v\n", arr4)

	// Arrays of same type and size can be compared
	arr5 := [3]int{1, 2, 3}
	arr6 := [3]int{1, 2, 3}
	arr7 := [3]int{1, 2, 4}
	fmt.Printf("arr5 == arr6: %v\n", arr5 == arr6)
	fmt.Printf("arr5 == arr7: %v\n", arr5 == arr7)

	// Passing array to function - passed by value
	modifyArray(arr2)
	fmt.Printf("Array after function call: %v\n", arr2)

	// Passing array pointer to function - can modify original
	modifyArrayPointer(&arr2)
	fmt.Printf("Array after pointer function call: %v\n", arr2)

	// Array vs Slice demonstration
	slice := arr2[:]
	fmt.Printf("Slice from array: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
	slice[0] = 999
	fmt.Printf("Array after slice modification: %v\n", arr2)
}

func modifyArray(arr [3]int) {
	arr[0] = 200
	fmt.Printf("Inside modifyArray: %v\n", arr)
}

func modifyArrayPointer(arr *[3]int) {
	arr[0] = 300
	fmt.Printf("Inside modifyArrayPointer: %v\n", *arr)
}
