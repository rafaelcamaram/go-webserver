package main

import (
	"fmt"
)

const (
	first = iota
	second
	third
)

func main() {
	// increment := 1

	// fmt.Println("Hello from a module, Gophers!")
	// fmt.Println(first+increment, second+increment, third+increment)

	// increment = 2

	// fmt.Println(first+increment, second+increment, third+increment)

	// Array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println(arr2[1])
}
