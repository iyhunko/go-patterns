package main

import "fmt"

func main() {
	// original slice
	s := []int{1, 2}

	// index of value that must be removed
	idxToRemove := 0

	// removeFromSliceByIndex returns a slice with the element at idxToRemove removed.
	// IMPORTANT: the returned slice may share the same backing array as `s`,
	// so operations inside the function can mutate the underlying array that `s` points to.
	fmt.Println(removeFromSliceByIndex(s, idxToRemove))

	// Because the implementation below reuses the backing array when possible,
	// `s` itself may be modified even though we never reassign `s` in main.
	// For this example the output will be:
	// [2]
	// [2 2]
	fmt.Println(s)
}

func removeFromSliceByIndex(s []int, index int) []int {
	// This line concatenates the portion before index with the portion after index:
	// - s[:index] produces a slice of all elements before index (length == index)
	// - s[index+1:] produces a slice of all elements after index
	// - append(a, b...) appends all elements of b to a and returns the resulting slice header.
	//
	// Important details and pitfalls:
	// 1. Bounds: if index is out of range (index < 0 or index >= len(s)),
	//    the slicing operations will panic. The function as written does not check bounds.
	// 2. Backing array reuse: append attempts to reuse the capacity of the first slice
	//    (s[:index]). If it fits, append will write into the existing backing array,
	//    overwriting elements starting at position `index`. The returned slice and the
	//    original `s` will share the same underlying array, so the original `s` may
	//    observe changes even when its length is unchanged.
	// 3. If you need an independent slice (no shared backing array), copy into a new
	//    backing array, for example:
	//      out := append([]int(nil), s[:index]...)
	//      out = append(out, s[index+1:]...)
	//    or
	//      out := make([]int, 0, len(s)-1)
	//      out = append(out, s[:index]...)
	//      out = append(out, s[index+1:]...)
	//
	// Example: s := []int{1,2}; index := 0
	// - s[:0] has len 0, cap 2 and points at the same array [1,2]
	// - append(s[:0], s[1:]...) writes the value 2 into the backing array at index 0,
	//   resulting in the backing array becoming [2,2].
	// - returned slice is [2] (len 1), but original `s` now observes [2,2].
	return append(s[:index], s[index+1:]...)
}
