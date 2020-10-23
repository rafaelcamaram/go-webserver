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
	increment := 1

	fmt.Println("Hello from a module, Gophers!")
	fmt.Println(first+increment, second+increment, third+increment)

	increment = 2

	fmt.Println(first+increment, second+increment, third+increment)
}
