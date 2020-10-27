package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"foo": 42}
	fmt.Println(m["foo"])

	m["foo"] = 23
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)
}
