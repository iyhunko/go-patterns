package interview_examples

import "fmt"

func main() {
	myMap := make(map[string]string)
	//var myMap map[string]string

	fmt.Println(myMap == nil)

	myMap["c"] = "v1"
	myMap["a"] = "v2"
	myMap["b"] = "v3"
	myMap["d"] = "v4"

	fmt.Printf("%v \n", myMap)

	for k, v := range myMap {
		fmt.Printf("%v %v \n", k, v)
	}
}
