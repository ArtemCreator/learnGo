package main

import "fmt"

func main() {
	a := 42
	b := 153

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	fmt.Println("arithmetic method")
	a = a + b // 42 + 153 = 195
	b = b - a // 153 - 192 = -42
	b = -b    // 42
	a = a - b // 195 - 42 = 153

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	fmt.Println("//multiplication/division")
	a = a * b
	b = a / b // деление НЕ целочисленное
	a = a / b
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	fmt.Println("XOR")
	a = a ^ b // 101010 ^ 10011001 = 11000011
	b = b ^ a // 10011001 ^ 11000011 = 101011100
	a = a ^ b //
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	fmt.Println("Other methods")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	a, b = b, a
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	c := 345
	a, b, c = c, a, b
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
