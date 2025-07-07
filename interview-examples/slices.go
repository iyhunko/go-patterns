package interview_examples

import "fmt"

func main() {
	urls1 := make([]string, 2, 2)

	urls1[0] = "first"

	urls2 := append(urls1, "second")

	urls2[0] = "new-first"

	fmt.Printf("len: %d \n", len(urls1))
	fmt.Printf("cap: %d \n", cap(urls1))
	fmt.Printf("%v \n", urls1)

	fmt.Printf("len: %d \n", len(urls2))
	fmt.Printf("cap: %d \n", cap(urls2))
	fmt.Printf("%v \n", urls2)
}
