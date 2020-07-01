package main

import (
	"fmt"
)

func main() {

	a := 42
	b := 153
	var temp int

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	temp = a
	a = b
	b = temp

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
