package main

import "fmt"

func solution03051()  {
	a := 42
	b := 153

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	// arithmetic method
	a = a + b // 42 + 153 = 195
	b = b - a // 153 - 192 = -42
	b = - b   // 42
	a = a - b // 195 - 42 = 153

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	//multiplication/division
	a = a * b
	b = a / b// деление НЕ целочисленное
	a = a / b
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	//XOR
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println("a:", a)
	fmt.Println("b:", b)



}
