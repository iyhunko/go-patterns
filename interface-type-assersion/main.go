package main

import "fmt"

func main() {
	var emptyInterface any
	emptyInterface = "123"

	//intV := emptyInterface.(int)
	//fmt.Println("String value:", intV)

	if intV, ok := emptyInterface.(int); ok {
		fmt.Println("Int value:", intV)
	} else {
		fmt.Println("Not a int")
	}
}
